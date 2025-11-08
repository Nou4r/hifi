package main

import (
	"api/config"
	"api/middleware"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

type routes struct {
	value bool
}

func (ch routes) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, _ = rw.Write([]byte(strconv.FormatBool(ch.value)))
}

func Handle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/signup", routes{value: middleware.SignupUser()})
	mux.Handle("/signin", routes{value: false})
	return mux
}

func main() {

	mux := http.NewServeMux()

	// API v1 routes
	mux.Handle("/v1/", http.StripPrefix("/v1", Handle()))

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

	select {}
}
