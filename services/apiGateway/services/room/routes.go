package roomHandler

import (
	"github.com/labstack/echo/v4"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

func RegisterRoomHandlers(e *echo.Echo, cfg *config.Config) {
	handlers := New(cfg)
	server := e.Group("/room")
	server.POST("/create", handlers.CreateRoom)
	server.GET("/:roomId", handlers.FetchRoom)
	server.PUT("/update", handlers.UpdateRoom)

}
