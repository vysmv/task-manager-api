package handlers

import (
	"net/http"

	"github.com/vysmv/task-manager-api/internal/http/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}