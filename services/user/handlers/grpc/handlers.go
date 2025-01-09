package grpcHandler

import (
	"context"

	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/service"
)

type Handler struct {
	service service.Service
	user.UnimplementedUserServiceServer
}

func NewGrpcServerHandler(srv service.Service) user.UserServiceServer {
	return &Handler{}
}

func (h *Handler) AuthenticateUser(_ context.Context, _ *user.AuthenticateRequest) (*user.AuthenticateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) GetUserDetails(_ context.Context, _ *user.GetUserRequest) (*user.GetUserResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) UpdateUserDetails(_ context.Context, _ *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	panic("not implemented") // TODO: Implement
}
