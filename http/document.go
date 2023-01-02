package http

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

type DocData struct {
	Text string `json:"text"`
}

func NewDocData() *DocData {
	return &DocData{Text: `(lambda (x) (display "hello" x))`}
}

func (s *Server) documentStream(c *gin.Context) {

	d := NewDocData()
	c.JSON(http.StatusOK, d)

}
