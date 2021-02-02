package postgres

import (
	"errors"
	"fmt"

	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/jmoiron/sqlx"
)

func NewWH(db *sqlx.DB) *WHPostgres {
	return &WHPostgres{db: db}
}

func (r *WHPostgres) CreateItem(item structs.WHitem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var (
		id       int
		query    string
		args     []interface{}
		itemType string
	)
	if strg := &item.Item.Strorage; strg != nil {
		itemType = "storage"
		fmt.Println(item.Item.Strorage)
		query = fmt.Sprintf("insert into %s (title, volume, type, size) values($1,$2,$3,$4) returning id", storageTable)
		args = append(args, strg.Title, strg.Volume, strg.Type, strg.Size)
	} else if &item.Item.Monitor != nil {
		fmt.Println(item.Item.Monitor)
	} else {
		return 0, errors.New("invalid body")
	}

	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	_, err = tx.Exec(fmt.Sprintf("insert into %s (item_id, items_type, in_stock) values($1,$2,$3)", itemTable),
		id, itemType, true)
	if err != nil {
		return 0, err
	}
	return id, tx.Commit()
}

func (r *WHPostgres) GetItem(id int) (structs.WHitem, error) {
	item := structs.WHitem{}
	var (
		itemsType string
		itemsID   int
	)
	query := fmt.Sprintf("select item_id, items_type from %s where id=$1", itemTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&itemsID, &itemsType); err != nil {
		return item, err
	}
	item.ID = id
	switch itemsType {
	case "storage":
		query := fmt.Sprintf("select title, volume, size, type from %s where id=$1", storageTable)
		err := r.db.Get(&item.Item.Strorage, query, itemsID)
		return item, err
	}
	return item, errors.New("invalid")
}
