package sqlRepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *sqliteRepo) createUser(tx *sql.Tx, user *model.UserModel) error {
	_, err := tx.Exec("insert into User (user_id, username,email,password_hash,role) values (?,?,?,?,?,?)", user.UserId, user.UserName,
		user.PasswordHash, user.Role)
	return err
}

func (r *sqliteRepo) SaveUser(ctx context.Context, user *model.UserModel) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	row := tx.QueryRow("select * from User where user_id=?", user.UserId)
	if row.Err() != nil && errors.Is(row.Err(), sql.ErrNoRows) {
		// create user
		return r.createUser(tx, user)
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

func (r *sqliteRepo) FindUserByIndentifier(ctx context.Context, query string) (*model.UserModel, error) {
	row := r.db.QueryRowContext(ctx, fmt.Sprintf("select * from User where %s", query))
	var usr model.UserModel
	if row.Err() != nil {
		return &usr, row.Err()
	}
	err := row.Scan(
		usr.UserId,
		usr.UserName,
		usr.Email,
		usr.PasswordHash,
		usr.Role,
	)
	if err != nil {
		return &usr, err
	}
	return &usr, nil
}
