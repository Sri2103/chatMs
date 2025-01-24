package grpcHandler

import (
	"context"
	"fmt"

	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type Handler struct {
	service service.Service
	user.UnimplementedUserServiceServer
}

func NewGrpcServerHandler(srv service.Service) user.UserServiceServer {
	return &Handler{service: srv}
}

func (h *Handler) AuthenticateUser(ctx context.Context, gr *user.AuthenticateRequest) (*user.AuthenticateResponse, error) {
	email := gr.GetEmail()
	password := gr.GetPassword()
	usr, err := h.service.AuthenticateUser(ctx, &model.Credentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	return &user.AuthenticateResponse{
		UserId: usr.UserId,

		Username:      usr.UserName,
		Authenticated: true,
	}, nil

}

func (h *Handler) GetUserDetails(ctx context.Context, gr *user.GetUserRequest) (*user.GetUserResponse, error) {
	user_id := gr.GetUserId()
	usr, err := h.service.FindUser(ctx, fmt.Sprintf("%s=%s", "user_id", user_id))
	if err != nil {
		return nil, err
	}
	return &user.GetUserResponse{
		UserId:   usr.UserId,
		Email:    usr.Email,
		Username: usr.UserName,
	}, nil

}

func (h *Handler) UpdateUserDetails(ctx context.Context, gr *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	var usr model.UserModel
	usr.UserId = gr.GetUserId()
	usr.Email = gr.GetEmail()
	usr.UserName = gr.GetUsername()
	err := h.service.SaveUser(ctx, &usr)
	if err != nil {
		return nil, err
	}

	resp := &user.UpdateUserResponse{
		Success: true,
	}
	return resp, nil

}

func (h *Handler) RegisterUserDetails(ctx context.Context, rr *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	var usr model.UserModel
	usr.Email = rr.GetEmail()
	usr.Role = model.Participant
	usr.UserName = rr.GetUsername()
	usr.PasswordHash = rr.GetPassword()

	fmt.Println("register user details starting")
	err := h.service.CreateUser(ctx, &usr)
	if err != nil {
		return nil, err
	}

	return &user.RegisterUserResponse{
		Success: true,
	}, nil

}
