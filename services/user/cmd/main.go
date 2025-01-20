package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/config"
	grpcHandler "github.com/sri2103/chat_me/services/user/handlers/grpc"
	"github.com/sri2103/chat_me/services/user/repository"
	"github.com/sri2103/chat_me/services/user/repository/memory"
	"github.com/sri2103/chat_me/services/user/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting user service")
	port := flag.Int("port", 8081, "port for service")
	sqlitePath := flag.String("sqlite", "", "sqlite Databases path")
	env := flag.String("env", "dev", "Environment for the service")
	flag.Parse()
	cfg := &config.Config{
		Port:        *port,
		SqlitePath:  *sqlitePath,
		Environment: *env,
	}
	err := SetUpGrpCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func SetUpGrpCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}
	repo := memory.New()
	service := service.New(repo)
	AuthRepo := repository.AuthRepositoryFaactory(cfg)
	s := grpc.NewServer()
	h := grpcHandler.NewGrpcServerHandler(service)
	authHandler := grpcHandler.NewGrpcAuthHandler(AuthRepo)
	pb.RegisterUserServiceServer(s, h)
	pb.RegisterAuthServiceServer(s, authHandler)
	err = s.Serve(lis)
	if err != nil {
		return nil
	}
	return nil
}
