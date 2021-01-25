package service

import (
	server "github.com/andy-smoker/wh-server"
	"github.com/andy-smoker/wh-server/repository"
)

type Authorization interface {
	CreateUser(user server.User) (int, error)
	//GenerateToken(username, password string) (string, error)
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
