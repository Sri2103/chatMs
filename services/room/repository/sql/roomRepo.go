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
	_, err := r.Db.Exec(
		"insert into room (room_id, room_name,created_by, users) values (?,?,?,?)",
		room.RoomId, room.RoomName, room.CreatedBy, room.Users)

	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r *SqlRepo) FindById(ctx context.Context, roomId string) (*model.Room, error) {
	row := r.Db.QueryRow(
		"select * from room where room_id=?",
		roomId,
	)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var room model.Room
	err := row.Scan(&room.RoomId, &room.RoomName, &room.CreatedBy, &room.Users)
	if err != nil {
		return nil, err
	}

	return &room, nil

}

func (r *SqlRepo) UpdateRoom(ctx context.Context, room *model.Room) error {
	_, err := r.Db.Exec(
		"update room set created_by=?, users=?, room_name=? where room_id=?",
		room.CreatedBy, room.Users, room.RoomName, room.RoomId,
	)

	if err != nil {
		return err
	}

	return nil

}
