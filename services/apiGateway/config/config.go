package config

import (
	userpb "github.com/sri2103/chat_me/protos/user"
	"go.uber.org/zap"
)

type Config struct {
	Log               *zap.Logger
	Port              string
	UserClient        string
	UserClientService userpb.UserServiceClient
}
