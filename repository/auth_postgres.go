package repository

import (
	"fmt"

	server "github.com/andy-smoker/wh-server"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user server.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(name, username, hash_pass) values($1,$2,$3) return id", userTable)
	row := r.db.QueryRow(query, user.Login, user.Username, user.Pass)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
