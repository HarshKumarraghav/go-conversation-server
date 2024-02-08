package routes

import (
	"conversationserver/pkg/chat"

	"github.com/gin-gonic/gin"
)

func WebSocketRoomConnectionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomId := c.Param("roomId")
		chat.ServeWs(c.Writer, c.Request, roomId)
	}
}

func CreateChatRoomRoute(app *gin.Engine) {
	app.GET("/ws/:roomId", WebSocketRoomConnectionHandler())
}
