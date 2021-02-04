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

type Warehouse interface {
	CreateItem(item structs.WHitem) (int, error)
	GetItem(id int) (structs.WHitem, error)
	UpdateItem(item structs.WHitem) (structs.WHitem, error)
	DeleteItem(id int, itemType string) error
}

// сруктура репозитория
type Repository struct {
	Authorization
	Warehouse
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
		Warehouse:     postgres.NewWH(db),
	}
}
