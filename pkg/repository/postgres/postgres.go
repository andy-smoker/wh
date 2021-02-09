package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	// название таблиц в базе данных
	userTable    = "users"
	storageTable = "storages"
	itemTable    = "items"
)

type AuthPostgres struct {
	db *sqlx.DB
}

type WHPostgres struct {
	db *sqlx.DB
}

type PostgresCFG struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

//go:generate mockgen -source=postgres.go -destination mocks/postgres.go

// создаём новое подключение к postgresql
func NewDB(cfg PostgresCFG) (*sqlx.DB, error) {
	// connect to postgresdb
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
