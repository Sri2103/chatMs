package service

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
)

type AuthService interface {
	CreateAuth(context.Context, *model.AuthDetails) error
	FetchAuth(context.Context, *model.AuthDetails) (*model.AuthModel, error)
	DeleteAuth(context.Context, string) error
}

type AuthServiceRepo interface {
	CreateAuth(context.Context, *model.AuthDetails) error
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

func (a AuthServiceReceiver) CreateAuth(_ context.Context, _ *model.AuthDetails) error {
	panic("not implemented") // TODO: Implement
}

func (a AuthServiceReceiver) FetchAuth(_ context.Context, _ *model.AuthDetails) (*model.AuthModel, error) {
	panic("not implemented") // TODO: Implement
}

func (a AuthServiceReceiver) DeleteAuth(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
