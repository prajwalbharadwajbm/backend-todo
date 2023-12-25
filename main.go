package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

var tasks []Task

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	newTask.ID = "task" + fmt.Sprint(len(tasks))
	tasks = append(tasks, newTask)
	json.NewEncoder(w).Encode(newTask)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid task ID"})
		return
	}

	var taskIndex = -1
	if tasks == nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Tasks is nil."})
	} else {
		for i, task := range tasks {
			if task.ID == taskID {
				taskIndex = i
				break
			}
		}
	}
	if taskIndex != -1 {
		w.WriteHeader(http.StatusOK)
		tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
		json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted Successfully"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
	}
}

func updateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	taskID := vars["id"]

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	var taskIndex = -1
	for i, task := range tasks {
		if task.ID == taskID {
			taskIndex = i
			break
		}
	}
	if taskIndex != -1 {
		updatedTask.ID = taskID
		tasks[taskIndex] = updatedTask
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Task Updated Successfully"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update task"})
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTasks).Methods("PUT")

	http.ListenAndServe(":8080", router)
}
