package main

import (
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

	db, err := storage.NewPostgresStorage(cfg)
	if err != nil {
		log.Error("failed to setup storage", sl.Err(err))
		os.Exit(1)
	}

	defer db.Close()
	defer db.Reset()

	if err = db.SaveUser("test", "test"); err != nil {
		log.Error("failed to save user", sl.Err(err))
	}

	if exists, err := db.UsernameExists("test"); err != nil {
		log.Error("failed to check if username exists", sl.Err(err))
	} else if exists {
		log.Info("username exists")
	} else {
		log.Info("username does not exist")
	}

	if user, err := db.GetUserByID(1); err != nil {
		log.Error("failed to get user by ID", sl.Err(err))
		os.Exit(1)
	} else {
		log.Info("user found", sl.Any("user", user))
	}

	if err = db.DeleteUser(1); err != nil {
		log.Error("failed to delete user by ID", sl.Err(err))
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