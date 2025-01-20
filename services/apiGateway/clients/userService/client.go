package userClient

import (
	"github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectUserClient(cfg *config.Config) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		cfg.UserClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	cfg.UserClientService = user.NewUserServiceClient(conn)
	cfg.AuthClientService = user.NewAuthServiceClient(conn)
	return conn, nil
}
