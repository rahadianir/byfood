package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
)

type BookRepo struct {
	deps *core.Dependency
}

func NewSQLRepo(deps *core.Dependency) *BookRepo {
	return &BookRepo{
		deps: deps,
	}
}

func (repo *BookRepo) GetBooks(ctx context.Context, params model.BookSearchParams, page pagination.Page) ([]model.Book, pagination.Metadata, error) {
	return nil, pagination.Metadata{}, nil
}
func (repo *BookRepo) GetBookByID(ctx context.Context, id int64) (model.Book, error) {
	return model.Book{}, nil
}
func (repo *BookRepo) UpdateBook(ctx context.Context, data model.Book) (model.Book, error) {
	return model.Book{}, nil
}
func (repo *BookRepo) DeleteBook(ctx context.Context, id int64) error {
	return nil
}
