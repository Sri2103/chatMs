package sqlRepo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	database "github.com/sri2103/chat_me/DB"
	"github.com/sri2103/chat_me/services/user/config"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type sqliteRepo struct {
	db *sql.DB
}

func NewPostgresRepo(cfg *config.Config) service.RepoInterface {
	db, err := database.NewPostgresConnection(database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "harsha",
		Password: "password",
		DBName:   "main",
	})

	if err != nil {
		log.Fatal(err)
	}
	return &sqliteRepo{
		db: db,
	}
}

func (r *sqliteRepo) createUser(tx *sql.Tx, user *model.UserModel) error {
	query := `
		INSERT INTO public.users (user_id, username, email, "role", password_hash)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := tx.Exec(query, user.UserId, user.UserName, user.Email, user.Role, user.PasswordHash)

	if err != nil {
		return err
	}
	return nil
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
	query := `select * from public.users where user_id=$1`
	row := tx.QueryRow(query, user.UserId)
	if row.Err() != nil && errors.Is(row.Err(), sql.ErrNoRows) {
		// create user
		return r.createUser(tx, user)
	} else if row.Err() != nil {
		return row.Err()
	}

	_, err = tx.Exec("update public.users set username=?,email=?,role=?", user.UserName, user.Email, user.Role)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func (r *sqliteRepo) FindUserByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	email = strings.TrimSpace(email)
	query := `select user_id, username, email, "role", password_hash from users u where email=$1`
	row := r.db.QueryRowContext(ctx, query, email)
	var usr model.UserModel
	if row.Err() != nil {
		return &usr, row.Err()
	}
	err := row.Scan(
		&usr.UserId,
		&usr.UserName,
		&usr.Email,
		&usr.Role,
		&usr.PasswordHash,
	)
	if err != nil {
		return &usr, err
	}
	return &usr, nil
}

func (r *sqliteRepo) CreateUser(ctx context.Context, user *model.UserModel) error {
	id := uuid.New().String()
	user.UserId = id
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	err = r.createUser(tx, user)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
