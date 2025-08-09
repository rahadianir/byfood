package model

import (
	"database/sql"
	"time"
)

type BaseAudit struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type SQLBaseAudit struct {
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}


