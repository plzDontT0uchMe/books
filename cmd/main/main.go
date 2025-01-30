package main

import (
	"fmt"
	"os"

	"backend/go/books/internal/app"
	"backend/go/books/internal/config"
	"backend/go/books/pkg/logger"
)

// Точка входа в приложение

// main точка входа в приложение.
func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		fmt.Printf("config load error: %s", err)
		os.Exit(1)
	}

	if err = logger.MustLoad(cfg.Environment); err != nil {
		fmt.Printf("logger load error: %s", err)
		os.Exit(1)
	}

	if err = app.Run(cfg); err != nil {
		fmt.Printf("app run error: %s", err)
		os.Exit(1)
	}
}
