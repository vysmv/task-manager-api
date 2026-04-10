package http_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vysmv/task-manager-api/internal/http/handlers"
	"github.com/vysmv/task-manager-api/internal/tasks/repository"

	"database/sql"
	_ "github.com/lib/pq"
)

func setupServer(t *testing.T) http.Handler {
	dsn := "host=localhost port=5432 user=task_manager password=task_manager dbname=task_manager sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("TRUNCATE TABLE tasks RESTART IDENTITY")
	if err != nil {
		t.Fatal(err)
	}

	repo := repository.NewTasksRepository(db)
	h := handlers.NewTasksHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("GET /tasks", h.List)

	return mux
}

func TestAPI_CreateAndList(t *testing.T) {
	server := setupServer(t)

	// === Create ===
	body := []byte(`{"title":"e2e task"}`)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// === List ===
	req = httptest.NewRequest("GET", "/tasks", nil)
	w = httptest.NewRecorder()

	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if !bytes.Contains(w.Body.Bytes(), []byte("e2e task")) {
		t.Fatalf("task not found in response")
	}
}