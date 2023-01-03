package http

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type EventData struct {
	Created time.Time `json:"created"`
	Text    string    `json:"text"`
}

func (s *Server) timeStream(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")

	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		ticker := time.NewTicker(500 * time.Millisecond)
		//done := make(chan bool)
		tch := make(chan time.Time)
		cCp := c.Copy()
		go func() {
			for {
				select {
				case <-cCp.Done():

					return
				case t := <-ticker.C:
					tch <- t
				}
			}
		}()
		ev := EventData{Created: <-tch, Text: "Merhaba"}

		c.SSEvent("message", ev)
		return true

	})
	//c.JSON(200, gin.H{"subs": s.EventService.LenSubs()})
}
