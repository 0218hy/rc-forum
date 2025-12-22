package main

import (
	"context"
	"log/slog"
	"os"
	"rc-forum-backend/internal/env"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	
	cfg := config{
		addr: ":8080",
		db:   dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host = localhost user=rc_user password=rc_password dbname=rc_forum port=5432 sslmode=disable"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	logger.Info("Database connection established", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
