package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewPool(ctx context.Context, dsn string) (*Database, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to parse dsn: %s", err)
	}
	config.MaxConns = 5
	config.MinConns = 1
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 10 * time.Minute
	config.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to create connection pool:%w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres: unable to ping database: %w", err)
	}
	return &Database{Pool: pool}, nil
}

func (db *Database) Close() {
	if db == nil || db.Pool == nil {
		return
	}
	db.Pool.Close()
}
func (db *Database) Ping(ctx context.Context) error {
	if db == nil || db.Pool == nil {
		return fmt.Errorf("postgres: pool is nil")
	}
	return db.Pool.Ping(ctx)
}
