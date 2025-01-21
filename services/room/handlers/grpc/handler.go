package handler

import (
	"context"

	"github.com/sri2103/chat_me/protos/room"
	"github.com/sri2103/chat_me/services/room/model"
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

func (h *Handler) CreateRoom(ctx context.Context, req *room.CreateRoomRequest) (*room.CreateRoomResponse, error) {
	rm := &model.Room{
		RoomName:  req.RoomName,
		CreatedBy: req.CreatedBy,
		Users:     req.GetUserIds(),
	}
	rd, err := h.service.CreateRoom(ctx, rm)

	if err != nil {
		return nil, err
	}

	var res room.CreateRoomResponse
	res.Success = true
	res.RoomId = rd.RoomId
	return &res, nil
}

func (h *Handler) GetRoomDetails(ctx context.Context, req *room.GetRoomRequest) (*room.GetRoomResponse, error) {
	rId := req.RoomId
	rm, err := h.service.GetRoomDetails(ctx, rId)
	if err != nil {
		return nil, err
	}
	var res room.GetRoomResponse
	res.RoomId = rm.RoomId
	res.RoomName = rm.RoomName
	res.CreatedBy = rm.CreatedBy
	res.UserIds = rm.Users
	return &res, nil
}

func (h *Handler) UpdateRoomSettings(ctx context.Context, req *room.UpdateRoomRequest) (*room.UpdateRoomResponse, error) {
	rm := &model.Room{
		RoomName: req.RoomName,
		Users:    req.GetUserIds(),
	}

	err := h.service.UpdateRoom(ctx, rm)

	if err != nil {
		return nil, err
	}
	var res room.UpdateRoomResponse
	res.Success = true
	return &res, nil

}
