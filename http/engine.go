package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narslan/dezge"
)

func (s *Server) ListEngines(c *gin.Context) {

	allEngines, _, err := s.EngineInfoService.Find(context.Background(), dezge.EngineFilter{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, allEngines)

}
