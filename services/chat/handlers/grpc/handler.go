package handler

import (
	"context"

	"github.com/sri2103/chat_me/protos/chat"
	"github.com/sri2103/chat_me/services/chat/service"
)

type Handler struct {
	service service.Service
	chat.UnimplementedChatServiceServer
}

func New(svc service.Service) chat.ChatServiceServer {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) SendMessage(_ context.Context, _ *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (h *Handler) GetMessages(_ context.Context, _ *chat.GetMessagesRequest) (*chat.GetMessagesResponse, error) {
	panic("not implemented") // TODO: Implement
}
