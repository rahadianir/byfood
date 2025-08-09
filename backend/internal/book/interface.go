package book

import (
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"context"
)

type RepositoryInterface interface {
	GetBooks(ctx context.Context, params model.BookSearchParams, page pagination.Page) ([]model.Book, pagination.Metadata, error)
	GetBookByID(ctx context.Context, id int64) (model.Book, error)
	UpdateBook(ctx context.Context, data model.Book) (model.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type LogicInterface interface {
	GetBooks(ctx context.Context, params model.BookSearchParams, page pagination.Page) ([]model.Book, pagination.Metadata, error)
	GetBookByID(ctx context.Context, id int64) (model.Book, error)
	UpdateBook(ctx context.Context, data model.Book) (model.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}
