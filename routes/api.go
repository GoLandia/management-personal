package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Api() {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"status":    http.StatusOK,
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	route.Run()
}
