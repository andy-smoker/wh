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

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	{
		return &Service{
			Authorization: NewAuthService(repo),
		}
	}
}
