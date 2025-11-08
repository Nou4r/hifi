package main

import (
	"api/config"
	"api/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	port := middleware.PortRotate()
	handler := middleware.Recovery(mux)

	// Server setup

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, port),
		Handler: handler,
	}

	slog.Info("Hifi Web server running",
		"host", config.Host,
		"port", port,
		"url", fmt.Sprintf("%s://%s:%s", config.HifiScheme, config.Host, port),
	)

	// Run server in background
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()
}
