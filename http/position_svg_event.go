package http

import (
	"encoding/json"
	"expvar"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/narslan/dezge"
)

func (s *Server) subscribePositionSVG(c *gin.Context) {

	id := c.Query("id")

	uid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid is not a valid number"})
		return
	}
	var u dezge.User
	u.ID = int(uid)
	u.Name = "none"

	ctx := dezge.NewContextWithUser(c, &u)
	sub, err := s.EventService.Subscribe(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription err" + err.Error()})
		return
	}

	defer sub.Close()

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {

		c.Abort()
	}
	//g, _ := errgroup.WithContext(c)
	//errCh := make(chan error)
	s.stat.Get("counter").(*expvar.Int).Add(1)
	for {
		select {
		case <-c.Done():
			s.stat.Get("counter").(*expvar.Int).Add(-1)
			return // disconnect when HTTP connection disconnects
		case event, ok := <-sub.C():
			// If subscription is closed then exit.
			if !ok {
				s.stat.Get("counter").(*expvar.Int).Add(-1)
				return
			}

			// Marshal event data to JSON.
			buf, err := json.Marshal(event)
			if err != nil {
				s.stat.Get("counter").(*expvar.Int).Add(-1)

				return
			}
			err = wsutil.WriteServerMessage(conn, ws.OpCode(ws.StateServerSide), buf)
			if err != nil {
				s.stat.Get("counter").(*expvar.Int).Add(-1)
				return
			}
		}

	}
}
