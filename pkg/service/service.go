package service

import (
	"github.com/andy-smoker/wh-server/pkg/repository"
	"github.com/andy-smoker/wh-server/pkg/structs"
)

type Authorization interface {
	CreateUser(user structs.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Warehouse interface {
	CreateItem(item structs.WHitem) (int, error)
	GetItem(itemID int) (structs.WHitem, error)
	UpdateItem(item structs.WHitem) (structs.WHitem, error)
	DeleteItem(id int, itemType string) error
	GetItemsList(itemsType string) ([]interface{}, error)
}

type Service struct {
	Authorization
	Warehouse
}

func NewService(repo *repository.Repository) *Service {
	{
		return &Service{
			Authorization: NewAuthService(repo),
			Warehouse:     NewWHService(repo),
		}
	}
}
