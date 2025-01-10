package userHandler

import (
	"github.com/labstack/echo/v4"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

func RegisterUserHandler(e *echo.Echo, cfg *config.Config) {
	handlers := New(cfg)
	server := e.Group("/user")
	server.POST("/register", handlers.Register)
}
