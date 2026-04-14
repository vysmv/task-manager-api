package routes

import (
	"net/http"

	"github.com/vysmv/task-manager-api/internal/http/handlers"
	"github.com/vysmv/task-manager-api/internal/http/middleware"
)

func Register(tasksHandler *handlers.TasksHandler, mux *http.ServeMux) {
	mux.Handle(
		"POST /tasks",
		middleware.BasicAuth(http.HandlerFunc(tasksHandler.Create)),
	)

	mux.Handle(
		"GET /tasks",
		middleware.BasicAuth(middleware.Logging(http.HandlerFunc(tasksHandler.List))),
	)

	mux.Handle(
		"GET /tasks/",
		middleware.BasicAuth(http.HandlerFunc(tasksHandler.Get)),
	)

	mux.Handle(
		"PATCH /tasks/",
		middleware.BasicAuth(http.HandlerFunc(tasksHandler.Update)),
	)

	mux.Handle(
		"DELETE /tasks/",
		middleware.APIKeyAuth(http.HandlerFunc(tasksHandler.Delete)),
	)

}
