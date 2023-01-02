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
	"github.com/rs/zerolog/log"
)

func (s *Server) subscribePositionSVG(c *gin.Context) {

	id := c.Query("id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uid is not a valid number"})
		return
	}
	var u dezge.User
	u.ID = uid
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

func (s *Server) publishPositionSVG(c *gin.Context) {
	// if s.EventService.LenSubs() == 0 {
	// 	c.JSON(http.StatusNotAcceptable, gin.H{"error": "no subscribers"})
	// 	return
	// }

	// var u dezge.User

	// if err := c.ShouldBindJSON(&u); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// var b strings.Builder

	// // create board position
	// fenStr := "rnbqkbnr/pppppppp/8/8/3P4/8/PPP1PPPP/RNBQKBNR b KQkq - 0 1"
	// pos := &chess.Position{}
	// if err := pos.UnmarshalText([]byte(fenStr)); err != nil {
	// 	log.Fatal().Msg("")
	// }

	// // write board SVG to file
	// yellow := color.RGBA{255, 255, 0, 1}
	// mark := chessimg.MarkSquares(yellow, chess.D2, chess.D4)
	// if err := chessimg.SVG(&b, pos.Board(), mark); err != nil {
	// 	log.Fatal().Msg("")
	// }
	// var ev dezge.Event
	// ev.Type = "PGNSVG"
	// ev.Payload = b.String()

	// s.EventService.PublishEvent(u.ID, ev)

}
