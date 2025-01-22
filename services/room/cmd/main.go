package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sri2103/chat_me/protos/room"
	"github.com/sri2103/chat_me/services/room/config"
	handler "github.com/sri2103/chat_me/services/room/handlers/grpc"
	sqlRepo "github.com/sri2103/chat_me/services/room/repository/sql"
	"github.com/sri2103/chat_me/services/room/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting room service")
	port := flag.Int("port", 8082, "port for the room service")
	env := flag.String("env", "dev", "environment for the service")
	flag.Parse()
	cfg := &config.Config{
		Port:        *port,
		Environment: *env,
	}
	err := SetupGRPCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func SetupGRPCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}

	repo := sqlRepo.SetUpRepo(cfg)
	svc := service.New(repo)
	s := grpc.NewServer()
	h := handler.NewGrpcHandler(svc)
	room.RegisterRoomServiceServer(s, h)
	err = s.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
