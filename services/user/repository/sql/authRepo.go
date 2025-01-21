package sqlRepo

import (
	"context"
	"database/sql"
	"log"

	"github.com/sri2103/chat_me/services/user/config"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type SqlAuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(cfg *config.Config) service.AuthService {
	Db, err := sql.Open("sqlite3", cfg.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	return &SqlAuthRepo{
		db: Db,
	}

}

func (r *SqlAuthRepo) CreateAuth(_ context.Context, _ *model.AuthDetails) error {
	panic("not implemented") // TODO: Implement
}

func (r *SqlAuthRepo) FetchAuth(_ context.Context, _ *model.AuthDetails) (*model.AuthModel, error) {
	panic("not implemented") // TODO: Implement
}

func (r *SqlAuthRepo) DeleteAuth(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
