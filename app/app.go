package app

import (
	"database/sql"

	"github.com/manuelxantony/blog/config"
)

type App struct {
	Config *config.Configuration
	DB     *sql.DB
}
