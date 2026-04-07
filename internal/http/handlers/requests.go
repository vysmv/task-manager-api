package handlers

import "strings"

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func (r *CreateTaskRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(r.Title) == "" {
		errors["title"] = "required"
	}

	return errors
}