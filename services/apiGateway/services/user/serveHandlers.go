package userHandler

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

func RegisterUserHandler(e *echo.Echo, cfg *config.Config) {
	handlers := New(cfg)
	server := e.Group("/user")

	claimConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	server.POST("/register", handlers.Register)
	server.GET("/logout", func(c echo.Context) error {
		return nil
	}, echojwt.WithConfig(claimConfig))
}

type jwtCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}
