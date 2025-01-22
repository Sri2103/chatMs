package service

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
)

type RepoInterface interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUserByIndentifier(ctx context.Context, query string) (*model.UserModel, error)
	CreateUser(ctx context.Context, user *model.UserModel) error
}

type Service interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUser(ctx context.Context, query string) (*model.UserModel, error)
	AuthenticateUser(ctx context.Context, auth *model.Credentials) (*model.UserModel, error)
	CreateUser(ctx context.Context, user *model.UserModel) error
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
	return f.repo.SaveUser(ctx, user)
}

func (f *service) FindUser(ctx context.Context, query string) (*model.UserModel, error) {
	return f.repo.FindUserByIndentifier(ctx, query)
}

func (f *service) AuthenticateUser(ctx context.Context, auth *model.Credentials) (*model.UserModel, error) {
	panic("Not implemented") //TODO: implement this method
}

func (f *service) CreateUser(ctx context.Context, user *model.UserModel) error {
	return f.repo.CreateUser(ctx, user)
}
