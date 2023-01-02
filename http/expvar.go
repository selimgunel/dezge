package http

import (
	"bytes"
	"expvar"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"golang.org/x/sync/errgroup"
)

func (s *Server) handleDebug(c *gin.Context) {

	w := c.Writer
	c.Header("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte("{\n"))
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			_, _ = w.Write([]byte(",\n"))
		}
		first = false
		fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
	})
	_, _ = w.Write([]byte("\n}\n"))
	c.AbortWithStatus(200)

}

func (s *Server) handleDebugWS(c *gin.Context) {

	conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
	if err != nil {

		c.Abort()
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	g, _ := errgroup.WithContext(c)
	//errCh := make(chan error)

	g.Go(func() error {
		defer conn.Close()

		for range ticker.C {

			metrics := harvestMetrics()
			err = wsutil.WriteServerMessage(conn, ws.OpCode(ws.StateServerSide), metrics)
			if err != nil {
				return err
			}

		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		c.Abort()
	}
}

func harvestMetrics() []byte {
	var w bytes.Buffer

	_, _ = w.Write([]byte("{\n"))
	first := true
	expvar.Do(func(kv expvar.KeyValue) {
		if !first {
			_, _ = w.Write([]byte(",\n"))
		}
		first = false
		fmt.Fprintf(&w, "%q: %s", kv.Key, kv.Value)
	})
	_, _ = w.Write([]byte("\n}\n"))
	return w.Bytes()
}
