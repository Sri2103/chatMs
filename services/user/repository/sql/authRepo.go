package sqlRepo

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/sri2103/chat_me/services/user/config"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type SqlAuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(cfg *config.Config) service.AuthServiceRepo {
	Db, err := sql.Open("sqlite3", cfg.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	return &SqlAuthRepo{
		db: Db,
	}

}

func (r *SqlAuthRepo) CreateAuth(ctx context.Context, m *model.AuthDetails) (*model.AuthModel, error) {
	ID := uuid.NewString()
	_, err := r.db.Exec("create into auth (Id,user_id,auth_id) values (?,?,?)", ID, m.UserId, m.AuthId)
	if err != nil {
		return nil, err
	}

	return &model.AuthModel{
		ID:     ID,
		UserId: m.UserId,
		AuthId: m.AuthId,
	}, nil
}

func (r *SqlAuthRepo) FetchAuth(ctx context.Context, m *model.AuthDetails) (*model.AuthModel, error) {
	row := r.db.QueryRow("select * from auth where auth_id=?, user_id=?", m.AuthId, m.UserId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var md model.AuthModel
	err := row.Scan(md.ID, md.UserId, md.AuthId)
	if err != nil {
		return nil, err
	}

	return &md, nil

}

func (r *SqlAuthRepo) DeleteAuth(ctx context.Context, authId string) error {
	_, err := r.db.Exec("delete from auth where auth_id=?", authId)
	if err != nil {
		return err
	}

	return nil
}
