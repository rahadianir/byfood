package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestBookRepo_GetBooks(t *testing.T) {
	type fields struct {
		deps *core.Dependency
	}
	type args struct {
		ctx    context.Context
		params model.BookSearchParams
		page   pagination.Page
	}

	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")

	mockFields := fields{
		deps: &core.Dependency{
			Logger: slog.Default(),
			DB:     dbx,
		},
	}

	now := time.Now()

	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []model.Book
		want1    pagination.Metadata
		wantErr  bool
		mockFunc func()
	}{
		{
			name:   "success get books with search params",
			fields: mockFields,
			args: args{
				ctx: context.Background(),
				params: model.BookSearchParams{
					Search: "oda",
				},
				page: pagination.Page{
					Page: 1,
					Size: 10,
				},
			},
			want: []model.Book{
				{
					ID:          int64(1),
					Title:       "One Piece",
					Author:      "Eiichiro Oda",
					PublishYear: 1997,
					BaseAudit: model.BaseAudit{
						CreatedAt: &now,
						UpdatedAt: &now,
					},
				},
			},
			want1: pagination.Metadata{
				CurrentPage:  1,
				PageSize:     10,
				FirstPage:    1,
				LastPage:     1,
				TotalRecords: 1,
			},
			wantErr: false,
			mockFunc: func() {
				// q := `SELECT id, title, author, publish_year, created_at, updated_at FROM library.books WHERE (title ILIKE $1 OR author ILIKE $2) AND deleted_at IS NULL ORDER BY id`
				expectedRows := sqlmock.NewRows([]string{"id", "title", "author", "publish_year", "created_at", "updated_at"})
				expectedRows.AddRow(1, "One Piece", "Eiichiro Oda", 1997, now, now)
				mockDB.ExpectQuery(`(?s)^.*$`).WithArgs("%oda%", "%oda%", 10, 0).WillReturnRows(expectedRows)
				mockDB.ExpectQuery(`(?s)^.*$`).WithArgs("%oda%", "%oda%").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &BookRepo{
				deps: tt.fields.deps,
			}

			tt.mockFunc()

			got, got1, err := repo.GetBooks(tt.args.ctx, tt.args.params, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookRepo.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookRepo.GetBooks() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BookRepo.GetBooks() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
