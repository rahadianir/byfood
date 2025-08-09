package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
)

type BookLogic struct {
	deps *core.Dependency
	repo RepositoryInterface
}

func NewBookLogic(deps *core.Dependency, repo RepositoryInterface) *BookLogic {
	return &BookLogic{
		deps: deps,
		repo: repo,
	}
}

func (repo *BookLogic) GetBooks(ctx context.Context, params model.BookSearchParams, page pagination.Page) ([]model.Book, pagination.Metadata, error) {
	return nil, pagination.Metadata{}, nil
}
func (repo *BookLogic) GetBookByID(ctx context.Context, id int64) (model.Book, error) {
	return model.Book{}, nil
}
func (repo *BookLogic) UpdateBook(ctx context.Context, data model.Book) (model.Book, error) {
	return model.Book{}, nil
}
func (repo *BookLogic) DeleteBook(ctx context.Context, id int64) error {
	return nil
}
