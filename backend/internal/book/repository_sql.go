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
	countQ.Select("COUNT(1)").From("project.books")

	q := sqlbuilder.NewSelectBuilder()
	q = q.Select("id", "title", "author", "publish_year", "created_at", "updated_at").From("project.books")

	if params.Search != "" {
		q.Where(
			q.Or(
				q.ILike("title", "%"+params.Search+"%"),
				q.ILike("author", "%"+params.Search+"%"),
			),
		)
	}

	q.Where(q.IsNull("deleted_at"))

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
			Author:      temp.Title.String,
			PublishYear: temp.PublishYear.Int64,
			BaseAudit: model.BaseAudit{
				CreatedAt: temp.CreatedAt.Time,
				UpdatedAt: temp.UpdatedAt.Time,
			},
		})
	}

	// build metadata
	var total int64
	countQ.WhereClause = q.WhereClause
	query, args = countQ.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err = repo.deps.DB.QueryRowxContext(ctx, query, args...).Scan(&total)
	if err != nil {
		return result, meta, err
	}

	meta.Compute(total, page.Size, page.Page)

	return result, meta, nil
}

func (repo *BookRepo) GetBookByID(ctx context.Context, id int64) (model.Book, error) {
	var result model.SQLBook

	q := `SELECT id, title, author, publish_year, created_at, updated_at FROM library.books WHERE id = $1 AND deleted_at ISNULL;`

	err := repo.deps.DB.QueryRowxContext(ctx, q, id).Scan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, xerrors.ErrDataNotFound
		}

		return model.Book{}, err
	}

	return model.Book{
		ID:          result.ID.Int64,
		Title:       result.Title.String,
		Author:      result.Author.String,
		PublishYear: result.PublishYear.Int64,
		BaseAudit: model.BaseAudit{
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		},
	}, nil
}

func (repo *BookRepo) StoreBook(ctx context.Context, data model.Book) (model.Book, error) {
	q := `
		INSERT INTO library.books (title, author, publish_year) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at
	`
	err := repo.deps.DB.QueryRowxContext(ctx, q, data.Title, data.Author, data.PublishYear).
		Scan(&data.ID, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return model.Book{}, err
	}

	return data, nil
}

func (repo *BookRepo) UpdateBook(ctx context.Context, data model.Book) (model.Book, error) {
	q := `
		UPDATE library.books
			SET 
				title = $1,
				author = $2,
				publish_year = $3
			WHERE
				id = $4
			AND
				deleted_at ISNULL;
	`
	tx, err := repo.deps.DB.BeginTxx(ctx, nil)
	if err != nil {
		return data, err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(ctx, q, data.Title, data.Author, data.PublishYear, data.ID)
	if err != nil {
		return data, err
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		return data, err
	}

	if rowsCount < 1 {
		return data, xerrors.ErrDataNotFound
	}

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
		return err
	}

	if rowsCount < 1 {
		return xerrors.ErrDataNotFound
	}
	return nil
}
