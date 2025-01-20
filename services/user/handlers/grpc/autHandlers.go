package grpcHandler

import (
	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/service"
)

type authHandler struct {
	authService service.AuthService
	user.UnimplementedAuthServiceServer
}

func NewGrpcAuthHandler(svc service.AuthService) user.AuthServiceServer {
	return &authHandler{
		authService: svc,
	}
}
