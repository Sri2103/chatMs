package sqlRepo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sri2103/chat_me/services/user/config"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type sqliteRepo struct {
	db *sql.DB
}

func NewSqlRepo(cfg *config.Config) service.RepoInterface {
	Db, err := sql.Open("sqlite3", cfg.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	return &sqliteRepo{
		db: Db,
	}
}

func (r *sqliteRepo) SaveUser(user *model.UserModel) error {
	panic("not implemented") // TODO: Implement
}

func (r *sqliteRepo) FindUserByIndentifier(query string) *model.UserModel {
	panic("not implemented") // TODO: Implement
}
