package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/prajwalbharadwajbm/todo-app-backend/pkg/task"
)

func main() {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		fmt.Println("Error opening SQLite database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id TEXT PRIMARY KEY,
			description TEXT,
			priority TEXT
		)
	`)
	if err != nil {
		fmt.Println("Error creating tasks table:", err)
		return
	}
	taskRepository := task.NewSQLiteRepository(db)
	taskHandler := &task.Handler{TaskRepository: taskRepository}
	router := task.NewRouter(taskHandler)

	http.ListenAndServe(":8080", router)
}
