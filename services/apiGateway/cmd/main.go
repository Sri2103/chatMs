package main

import (
	"context"
	"fmt"
	"log"

	roomClient "github.com/sri2103/chat_me/services/apiGateway/clients/roomService"
	userClient "github.com/sri2103/chat_me/services/apiGateway/clients/userService"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	httpServer "github.com/sri2103/chat_me/services/apiGateway/servers/http"
	"go.uber.org/zap"
)

var userClientUrl = "localhost:8081"
var gatewayPort = ":7080"
var roomClientUrl = "localhost:8082"

func main() {
	logger, _ := zap.NewProduction()
	ctx := context.Background()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Println("Error in zap")
		}

	}()
	fmt.Println("gateway")
	cfg := &config.Config{
		UserClient: userClientUrl,
		Port:       gatewayPort,
		RoomClient: roomClientUrl,
	}

	conn, err := userClient.ConnectUserClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	roomSvcConn, err := roomClient.ConnectRoomClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer roomSvcConn.Close()
	go httpServer.SetUpServer(ctx, cfg)

	select {}
}
