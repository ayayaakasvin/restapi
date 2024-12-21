package main

import (
	"fmt"
	"log/slog"

	"restapi/internal/config"
)

const (
	envProd = "prod"
	envDev 	= "dev"
	envLocal = "local"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Println(cfg)
}

func setupLogger(env string)  {
	
}