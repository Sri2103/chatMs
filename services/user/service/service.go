package service

import "github.com/sri2103/chat_me/services/user/model"

type RepoInterface interface {
	SaveUser(user *model.UserModel) error
	FindUserByIndentifier(query string) *model.UserModel
}

type Service interface {
	SaveUser(user *model.UserModel) error
	findUser(query string) *model.UserModel
}

type service struct {
	repo RepoInterface
}

func New(r RepoInterface) Service {
	return &service{
		repo: r,
	}
}

func (f *service) SaveUser(user *model.UserModel) error {
	panic("not implemented") // TODO: Implement
}

func (f *service) findUser(query string) *model.UserModel {
	panic("not implemented") // TODO: Implement
}
