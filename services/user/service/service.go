package service

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
)

type RepoInterface interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUserByIndentifier(ctx context.Context, query string) (*model.UserModel, error)
}

type Service interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUser(ctx context.Context, query string) (*model.UserModel, error)
}

type service struct {
	repo RepoInterface
}

func New(r RepoInterface) Service {
	return &service{
		repo: r,
	}
}

func (f *service) SaveUser(ctx context.Context, user *model.UserModel) error {
	return f.SaveUser(ctx, user)
}

func (f *service) FindUser(ctx context.Context, query string) (*model.UserModel, error) {
	return f.FindUser(ctx, query)
}
