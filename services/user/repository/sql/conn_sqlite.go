package sqlRepo

import (
	"context"
	"database/sql"
	"errors"
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

func (r *sqliteRepo) SaveUser(ctx context.Context, user *model.UserModel) error {
	tx, err := r.db.BeginTx(ctx, nil)
	defer func() {
		_ = tx.Rollback()
	}()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	row := tx.QueryRow("select * from User where user_id=?", user.UserId)
	if row.Err() != nil && errors.Is(row.Err(), sql.ErrNoRows) {
		// create user
	} else if row.Err() != nil {
		return row.Err()
	}

	_, err = tx.Exec("update User set username=?,email=?,role=?", user.UserName, user.Email, user.Role)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func (r *sqliteRepo) FindUserByIndentifier(ctx context.Context, query string) *model.UserModel {
	panic("not implemented") // TODO: Implement
}
