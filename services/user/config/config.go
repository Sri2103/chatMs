package config

import "database/sql"

type Config struct {
	Port        int
	Environment string
	Db          *sql.DB
}
