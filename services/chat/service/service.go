package service

import "context"

type Service interface {
	SendMessage(ctx context.Context) error
	GetMessage(ctx context.Context) error
}

type Repository interface {
	SendMessage(ctx context.Context) error
	GetMessage(ctx context.Context) error
}
