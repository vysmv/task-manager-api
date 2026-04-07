package app

import (
	"fmt"
	"net/http"

	"github.com/vysmv/task-manager-api/internal/config"
	"github.com/vysmv/task-manager-api/internal/http/handlers"
)

func Run() error {
	cfg := config.MustLoad()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.Health)

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	fmt.Printf("server started on :%s\n", cfg.HTTPPort)

	return server.ListenAndServe()
}