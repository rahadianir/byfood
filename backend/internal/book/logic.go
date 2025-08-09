package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"byfood-app/internal/pkg/xerrors"
	"context"
	"errors"
	"fmt"
	"log/slog"
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

func (logic *BookLogic) GetBooks(ctx context.Context, params model.BookSearchParams, page pagination.Page) ([]model.Book, pagination.Metadata, error) {
	data, meta, err := logic.repo.GetBooks(ctx, params, page)
	if err != nil {
		if errors.Is(err, xerrors.ErrDataNotFound) {
			return []model.Book{}, meta, err
		}

		logic.deps.Logger.ErrorContext(ctx, "failed to get books", slog.Any("error", err))
		return []model.Book{}, meta, err
	}
	return data, meta, nil
}

func (logic *BookLogic) GetBookByID(ctx context.Context, id int64) (model.Book, error) {
	if id <= 0 {
		return model.Book{}, xerrors.NewClientError(xerrors.ErrInvalidID)
	}

	data, err := logic.repo.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, xerrors.ErrDataNotFound) {
			return model.Book{}, err
		}

		logic.deps.Logger.ErrorContext(ctx, "failed to get book by id", slog.Any("error", err))
		return model.Book{}, err
	}

	return data, nil
}

func (logic *BookLogic) StoreBook(ctx context.Context, data model.Book) (model.Book, error) {
	switch {
	case data.Author == "":
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("author field must not empty"))
	case data.Title == "":
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("title field must not empty"))
	case data.PublishYear <= 0:
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("publish year field must not empty and greater than 0"))
	}

	result, err := logic.repo.StoreBook(ctx, data)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to store book data", slog.Any("error", err))
		return model.Book{}, err
	}

	return result, nil
}

func (logic *BookLogic) UpdateBook(ctx context.Context, data model.Book) (model.Book, error) {
	switch {
	case data.ID <= 0:
		return model.Book{}, xerrors.NewClientError(xerrors.ErrInvalidID)
	case data.Author == "":
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("author field must not empty"))
	case data.Title == "":
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("title field must not empty"))
	case data.PublishYear <= 0:
		return model.Book{}, xerrors.NewClientError(fmt.Errorf("publish year field must not empty and greater than 0"))
	}

	result, err := logic.repo.UpdateBook(ctx, data)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to update book data", slog.Any("error", err))
		return result, err
	}

	return result, nil
}

func (logic *BookLogic) DeleteBook(ctx context.Context, id int64) error {
	if id <= 0 {
		return xerrors.NewClientError(xerrors.ErrInvalidID)
	}

	err := logic.repo.DeleteBook(ctx, id)
	if err != nil {
		logic.deps.Logger.ErrorContext(ctx, "failed to delete book data", slog.Any("error", err))
		return err
	}
	return nil
}
