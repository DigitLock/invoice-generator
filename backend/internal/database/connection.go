package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/DigitLock/invoice-generator/backend/internal/config"
	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
)

type DB struct {
	Pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

func New(ctx context.Context, cfg config.DatabaseConfig) (*DB, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("parse database config: %w", err)
	}

	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &DB{
		Pool:    pool,
		Queries: sqlc.New(pool),
	}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}
