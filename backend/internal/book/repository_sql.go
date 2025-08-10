package book

import (
	"byfood-app/internal/core"
	"byfood-app/internal/model"
	"byfood-app/internal/pkg/pagination"
	"byfood-app/internal/pkg/xerrors"
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/huandu/go-sqlbuilder"
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
	var (
		result []model.Book
		meta   pagination.Metadata
	)

	// base query
	countQ := sqlbuilder.NewSelectBuilder()
	countQ.Select("COUNT(1)").From("library.books")

	q := sqlbuilder.NewSelectBuilder()
	q = q.Select("id", "title", "author", "publish_year", "created_at", "updated_at").From("library.books")

	if params.Search != "" {
		q.Where(
			q.Or(
				q.ILike("title", "%"+params.Search+"%"),
				q.ILike("author", "%"+params.Search+"%"),
			),
		)
	}

	q.Where(q.IsNull("deleted_at"))
	q.OrderBy("id")

	// store where clause for metadata query
	whereClauseNoPage := q.WhereClause

	// pagination compute
	page.Compute()
	q.Limit(page.Limit)
	q.Offset(page.Offset)

	// build and exec query
	query, args := q.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := repo.deps.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		return result, meta, err
	}
	defer rows.Close()

	var temp model.SQLBook
	for rows.Next() {
		err := rows.StructScan(&temp)
		if err != nil {
			repo.deps.Logger.WarnContext(ctx, "failed to scan book data", slog.Any("error", err))
			continue
		}

		result = append(result, model.Book{
			ID:          temp.ID.Int64,
			Title:       temp.Title.String,
			Author:      temp.Author.String,
			PublishYear: temp.PublishYear.Int64,
			BaseAudit: model.BaseAudit{
				CreatedAt: &temp.CreatedAt.Time,
				UpdatedAt: &temp.UpdatedAt.Time,
			},
		})
	}

	// build metadata
	var total int64
	countQ.WhereClause = whereClauseNoPage
	query, args = countQ.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err = repo.deps.DB.QueryRowxContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return result, meta, err
	}

	meta.Compute(total, page.Size, page.Page)

	if total == 0 {
		return result, meta, xerrors.ErrDataNotFound
	}

	return result, meta, nil
}

func (repo *BookRepo) GetBookByID(ctx context.Context, id int64) (model.Book, error) {
	var result model.SQLBook

	q := `SELECT id, title, author, publish_year, created_at, updated_at FROM library.books WHERE id = $1 AND deleted_at ISNULL;`

	err := repo.deps.DB.QueryRowxContext(ctx, q, id).StructScan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, xerrors.NewClientError(xerrors.ErrDataNotFound)
		}

		return model.Book{}, err
	}

	return model.Book{
		ID:          result.ID.Int64,
		Title:       result.Title.String,
		Author:      result.Author.String,
		PublishYear: result.PublishYear.Int64,
		BaseAudit: model.BaseAudit{
			CreatedAt: &result.CreatedAt.Time,
			UpdatedAt: &result.UpdatedAt.Time,
		},
	}, nil
}

func (repo *BookRepo) StoreBook(ctx context.Context, data model.Book) (model.Book, error) {
	var returned model.SQLBook
	q := `
		INSERT INTO library.books (title, author, publish_year) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`
	err := repo.deps.DB.QueryRowxContext(ctx, q, data.Title, data.Author, data.PublishYear).
		Scan(&returned.ID, &returned.CreatedAt, &returned.UpdatedAt)
	if err != nil {
		return model.Book{}, err
	}

	data.ID = returned.ID.Int64
	data.CreatedAt = &returned.CreatedAt.Time
	data.UpdatedAt = &returned.UpdatedAt.Time

	return data, nil
}

func (repo *BookRepo) UpdateBook(ctx context.Context, data model.Book) (model.Book, error) {
	var returned model.SQLBook

	q := `
		UPDATE library.books
			SET 
				title = $1,
				author = $2,
				publish_year = $3
			WHERE
				id = $4
			AND
				deleted_at ISNULL
		RETURNING updated_at;
	`
	tx, err := repo.deps.DB.BeginTxx(ctx, nil)
	if err != nil {
		return data, err
	}
	defer tx.Rollback()

	err = tx.QueryRowxContext(ctx, q, data.Title, data.Author, data.PublishYear, data.ID).Scan(&returned.UpdatedAt)
	if err != nil {
		// this means no data is updated
		// which probably caused by invalid id input (i.e. updating deleted entry)
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, xerrors.NewClientError(xerrors.ErrInvalidID)
		}

		return data, err
	}

	err = tx.Commit()
	if err != nil {
		repo.deps.Logger.ErrorContext(ctx, "failed to commit sql transaction", slog.Any("error", err))
		return data, err
	}

	data.UpdatedAt = &returned.UpdatedAt.Time

	return data, nil
}
func (repo *BookRepo) DeleteBook(ctx context.Context, id int64) error {
	q := `
		UPDATE library.books
			SET 
				deleted_at = now()
			WHERE
				id = $1
			AND
				deleted_at ISNULL;
	`
	tx, err := repo.deps.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		repo.deps.Logger.WarnContext(ctx, "failed to check affected row", slog.Any("error", err))
		return err
	}

	if rowsCount < 1 {
		// this means no data is soft deleted
		// which probably caused by invalid id input (i.e. deleting deleted entry)
		return xerrors.NewClientError(xerrors.ErrInvalidID)
	}

	err = tx.Commit()
	if err != nil {
		repo.deps.Logger.ErrorContext(ctx, "failed to commit sql transaction", slog.Any("error", err))
		return err
	}

	return nil
}

func (repo *BookRepo) GetBooksNoPagination(ctx context.Context, params model.BookSearchParams) ([]model.Book, error) {
	var result []model.Book

	// base query
	countQ := sqlbuilder.NewSelectBuilder()
	countQ.Select("COUNT(1)").From("library.books")

	q := sqlbuilder.NewSelectBuilder()
	q = q.Select("id", "title", "author", "publish_year", "created_at", "updated_at").From("library.books")

	if params.Search != "" {
		q.Where(
			q.Or(
				q.ILike("title", "%"+params.Search+"%"),
				q.ILike("author", "%"+params.Search+"%"),
			),
		)
	}

	q.Where(q.IsNull("deleted_at"))
	q.OrderBy("id")

	// build and exec query
	query, args := q.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := repo.deps.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var temp model.SQLBook
	for rows.Next() {
		err := rows.StructScan(&temp)
		if err != nil {
			repo.deps.Logger.WarnContext(ctx, "failed to scan book data", slog.Any("error", err))
			continue
		}

		result = append(result, model.Book{
			ID:          temp.ID.Int64,
			Title:       temp.Title.String,
			Author:      temp.Author.String,
			PublishYear: temp.PublishYear.Int64,
			BaseAudit: model.BaseAudit{
				CreatedAt: &temp.CreatedAt.Time,
				UpdatedAt: &temp.UpdatedAt.Time,
			},
		})
	}

	return result, nil
}
