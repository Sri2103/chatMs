package main

import (
	"fmt"
	"log"

	userClient "github.com/sri2103/chat_me/services/apiGateway/clients/userService"
	"github.com/sri2103/chat_me/services/apiGateway/config"
	"go.uber.org/zap"
)

var userClientUrl = "localhost:8081"
var gatewayPort = ":7080"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	fmt.Println("gateway")
	cfg := &config.Config{
		UserClient: userClientUrl,
		Port:       gatewayPort,
	}

	conn, err := userClient.ConnectUserClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	select {}
}
