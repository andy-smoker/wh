package postgres

import (
	"database/sql"
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
		colums   [2]string
		args     []interface{}
		itemType string
	)
	if strg := &item.Item.Strorage; strg != nil {
		itemType = "storage"
		colums[0] = "title, volume, type, size"
		colums[1] = "$1,$2,$3,$4"
		args = append(args, strg.Title, strg.Volume, strg.Type, strg.Size)
	} else if &item.Item.Monitor != nil {
		fmt.Println(item.Item.Monitor)
	} else {
		return 0, errors.New("invalid body")
	}
	query := fmt.Sprintf("insert into %s (%s) values(%s) returning id", storageTable, colums[0], colums[1])
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
		columns string
		table   string
	)
	query := fmt.Sprintf("select item_id, items_type from %s where id=$1", itemTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&item.ItemID, &item.ItemsType); err != nil {
		return item, err
	}
	item.ID = id
	switch item.ItemsType {
	case "storage":
		columns = "title, volume, size, type"
		table = storageTable
	}
	query = fmt.Sprintf("select %s from %s where id=$1", columns, table)
	err := r.db.Get(&item.Item.Strorage, query, item.ItemID)
	return item, err
}

func (r *WHPostgres) UpdateItem(item structs.WHitem) (structs.WHitem, error) {
	var (
		query   string
		columns string
		args    []interface{}
		table   string
	)
	query = fmt.Sprintf("update %s set in_stock=$1 where id=$2 returning item_id", itemTable)
	row := r.db.QueryRow(query, item.InStock, item.ID)
	if err := row.Scan(&item.ItemID); err != nil {
		return item, err
	}
	switch item.ItemsType {
	case "storage":
		tmp := item.Item.Strorage
		columns = "title=$2, volume=$3, size=$4, type=$5"
		args = append(args, item.ItemID, tmp.Title, tmp.Volume, tmp.Size, tmp.Type)
		table = storageTable
	}
	query = fmt.Sprintf("update %s set %s where id=$1", table, columns)
	fmt.Println(args)
	err := r.db.Get(&item.Item.Strorage, query, args...)
	if err == sql.ErrNoRows {
		return item, nil
	}
	return item, err
}
