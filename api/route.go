package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Service) GetRoutes() {
	prefix := s.Group("/api/")

	prefix.GET("ping", handlePing)
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"status":    http.StatusOK,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	})
}
