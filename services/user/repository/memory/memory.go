package memory

import (
	"context"

	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type repo struct{}

func New() service.RepoInterface {
	return &repo{}
}

func (f *repo) SaveUser(ctx context.Context, user *model.UserModel) error {
	panic("not implemented") // TODO: Implement
}

func (f *repo) FindUserByIndentifier(ctx context.Context, query string) (*model.UserModel, error) {
	panic("not implemented") // TODO: Implement
}

func (f *repo) CreateUser(ctx context.Context, user *model.UserModel) error {
	panic("not implementd")
}
