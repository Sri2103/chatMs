package main

import (
	"fmt"
	"log"

	userClient "github.com/sri2103/chat_me/services/apiGateway/clients/userService"
	"github.com/sri2103/chat_me/services/apiGateway/config"
)

var userClientUrl = "localhost:8081"
var gatewayPort = ":7080"

func main() {
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
