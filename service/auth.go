package service

import (
	"crypto/sha1"
	"fmt"

	server "github.com/andy-smoker/wh-server"
	"github.com/andy-smoker/wh-server/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "sgdfhemrlg;dv[gvs"
)

type tokenClaims struct {
	jwt.StandardClaims
	IserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user server.User) (int, error) {
	user.Pass = generatePasswordHash(user.Pass)
	return s.repo.CreateUser(user)
}

//func (s *AuthService) GenerateToken(username, password string) (string, error) {}

func generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
