package sqlRepo

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sri2103/chat_me/services/user/config"
	"github.com/sri2103/chat_me/services/user/model"
	"github.com/sri2103/chat_me/services/user/service"
)

type SqlAuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(cfg *config.Config) service.AuthServiceRepo {
	return &SqlAuthRepo{
		db: cfg.Db,
	}

}

func (r *SqlAuthRepo) CreateAuth(ctx context.Context, m *model.AuthDetails) (*model.AuthModel, error) {
	ID := uuid.NewString()
	query := `INSERT INTO auth (id, user_id, auth_id) VALUES($1, $2, $3);`
	_, err := r.db.Exec(query, ID, m.UserId, m.AuthId)
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
