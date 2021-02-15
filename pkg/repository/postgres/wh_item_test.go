package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/andy-smoker/wh-server/pkg/structs"
	"github.com/jmoiron/sqlx"
)

func TestCreateItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
	r := NewWH(sqlx.NewDb(db, "postgres"))
	testTable := []struct {
		name    string
		mock    func()
		input   structs.WHitem
		want    int
		wantErr bool
	}{
		{
			name: "storage_OK",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into storages").
					WithArgs("test", 200, "HDD", "2.5").WillReturnRows(rows)

				mock.ExpectExec("insert into items").
					WithArgs(1, "storage", true).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

			},
			input: structs.WHitem{
				ID:        1,
				ItemsType: "storage",
				ItemProps: structs.WHitemProps{
					ID:    1,
					Title: "test",
					Strorage: structs.Strorage{
						Size:   "2.5",
						Volume: 200,
						Type:   "HDD",
					},
					Vendor: "test",
				},
				InStock: true,
			},
			want: 1,
		},
		{
			name: "Empty Type",
			mock: func() {},
			input: structs.WHitem{
				ID:        1,
				ItemsType: "",
				ItemProps: structs.WHitemProps{
					ID:    1,
					Title: "test",
					Strorage: structs.Strorage{
						Size:   "2.5",
						Volume: 200,
						Type:   "HDD",
					},
					Vendor: "test",
				},
				InStock: true,
			},
			wantErr: true,
		},
		{
			name: "Empty Fields",
			mock: func() {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("insert into storages").
					WithArgs("test", 200, "HDD", "2.5").WillReturnRows(rows)

				mock.ExpectExec("insert into items").
					WithArgs(1, "storage", true).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()

			},
			input: structs.WHitem{
				ID:        1,
				ItemsType: "storage",
				ItemProps: structs.WHitemProps{
					ID:    1,
					Title: "",
					Strorage: structs.Strorage{
						Size:   "",
						Volume: 200,
						Type:   "HDD",
					},
					Vendor: "test",
				},
				InStock: true,
			},
			wantErr: true,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := r.CreateItem(test.input)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("db error: %s", err)
	}
	r := NewWH(sqlx.NewDb(db, "postgres"))
	testTable := []struct {
		name    string
		mock    func()
		input   int
		want    structs.WHitem
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				itemRows := sqlmock.NewRows([]string{"item_id", "items_type", "in_stock"}).
					AddRow(1, "storage", true)
				mock.ExpectQuery("select (.+) from items").WithArgs(1).WillReturnRows(itemRows)
				strgRows := sqlmock.NewRows([]string{"id", "title", "size", "volume", "type"}).
					AddRow(1, "test", "2.5", 200, "HDD")
				mock.ExpectQuery("select (.+) from storages").WithArgs(1).WillReturnRows(strgRows)
			},
			input: 1,
			want: structs.WHitem{
				ID:        1,
				ItemsType: "storage",
				ItemProps: structs.WHitemProps{
					ID:    1,
					Title: "test",
					Strorage: structs.Strorage{
						Size:   "2.5",
						Volume: 200,
						Type:   "HDD",
					},
				},
				InStock: true,
			},
		},
		{
			name:    "Invalid Input",
			mock:    func() {},
			input:   0,
			wantErr: true,
		},
	}
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := r.GetItem(test.input)
			if test.wantErr {
				t.Log(err)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
