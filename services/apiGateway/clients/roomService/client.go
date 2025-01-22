package roomClient

import (
	"github.com/sri2103/chat_me/protos/room"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectRoomClient(cfg *config.Config) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		cfg.RoomClient,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	cfg.RoomClientService = room.NewRoomServiceClient(conn)
	return conn, nil
}
