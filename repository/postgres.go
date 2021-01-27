package repository

import (
	"fmt"

	server "github.com/andy-smoker/wh-server"
	"github.com/jmoiron/sqlx"
)

const (
	// название таблиц в базе данных
	userTable = "users"
)

// создаём новое подключение к postgresql
func NewPostgresDB(cfg server.Config) (*sqlx.DB, error) {
	// connect to postgresdb
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.PGcfg.Host, cfg.PGcfg.Port, cfg.PGcfg.DBName, cfg.PGcfg.Username, cfg.PGcfg.Password, cfg.PGcfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
