package service

import (
	"context"
	"errors"

	"github.com/sri2103/chat_me/services/user/model"
)

type RepoInterface interface {
	SaveUser(ctx context.Context, user *model.UserModel) error
	FindUserByEmail(ctx context.Context, email string) (*model.UserModel, error)
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
	return f.repo.FindUserByEmail(ctx, query)
}

func (f *service) AuthenticateUser(ctx context.Context, auth *model.Credentials) (*model.UserModel, error) {
	md, err := f.repo.FindUserByEmail(ctx, auth.Email)
	if err != nil {
		return nil, err
	}
	if md.PasswordHash != auth.Password {
		return nil, errors.New("password error")
	}
	return md, nil
}

func (f *service) CreateUser(ctx context.Context, user *model.UserModel) error {
	return f.repo.CreateUser(ctx, user)
}
