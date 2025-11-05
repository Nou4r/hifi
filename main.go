package main

import (
	"context"
	"fmt"
	"hifi/config"
	"hifi/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		middleware.StartTidalRefresher()
	}()
	go func() {
		defer wg.Done()
		middleware.RecentAlbum()
	}()

	wg.Wait()

	// Define subsonic user credentials
	person := config.Person{
		UserName: "",
		PassWord: "",
	}

	// Hifi proxy
	validPaths := config.ValidPaths
	targetHost := config.TargetHost

	// HTTP server setup
	mux := http.NewServeMux()

	// Middleware setup
	session := middleware.Session(person.UserName, person.PassWord, targetHost, validPaths)(mux)

	cors := middleware.CORS(session)

	handler := middleware.Recovery(cors)

	// Server setup
	port := middleware.PortRotate()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, port),
		Handler: handler,
	}

	slog.Info("Hifi API server running",
		"host", config.Host,
		"port", port,
		"url", fmt.Sprintf("%s://%s:%s", config.Scheme, config.Host, port),
	)

	// Run server in background
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Shutdown complete")

}
