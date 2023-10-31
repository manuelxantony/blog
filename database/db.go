package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/manuelxantony/blog/config"
)

func OpenDB(config *config.Configuration) (*sql.DB, error) {
	cfg := mysql.Config{
		User:   config.Database.DBUser,
		Passwd: config.Database.DBPassword,
		Net:    config.Database.Network,
		Addr:   config.Database.Addr,
		DBName: config.Database.DBName,
	}

	var err error

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("DB connection created succesfully")

	return db, nil
}
