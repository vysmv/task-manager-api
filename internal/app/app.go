package app

import (
	"fmt"
	"net/http"

	"github.com/vysmv/task-manager-api/internal/config"
	"github.com/vysmv/task-manager-api/internal/http/handlers"
	"github.com/vysmv/task-manager-api/internal/tasks"
)

func Run() error {
	cfg := config.MustLoad()

	storage := tasks.NewMemoryStorage()
	tasksHandler := handlers.NewTasksHandler(storage)

	mux := http.NewServeMux()

	// health
	mux.HandleFunc("GET /health", handlers.Health)

	// tasks
	mux.HandleFunc("POST /tasks", tasksHandler.Create)
	mux.HandleFunc("GET /tasks", tasksHandler.List)
	mux.HandleFunc("GET /tasks/", tasksHandler.Get)
	mux.HandleFunc("DELETE /tasks/", tasksHandler.Delete)

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	fmt.Printf("server started on :%s\n", cfg.HTTPPort)

	return server.ListenAndServe()
}