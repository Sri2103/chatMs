package sqlRepo

import (
	"context"
	"database/sql"
	"log"

	"github.com/sri2103/chat_me/services/room/config"
	"github.com/sri2103/chat_me/services/room/model"
	"github.com/sri2103/chat_me/services/room/service"
)

type SqlRepo struct {
	Db *sql.DB
}

func NewSqlRepo(cfg *config.Config) service.RoomRepository {
	Db, err := sql.Open("sqlite3", cfg.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	return &SqlRepo{
		Db: Db,
	}
}

func (r *SqlRepo) SaveRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	panic("not implemented") // TODO: Implement
}

func (r *SqlRepo) FindById(ctx context.Context, roomId string) (*model.Room, error) {
	panic("not implemented") // TODO: Implement
}

func (r *SqlRepo) UpdateRoom(ctx context.Context, room *model.Room) error {
	panic("not implemented") // TODO: Implement
}
