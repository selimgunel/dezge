package http

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
)

func TestBasicRegister(t *testing.T) {
	r := gofight.New()

	r.GET("/ping").
		// turn on the debug mode.
		SetDebug(false).
		Run(BasicEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "pong", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
