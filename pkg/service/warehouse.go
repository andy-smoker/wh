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

func (s *WHservice) GetItem(itemID int) (structs.WHitem, error) {
	return s.repo.GetItem(itemID)
}

func (s *WHservice) UpdateItem(item structs.WHitem) (structs.WHitem, error) {
	return s.repo.UpdateItem(item)
}

func (s *WHservice) DeleteItem(id int, itemType string) error {
	return s.repo.DeleteItem(id, itemType)
}

func (s *WHservice) GetItemsList(itemsType string) ([]interface{}, error) {
	return s.repo.GetItemsList(itemsType)
}
