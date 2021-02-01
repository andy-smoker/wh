package repository

import (
	"github.com/andy-smoker/wh-server/pkg/repository/postgres"
	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/jmoiron/sqlx"
)

// интерфейс авторизаци пользователя в программе
type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GetUser(username, password string) (structs.User, error)
}

// сруктура репозитория
type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
	}
}
