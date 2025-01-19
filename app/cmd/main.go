package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/massih/moody/internal"
	"github.com/massih/moody/internal/config"
	"log/slog"
)

func main() {
	ctx = context.Background()

	err := run()
	if err != nil {
		slog.ErrorContext(ctx, "fail to start", "error", err)
	}
}

func run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting Moody application")
	cfg := config.Get()

	pool, err := pgxpool.New(ctx, cfg.Home)
	if err != nil {
		return fmt.Errorf("failed to create database connection, error: %w", err)
	}
	defer pool.Close()

	moodHandler := internal.NewMoodHandler(pool)

	return nil
}
