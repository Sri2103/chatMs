package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	database "github.com/sri2103/chat_me/DB"
	pb "github.com/sri2103/chat_me/protos/user"
	"github.com/sri2103/chat_me/services/user/config"
	grpcHandler "github.com/sri2103/chat_me/services/user/handlers/grpc"
	"github.com/sri2103/chat_me/services/user/repository"
	"github.com/sri2103/chat_me/services/user/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting user service")
	port := flag.Int("port", 8081, "port for service")
	env := flag.String("env", "dev", "Environment for the service")
	flag.Parse()

	cfg := &config.Config{
		Port:        *port,
		Environment: *env,
	}
	db, err := OpenDBPath(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = SetUpGrpCServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func SetUpGrpCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}

	repo := repository.RepositoryFactory(cfg)
	service := service.New(repo)
	AuthRepo := repository.AuthRepositoryFactory(cfg)
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

func OpenDBPath(cfg *config.Config) (*sql.DB, error) {
	db, err := database.NewPostgresConnection(database.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "harsha",
		Password: "password",
		DBName:   "main",
	})

	if err != nil {
		return nil, err
	}

	cfg.Db = db

	return db, nil

}
