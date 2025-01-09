package config

import (
	userpb "github.com/sri2103/chat_me/protos/user"
)

type Config struct {
	UserClient        string
	UserClientService userpb.UserServiceClient
}
