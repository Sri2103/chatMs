package grpcHandler

import (
	"context"

	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/service"
)

type AuthHandler struct {
	authService service.AuthService
	user.UnimplementedAuthServiceServer
}

func NewGrpcAuthHandler(svc service.AuthService) user.AuthServiceServer {
	return &AuthHandler{
		authService: svc,
	}
}

func (h *AuthHandler) CreateAuth(_ context.Context, _ *user.CreateAuthRequest) (*user.CreateAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *AuthHandler) DestroyAuth(_ context.Context, _ *user.DestroyAuthRequest) (*user.DestroyAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *AuthHandler) FetchAuth(_ context.Context, _ *user.FetchAuthRequest) (*user.FetchAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}
