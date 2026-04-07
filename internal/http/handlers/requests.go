package handlers

import "strings"

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func (r *CreateTaskRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(r.Title) == "" {
		errors["title"] = "required"
	}

	return errors
}

func (r *UpdateTaskRequest) Validate() map[string]string {
	errors := make(map[string]string)

	if r.Title == nil && r.Done == nil {
		errors["body"] = "at least one field must be provided"
		return errors
	}

	if r.Title != nil && strings.TrimSpace(*r.Title) == "" {
		errors["title"] = "cannot be empty"
	}

	return errors
}