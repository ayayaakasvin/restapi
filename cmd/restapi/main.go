package main

import (
	// "fmt"
	"log/slog"
	"os"

	"restapi/internal/config"
	"restapi/internal/lib/sl"
	"restapi/internal/storage"
)

const (
	envProd = "prod"
	envDev 	= "dev"
	envLocal = "local"
)

func main() {
	cfg := config.MustLoadConfig()

	log := setupLogger(cfg.Env)

	log.Info("starting RESTful API")
	log.Debug("debug message", slog.String("env", cfg.Env))

	db, err := storage.NewPostgresStorageWithConfig(cfg.Database)
	if err != nil {
		log.Error("failed to setup storage", sl.Err(err))
		os.Exit(1)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Error("failed to ping database", sl.Err(err))
		os.Exit(1)
	}

	// TODO: implement the RESTful API with Gin

	// TODO: run the HTTP server
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, 
			&slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envDev, envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return logger
}