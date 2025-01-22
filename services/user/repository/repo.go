package repository

import (
	"github.com/sri2103/chat_me/services/user/config"
	sqlRepo "github.com/sri2103/chat_me/services/user/repository/sql"
	"github.com/sri2103/chat_me/services/user/service"
)

func RepositoryFactory(cfg *config.Config) service.RepoInterface {
	switch cfg.Environment {
	case "dev":
		return sqlRepo.NewPostgresRepo(cfg)
	case "prod":
		return nil
	default:
		return nil
	}
}

func AuthRepositoryFactory(cfg *config.Config) service.AuthServiceRepo {
	switch cfg.Environment {
	case "dev":
		return sqlRepo.NewAuthRepo(cfg)
	default:
		return nil
	}
}
