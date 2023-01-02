package http

import (
	"bytes"
	"encoding/json"
	"expvar"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/narslan/dezge"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func (s *Server) handleEcho(c *gin.Context) {

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {

		c.Abort()
	}
	g, _ := errgroup.WithContext(c)
	//errCh := make(chan error)

	g.Go(func() error {
		defer conn.Close()

		for {
			_, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				return err
			}
			var buf bytes.Buffer
			buf.WriteString(`"`)

			// for _, m := range msg {
			// 	buf.WriteByte(m)
			// }
			buf.WriteString(time.Now().String())
			buf.WriteString(`"`)
			err = wsutil.WriteServerMessage(conn, op, buf.Bytes())
			if err != nil {
				return err
			}
		}
	})
	err = g.Wait()
	if err != nil {
		c.Abort()
	}
}

func (s *Server) handleEvent(c *gin.Context) {

	id := c.Query("id")

	uid, err := strconv.ParseUint(id, 10, 64)
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

	log.Debug().Msgf("ID: %d Name: %s", u.ID, u.Name)

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

func (s *Server) handleEventRandomUser(c *gin.Context) {

	var u dezge.User

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	u.ID = r1.Int()
	u.Name = fmt.Sprintf("no-name: %d", u.ID)

	ctx := dezge.NewContextWithUser(c, &u)
	sub, err := s.EventService.Subscribe(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "subscription err" + err.Error()})
		return
	}

	defer sub.Close()

	log.Debug().Msgf("ID: %d Name: %s", u.ID, u.Name)

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {

		c.Abort()
	}
	//g, _ := errgroup.WithContext(c)
	//errCh := make(chan error)

	for {
		select {
		case <-c.Done():
			return // disconnect when HTTP connection disconnects
		case event, ok := <-sub.C():
			// If subscription is closed then exit.
			if !ok {
				return
			}

			// Marshal event data to JSON.
			buf, err := json.Marshal(event)
			if err != nil {

				return
			}

			err = wsutil.WriteServerMessage(conn, ws.OpCode(ws.StateServerSide), buf)
			if err != nil {
				return
			}
		}

	}
}

func (s *Server) handlePublishEvent(c *gin.Context) {

	if s.EventService.LenSubs() == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "no subscribers"})
		return
	}

	var u dezge.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ev dezge.Event
	ev.Type = "Publish"
	ev.Payload = fmt.Sprintf("ID: %d on %s", u.ID, time.Now().Format(time.UnixDate))

	s.EventService.PublishEvent(u.ID, ev)

	c.JSON(http.StatusAccepted, gin.H{"message": "publication succesful"})

}
