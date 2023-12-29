package task

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	CreateTask(task Task) error
	UpdateTask(task Task) error
	DeleteTask(id string) error
}

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) GetAllTasks() ([]Task, error) {
	rows, err := r.db.Query("SELECT id, description, priority FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Description, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *SQLiteRepository) GetTaskByID(id string) (Task, error) {
	var task Task
	err := r.db.QueryRow("SELECT id, description, priority FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Description, &task.Priority)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *SQLiteRepository) CreateTask(task Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (id, description, priority) VALUES (?, ?, ?)",
		task.ID, task.Description, task.Priority)
	return err
}

func (r *SQLiteRepository) UpdateTask(task Task) error {
	_, err := r.db.Exec("UPDATE tasks SET description = ?, priority = ? WHERE id = ?",
		task.Description, task.Priority, task.ID)
	return err
}

func (r *SQLiteRepository) DeleteTask(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
