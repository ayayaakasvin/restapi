package main

import (
	"fmt"
	"restapi/internal/config"
	"log/slog"
)

const (
	envProd = "prod"
	envDev 	= "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Println(cfg)
}

func setupLogger(env string)  {
	
}