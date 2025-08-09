package model

type Book struct {
	ID          int64
	Title       string
	Author      string
	PublishYear int64
	BaseAudit
}

type SQLBook struct {
	ID          int64
	Title       string
	Author      string
	PublishYear int64
}

type BookSearchParams struct {
	Search string
}