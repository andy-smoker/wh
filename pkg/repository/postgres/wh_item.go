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

// CreateItem -  create new item
func (r *WHPostgres) CreateItem(item structs.WHitem) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}
	var (
		props []pgQuery
		args  []interface{}
	)

	itemProps := &item.ItemProps
	if item.ItemsType == "storage" {
		props = append(props, pgQuery{
			table:     storageTable,
			columns:   "title, volume, type, size",
			values:    "$1,$2,$3,$4",
			returning: "id",
		})
		args = append(args, itemProps.Title, itemProps.Volume, itemProps.Type, itemProps.Size)

	} else if &item.ItemProps.Monitor != nil {
		fmt.Println(item.ItemProps.Monitor)
	} else {
		tx.Rollback()
		return 0, errors.New("invalid body")
	}
	query := makeQuery(props, "insert")
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(&item.ItemProps.ID); err != nil {
		tx.Rollback()
		return 0, err
	}
	props[0] = pgQuery{
		table:     itemTable,
		columns:   "item_id, items_type, in_stock",
		values:    "$1,$2,$3",
		returning: "id",
	}
	query = makeQuery(props, "insert")
	res, err := tx.Exec(query, item.ItemProps.ID, item.ItemsType, true)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, _ := res.RowsAffected()
	return int(id), tx.Commit()
}

// GetItem - get item by ID
func (r *WHPostgres) GetItem(id int) (structs.WHitem, error) {
	dest := struct {
		ID      int    `db:"item_id"`
		Type    string `db:"items_type"`
		InStock bool   `db:"in_stock"`
	}{}
	query := fmt.Sprintf("select item_id, items_type, in_stock from %s where id=$1", itemTable)
	err := r.db.Get(&dest, query, id)
	item := structs.WHitem{
		ID:        id,
		ItemsType: dest.Type,
		InStock:   dest.InStock,
		ItemProps: structs.WHitemProps{
			ID: dest.ID,
		},
	}

	if err != nil {
		return item, err
	}

	var (
		columns string
		table   string
		pgS     []pgQuery
	)
	pgS = append(pgS, pgQuery{
		columns: "i.id",
		table:   "items as i",
		sample:  "i.id=$1",
	})
	itemProps := &item.ItemProps

	switch item.ItemsType {
	case "storage":
		pgS = append(pgS, pgQuery{
			columns: "s.id, s.volume, s.type, s.size",
			table:   "join storage as s",
			on:      "on s.id = i.item_id",
			sample:  `i.items_type = 'storage'`,
		})
		columns = "title, volume, size, type"
		table = storageTable
	default:
		return item, errors.New("invalid type")
	}
	query = fmt.Sprintf("select %s from %s where id=$1", columns, table)
	err = r.db.Get(itemProps, query, item.ItemProps.ID)
	return item, err
}

// GetItemsList - get item list by filter
func (r *WHPostgres) GetItemsList(filter string) ([]interface{}, error) {
	var (
		items []interface{}
		query string
		joins []pgQuery
		item  func() (*structs.WHitem, []interface{})
	)

	joins = append(joins, pgQuery{
		columns: "i.id, i.in_stock",
		table:   "items as i ",
		sample:  `i.items_type='storage'`,
		orderBy: "i.id",
	})

	switch filter {
	case "storage":
		joins = append(joins, pgQuery{
			columns: ", s.title, s.volume, s.type, s.size ",
			table:   fmt.Sprintf("join %s as s ", storageTable),
			on:      "on s.id = i.item_id ",
		})
		item = func() (pointer *structs.WHitem, dest []interface{}) {
			i := structs.WHitem{}
			i.ItemsType = filter
			dest = append(dest, &i.ID, &i.InStock, &i.ItemProps.Title, &i.ItemProps.Volume, &i.ItemProps.Type, &i.ItemProps.Size)
			pointer = &i
			return
		}
	case "all":
		joins[0] = pgQuery{
			columns: ` a.id, a.items_type, a.title, a.in_stock`,
			table: `(select i.id,  s.title, i.items_type, i.in_stock
			from items as i, storages as s 
			where s.id = i.item_id and i.items_type = 'storage'
			union select i.id,  m.title, i.items_type, i.in_stock
			from items as i, monitors as m 
			where m.id = i.item_id and i.items_type = 'monitor'
			) a`,
			orderBy: `a.id`,
		}
		item = func() (pointer *structs.WHitem, dest []interface{}) {
			i := structs.WHitem{}

			dest = append(dest, &i.ID, &i.ItemsType, &i.ItemProps.Title, &i.InStock)
			pointer = &i
			return
		}
	default:
		return items, errors.New("invalid type")
	}
	query = makeQuery(joins, "select")
	rows, err := r.db.Query(query)
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
		query string
		props []pgQuery
		args  []interface{}
	)
	query = fmt.Sprintf("update %s set in_stock=$1 where id=$2 returning item_id", itemTable)
	row := r.db.QueryRow(query, item.InStock, item.ID)
	if err := row.Scan(&item.ItemProps.ID); err != nil {
		return item, err
	}
	switch item.ItemsType {
	case "storage":
		tmp := item.ItemProps
		props = append(props, pgQuery{
			table: storageTable,
			set:   "title=$2, volume=$3, size=$4, type=$5",
		})
		args = append(args, item.ItemProps.ID, tmp.Title, tmp.Volume, tmp.Size, tmp.Type)
	default:
		return item, nil
	}
	props[0].sample = "id=$1"
	query = makeQuery(props, "update")
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
