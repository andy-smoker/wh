package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

const (
	// название таблиц в базе данных
	userTable    = "users"
	storageTable = "storages"
	itemTable    = "items"
)

type AuthPostgres struct {
	db *sqlx.DB
}

type WHPostgres struct {
	db *sqlx.DB
}

type PostgresCFG struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	SSLMode  string `toml:"sslmode"`
}

// создаём новое подключение к postgresql
func NewDB(cfg PostgresCFG) (*sqlx.DB, error) {
	// connect to postgresdb
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type pgQuery struct {
	columns   string
	table     string
	sample    string
	values    string
	set       string
	on        string
	orderBy   string
	returning string
}

func makeQuery(s []pgQuery, method string) string {
	var (
		columns   string
		tables    string
		values    string
		set       string
		sample    string
		orderBy   string
		returning string
		query     string
	)

	switch method {
	case "select":
		for _, join := range s {
			columns += join.columns
			tables += fmt.Sprintf("%s %s", join.table, join.on)
			sample += (join.sample)
			orderBy += (join.orderBy)
		}
		query = fmt.Sprintf("select %s from %s", columns, tables)
	case "update":
		for _, join := range s {
			tables += fmt.Sprintf("%s", join.table)
			sample += (join.sample)
			set += (join.set)
			returning += (join.returning)
		}
		query = fmt.Sprintf("update %s set %s", tables, set)
	case "insert":
		for _, join := range s {
			tables += fmt.Sprintf("%s(%s)", join.table, join.columns)
			values += (join.values)
			returning += (join.returning)
		}
		query = fmt.Sprintf("insert into %s values(%s) returning %s", tables, values, returning)
	case "delete":
		for _, join := range s {
			tables += fmt.Sprintf("%s", join.table)
			returning += (join.returning)
		}
		query = fmt.Sprintf("delete from %s", tables)
	}
	if sample != "" {
		query += fmt.Sprintf(" where %s", sample)
	}
	if orderBy != "" {
		query += fmt.Sprintf(" order by %s", orderBy)
	}
	return query
}
