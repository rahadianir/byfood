package core

import (
	"byfood-app/internal/config"
	"context"
	"log/slog"

	"byfood-app/internal/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Dependency struct {
	*sqlx.DB
	*config.Config
	*slog.Logger
}

func NewDependency(ctx context.Context, cfg *config.Config) *Dependency {
	// setup logger
	logger := logger.NewLogger()

	// setup dependencies
	// setup database
	db, err := sqlx.ConnectContext(ctx, "postgres", cfg.DBURL)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	return &Dependency{
		DB:     db,
		Config: cfg,
		Logger: logger,
	}
}
