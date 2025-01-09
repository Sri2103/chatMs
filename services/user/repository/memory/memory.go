package memory

import (
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type repo struct{}

func New() service.RepoInterface {
	return &repo{}
}

func (f *repo) SaveUser(user *model.UserModel) error {
	panic("not implemented") // TODO: Implement
}

func (f *repo) FindUserByIndentifier(query string) *model.UserModel {
	panic("not implemented") // TODO: Implement
}
