package repository

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) *sql.DB {
	dsn := "host=localhost port=5432 user=task_manager password=task_manager dbname=task_manager sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	// очистка таблицы
	_, err = db.Exec("TRUNCATE TABLE tasks RESTART IDENTITY")
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestTasksRepository_CreateAndGet(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTasksRepository(db)

	// Act
	task, err := repo.Create("test task")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}

	// Assert
	if task.ID == 0 {
		t.Fatalf("expected ID to be set")
	}

	got, err := repo.Get(task.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}

	if got.Title != "test task" {
		t.Fatalf("expected title 'test task', got %s", got.Title)
	}
}