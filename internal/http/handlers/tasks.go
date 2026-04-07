package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/vysmv/task-manager-api/internal/http/response"
	"github.com/vysmv/task-manager-api/internal/tasks"
)

type TasksHandler struct {
	storage *tasks.MemoryStorage
}

func NewTasksHandler(storage *tasks.MemoryStorage) *TasksHandler {
	return &TasksHandler{storage: storage}
}

func (h *TasksHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if errs := req.Validate(); len(errs) > 0 {
		response.WriteValidationError(w, errs)
		return
	}

	task := h.storage.Create(req.Title)

	response.WriteJSON(w, http.StatusCreated, task)
}

func (h *TasksHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks := h.storage.List()

	response.WriteJSON(w, http.StatusOK, tasks)
}

func (h *TasksHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
		return
	}

	task, ok := h.storage.Get(id)
	if !ok {
		response.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "task not found",
		})
		return
	}

	response.WriteJSON(w, http.StatusOK, task)
}

func (h *TasksHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
		return
	}

	ok := h.storage.Delete(id)
	if !ok {
		response.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "task not found",
		})
		return
	}

	response.WriteJSON(w, http.StatusNoContent, nil)
}