package repository

import (
	server "github.com/andy-smoker/wh-server"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user server.User) (int, error)
	//GetUser(username, password string) (server.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
