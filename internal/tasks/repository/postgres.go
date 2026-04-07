package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/vysmv/task-manager-api/internal/tasks"
)

var ErrTaskNotFound = errors.New("task not found")

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) *TasksRepository {
	return &TasksRepository{db: db}
}

func (r *TasksRepository) Create(title string) (tasks.Task, error) {
	var task tasks.Task

	err := r.db.QueryRow(
		`
		INSERT INTO tasks (title, done)
		VALUES ($1, false)
		RETURNING id, title, done
		`,
		title,
	).Scan(&task.ID, &task.Title, &task.Done)
	if err != nil {
		return tasks.Task{}, err
	}

	return task, nil
}

func (r *TasksRepository) List() ([]tasks.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, done
		FROM tasks
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]tasks.Task, 0)

	for rows.Next() {
		var task tasks.Task

		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			return nil, err
		}

		result = append(result, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TasksRepository) Get(id int64) (tasks.Task, error) {
	var task tasks.Task

	err := r.db.QueryRow(
		`
		SELECT id, title, done
		FROM tasks
		WHERE id = $1
		`,
		id,
	).Scan(&task.ID, &task.Title, &task.Done)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tasks.Task{}, ErrTaskNotFound
		}

		return tasks.Task{}, err
	}

	return task, nil
}

func (r *TasksRepository) Update(id int64, title *string, done *bool) (tasks.Task, error) {
	// формируем динамический UPDATE
	query := "UPDATE tasks SET "
	args := make([]any, 0)
	argID := 1

	if title != nil {
		query += "title = $" + strconv.Itoa(argID)
		args = append(args, *title)
		argID++
	}

	if done != nil {
		if len(args) > 0 {
			query += ", "
		}
		query += "done = $" + strconv.Itoa(argID)
		args = append(args, *done)
		argID++
	}

	query += " WHERE id = $" + strconv.Itoa(argID)
	args = append(args, id)

	query += " RETURNING id, title, done"

	var task tasks.Task

	err := r.db.QueryRow(query, args...).Scan(
		&task.ID,
		&task.Title,
		&task.Done,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tasks.Task{}, ErrTaskNotFound
		}
		return tasks.Task{}, err
	}

	return task, nil
}

func (r *TasksRepository) Delete(id int64) error {
	result, err := r.db.Exec(
		`
		DELETE FROM tasks
		WHERE id = $1
		`,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}
