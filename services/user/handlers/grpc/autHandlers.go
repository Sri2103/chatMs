package grpcHandler

import (
	"context"

	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/model"
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

func (h *AuthHandler) CreateAuth(ctx context.Context, req *user.CreateAuthRequest) (*user.CreateAuthResponse, error) {
	authModel := model.AuthDetails{
		AuthId: req.GetAuthId(),
		UserId: req.GetUserId(),
	}

	mdl, err := h.authService.CreateAuth(ctx, &authModel)
	if err != nil {
		return nil, err
	}

	var resp user.CreateAuthResponse
	resp.UserId = mdl.UserId
	resp.Id = mdl.ID
	resp.AuthId = mdl.AuthId

	return &resp, nil
}

func (h *AuthHandler) DestroyAuth(_ context.Context, _ *user.DestroyAuthRequest) (*user.DestroyAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *AuthHandler) FetchAuth(_ context.Context, _ *user.FetchAuthRequest) (*user.FetchAuthResponse, error) {
	panic("not implemented") // TODO: Implement
}
