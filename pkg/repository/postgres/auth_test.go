package postgres

import (
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"

	"github.com/andy-smoker/wh-server/pkg/structs"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
	r := NewAuth(sqlx.NewDb(db, "postgres"))
	testTable := []struct {
		name    string
		mock    func()
		input   structs.User
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into user").
					WithArgs("test", "test", "pass").WillReturnRows(rows)
			},
			input: structs.User{
				Login:    "test",
				Username: "test",
				Pass:     "pass",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert inti users").
					WithArgs("test", "test", "").WillReturnRows(rows)
			},
			input: structs.User{
				Login:    "test",
				Username: "test",
				Pass:     "",
			},
			wantErr: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.CreateUser(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
	type args struct {
		login string
		pass  string
	}
	r := NewAuth(sqlx.NewDb(db, "postgres"))
	testTable := []struct {
		name    string
		mock    func()
		input   args
		want    structs.User
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "login", "username", "pass"}).
					AddRow(1, "test", "test", "pass")
				mock.ExpectQuery("select (.+) from user").
					WithArgs("test", "pass").WillReturnRows(rows)
			},
			input: args{"test", "pass"},
			want: structs.User{
				ID:       1,
				Login:    "test",
				Username: "test",
				Pass:     "pass",
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "login", "username", "pass"})
				mock.ExpectQuery("select (.+) from user").
					WithArgs("test", "pass").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.GetUser(test.input.login, test.input.pass)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
