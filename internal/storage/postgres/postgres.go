package postgres

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func NewConnectDB(cfg config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBName,
	)

	configPool, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config error: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), configPool)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool with config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}
