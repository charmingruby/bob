package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{ .RootPath }}/config"
	"{{ .RootPath }}/internal/{{ .Module }}"
	"{{ .RootPath }}/internal/shared/transport/rest"
	"{{ .RootPath }}/internal/shared/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	config, err := config.New()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %v", err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	restServer := rest.NewServer(config.ServerConfig.Port, router)

	initModules(router)

	shutdown := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		slog.Info(fmt.Sprintf("SHUTDOWN: signal caught %s", s))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		slog.Info("SHUTDOWN: Initiating graceful shutdown")
		shutdown <- restServer.Shutdown(ctx)
	}()

	slog.Info(fmt.Sprintf("REST SERVER: Running on port %s", config.ServerConfig.Port))
	if err := restServer.Run(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %v", err))
			os.Exit(1)
		}
	}

	err = <-shutdown
	if err != nil {
		slog.Error(fmt.Sprintf("REST SERVER: %v", err))
		os.Exit(1)
	}

	slog.Info("REST SERVER: has gracefully shutdown")
}

func initModules(r *chi.Mux) {
	endpoint.New(r).Register()

	{{ .Module }}Svc := {{ .Module }}.NewService()
	{{ .Module }}.NewHTTPHandler(r, {{ .Module }}Svc)
}
