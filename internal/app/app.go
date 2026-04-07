package app

import (
	"fmt"
	"net/http"

	"github.com/vysmv/task-manager-api/internal/config"
	"github.com/vysmv/task-manager-api/internal/http/handlers"
	"github.com/vysmv/task-manager-api/internal/storage/postgres"
	"github.com/vysmv/task-manager-api/internal/tasks/repository"
)

func Run() error {
	cfg := config.MustLoad()

	db, err := postgres.New(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	tasksRepo := repository.NewTasksRepository(db)
	tasksHandler := handlers.NewTasksHandler(tasksRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

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