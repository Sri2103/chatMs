package userHandler

import (
	"time"

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
	server.POST("/login", handlers.LoginUser)
	server.GET("/logout", handlers.Logout, echojwt.WithConfig(claimConfig))
}

type jwtCustomClaims struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	AuthId string `json:"auth_id"`
	jwt.RegisteredClaims
}

type TokenPayload struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	AuthId string `json:"auth_id"`
}

func CreateToken(p *TokenPayload) (*string, error) {
	claims := &jwtCustomClaims{
		p.Name,
		p.UserId,
		p.AuthId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
