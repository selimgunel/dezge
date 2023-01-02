package http

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

const httpPort = ":8888"

func BasicEngine() http.Handler {
	s := NewServer(httpPort, nil)
	return s.server.Handler

}
func TestBasicPing(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		// turn on the debug mode.
		SetDebug(true).
		Run(BasicEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "pong", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
