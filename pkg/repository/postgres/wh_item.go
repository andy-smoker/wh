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
	var (
		id    int
		query string
		args  []interface{}
	)
	if strg := &item.Item.Strorage; strg != nil {
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
		return 0, err
	}
	return id, nil
}
