package handler

import (
	"context"

	"github.com/sri2103/chat_me/protos/room"
	"github.com/sri2103/chat_me/services/room/service"
)

type Handler struct {
	service service.Service
	room.UnimplementedRoomServiceServer
}

func NewGrpcHandler(svc service.Service) room.RoomServiceServer {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) CreateRoom(_ context.Context, _ *room.CreateRoomRequest) (*room.CreateRoomResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) GetRoomDetails(_ context.Context, _ *room.GetRoomRequest) (*room.GetRoomResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) UpdateRoomSettings(_ context.Context, _ *room.UpdateRoomRequest) (*room.UpdateRoomResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) mustEmbedUnimplementedRoomServiceServer() {
	panic("not implemented") // TODO: Implement
}
