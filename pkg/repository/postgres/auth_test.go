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
	xdb := sqlx.NewDb(db, "postgres")
	r := NewAuth(xdb)
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
