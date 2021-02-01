package service

import (
	"github.com/andy-smoker/wh-server/pkg/repository"
	"github.com/andy-smoker/wh-server/pkg/structs"
)

type WHservice struct {
	repo repository.Warehouse
}

func NewWHService(repo repository.Warehouse) *WHservice {
	return &WHservice{repo: repo}
}

func (s *WHservice) CreateItem(item structs.WHitem) (int, error) {
	return s.repo.CreateItem(item)
}
