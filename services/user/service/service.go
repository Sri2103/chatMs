package service

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
)

type RepoInterface interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUserByIndentifier(ctx context.Context, query string) *model.UserModel
}

type Service interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	findUser(ctx context.Context, query string) *model.UserModel
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
	panic("not implemented") // TODO: Implement
}

func (f *service) findUser(ctx context.Context, query string) *model.UserModel {
	panic("not implemented") // TODO: Implement
}
