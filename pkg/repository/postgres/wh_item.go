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
		columns [2]string
		args    []interface{}
	)
	if strg := &item.Item.Strorage; strg != nil {
		item.ItemsType = "storage"
		columns[0] = "title, volume, type, size"
		columns[1] = "$1,$2,$3,$4"
		args = append(args, strg.Title, strg.Volume, strg.Type, strg.Size)
	} else if &item.Item.Monitor != nil {
		fmt.Println(item.Item.Monitor)
	} else {
		return 0, errors.New("invalid body")
	}
	query := fmt.Sprintf("insert into %s (%s) values(%s) returning id", storageTable, columns[0], columns[1])
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&item.ItemID); err != nil {
		tx.Rollback()
		return 0, err
	}
	_, err = tx.Exec(fmt.Sprintf("insert into %s (item_id, items_type, in_stock) values($1,$2,$3)", itemTable),
		item.ItemID, item.ItemsType, true)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return item.ItemID, tx.Commit()
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

func (r *WHPostgres) GetItemsList(itemsType string) ([]interface{}, error) {
	var (
		items   []interface{}
		table   string
		columns string
		args    []interface{}
		item    func() (interface{}, []interface{})
	)
	switch itemsType {
	case "storage":
		columns = "title, volume, type, size"
		table = storageTable
		item = func() (pointer interface{}, dest []interface{}) {
			tmp := new(structs.Strorage)

			dest = append(dest, &tmp.Title, &tmp.Volume, &tmp.Type, &tmp.Size)
			pointer = tmp
			return
		}
	default:
		return items, errors.New("invalid type")
	}
	query := fmt.Sprintf("select %s from %s", columns, table)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return items, err
	}
	for rows.Next() {
		pointer, dest := item()
		rows.Scan(dest...)
		items = append(items, pointer)
	}
	return items, nil

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
	default:
		return item, nil
	}
	query = fmt.Sprintf("update %s set %s where id=$1", table, columns)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *WHPostgres) DeleteItem(id int, itemsType string) error {
	tx, err := r.db.Begin()
	var (
		itemID int
		query  string
		table  string
	)
	query = fmt.Sprintf("delete from %s where id=$1 returning item_id", itemTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&itemID); err != nil {
		tx.Rollback()
		return err
	}
	switch itemsType {
	case "storage":
		table = storageTable
	}
	query = fmt.Sprintf("delete from %s where id=$1", table)
	_, err = r.db.Exec(query, itemID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
