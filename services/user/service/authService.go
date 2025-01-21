package service

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
)

type AuthService interface {
	CreateAuth(context.Context, *model.AuthDetails) (*model.AuthModel, error)
	FetchAuth(context.Context, *model.AuthDetails) (*model.AuthModel, error)
	DeleteAuth(context.Context, string) error
}

type AuthServiceRepo interface {
	CreateAuth(context.Context, *model.AuthDetails) (*model.AuthModel, error)
	FetchAuth(context.Context, *model.AuthDetails) (*model.AuthModel, error)
	DeleteAuth(context.Context, string) error
}

type AuthServiceReceiver struct {
	repo AuthServiceRepo
}

func NewAuth(r AuthServiceRepo) AuthService {
	return AuthServiceReceiver{
		repo: r,
	}
}

func (a AuthServiceReceiver) CreateAuth(ctx context.Context, auth *model.AuthDetails) (*model.AuthModel, error) {
	return a.repo.CreateAuth(ctx, auth)
}

func (a AuthServiceReceiver) FetchAuth(ctx context.Context, auth *model.AuthDetails) (*model.AuthModel, error) {
	return a.repo.FetchAuth(ctx, auth)
}

func (a AuthServiceReceiver) DeleteAuth(ctx context.Context, authId string) error {
	return a.repo.DeleteAuth(ctx, authId)
}
