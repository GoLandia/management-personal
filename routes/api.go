package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func Api() {
	route := gin.Default()

	prefix := route.Group("/api/")

	prefix.GET("ping", handlePing)
	prefix.GET("websocket", handleWebsocket)
	route.Run()
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":   "pong",
		"status":    http.StatusOK,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	})
}

func handleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   string(message),
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
}
