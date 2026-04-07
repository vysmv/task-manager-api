package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/vysmv/task-manager-api/internal/http/response"
	"github.com/vysmv/task-manager-api/internal/tasks/repository"
)

type TasksHandler struct {
	repo *repository.TasksRepository
}

func NewTasksHandler(repo *repository.TasksRepository) *TasksHandler {
	return &TasksHandler{repo: repo}
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

	task, err := h.repo.Create(req.Title)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to create task")
		return
	}

	response.WriteJSON(w, http.StatusCreated, task)
}

func (h *TasksHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.List()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to fetch tasks")
		return
	}

	response.WriteJSON(w, http.StatusOK, tasks)
}

func (h *TasksHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	task, err := h.repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "failed to fetch task")
		return
	}

	response.WriteJSON(w, http.StatusOK, task)
}

func (h *TasksHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if errs := req.Validate(); len(errs) > 0 {
		response.WriteValidationError(w, errs)
		return
	}

	task, err := h.repo.Update(id, req.Title, req.Done)
	if err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "failed to update task")
		return
	}

	response.WriteJSON(w, http.StatusOK, task)
}

func (h *TasksHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		if errors.Is(err, repository.ErrTaskNotFound) {
			response.WriteError(w, http.StatusNotFound, "task not found")
			return
		}

		response.WriteError(w, http.StatusInternalServerError, "failed to delete task")
		return
	}

	response.WriteJSON(w, http.StatusNoContent, nil)
}