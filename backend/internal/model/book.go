package model

import "database/sql"

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishYear int64  `json:"publish_year"`
	BaseAudit
}

type SQLBook struct {
	ID          sql.NullInt64  `db:"id"`
	Title       sql.NullString `db:"title"`
	Author      sql.NullString `db:"author"`
	PublishYear sql.NullInt64  `db:"publish_year"`
	SQLBaseAudit
}

type BookSearchParams struct {
	Search string
}
