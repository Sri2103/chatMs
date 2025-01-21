package service

import (
	"context"

	"github.com/sri2103/chat_me/services/room/model"
)

type Service interface {
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	GetRoomDetails(ctx context.Context, roomId string) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
}

type RoomRepository interface {
	SaveRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	FindById(ctx context.Context, roomId string) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
}

type RoomService struct {
	repo RoomRepository
}

func New(r RoomRepository) Service {
	return &RoomService{
		repo: r,
	}
}

func (r *RoomService) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RoomService) GetRoomDetails(ctx context.Context, roomId string) (*model.Room, error) {
	panic("not implemented") // TODO: Implement
}

func (r *RoomService) UpdateRoom(ctx context.Context, room *model.Room) error {
	panic("not implemented") // TODO: Implement
}
