package postgres

import (
	"fmt"

	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/jmoiron/sqlx"
)

func NewAuth(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structs.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (login, username, hash_pass) values($1,$2,$3) returning id", userTable)
	row := r.db.QueryRow(query, user.Login, user.Username, user.Pass)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (structs.User, error) {
	user := structs.User{}
	query := fmt.Sprintf("select id from %s where login=$1 and hash_pass=$2", userTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}
