package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"backend/go/books/internal/config"
	"backend/go/books/internal/rest"
	"backend/go/books/internal/service"
	"backend/go/books/internal/storage"
)

// Сборка и запуск приложения

// Run собирает и запускает приложение.
func Run(cfg *config.Config) error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	listener, errListen := net.Listen("tcp", cfg.Port)
	fmt.Printf("asdads")
	if errListen != nil {
		return fmt.Errorf("listen error: %w", errListen)
	}
	defer listener.Close()

	storageLayer := storage.New()
	serviceLayer := service.New(storageLayer)

	server := rest.New(cfg.Server, serviceLayer)

	go func() {
		<-ctx.Done()

		slog.InfoContext(ctx, "server shutdown")
		if errShutdown := server.Shutdown(ctx); errShutdown != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("error shutdown server, %s", errShutdown.Error()))
		}
	}()

	slog.InfoContext(ctx, fmt.Sprintf("server serve on %s", cfg.Server.Address))
	if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error serve http, %w", err)
	}

	return nil
}
