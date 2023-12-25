package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTasks(t *testing.T) {
	tasks = []Task{
		{ID: "task1", Description: "Task 1", Priority: 1},
		{ID: "task2", Description: "Task 2", Priority: 2},
	}

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getTasks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `[{"id":"task1","description":"Task 1","priority":1},{"id":"task2","description":"Task 2","priority":2}]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreateTask(t *testing.T) {
	tasks = nil

	newTask := Task{Description: "New Task", Priority: 3}

	jsonData, err := json.Marshal(newTask)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(createTask)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if len(tasks) != 1 || tasks[0].Description != "New Task" || tasks[0].Priority != 3 {
		t.Errorf("Handler did not add the task correctly to the tasks list")
	}
}

func TestDeleteTask(t *testing.T) {
	tasks = []Task{
		{ID: "task1", Description: "Task 1", Priority: 1},
		{ID: "task2", Description: "Task 2", Priority: 2},
	}
	req, err := http.NewRequest("DELETE", "/tasks/task1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if len(tasks) != 1 || tasks[0].ID != "task2" {
		t.Errorf("Handler did not delete the task correctly from the tasks list")
	}
}

func TestUpdateTask(t *testing.T) {
	tasks = []Task{
		{ID: "task1", Description: "Task 1", Priority: 1},
		{ID: "task2", Description: "Task 2", Priority: 2},
	}

	updatedTask := Task{Description: "Updated Task", Priority: 3}

	jsonData, err := json.Marshal(updatedTask)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/tasks/task1", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", updateTasks).Methods("PUT")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if len(tasks) != 2 || tasks[0].Description != "Updated Task" {
		t.Errorf("Handler did not update the task correctly in the tasks list")
	}
}
func TestDeleteTaskNonExistent(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/tasks/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", status)
	}
}

func TestUpdateTaskInvalidRequestBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/tasks/task1", strings.NewReader("invalid json"))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/{id}", updateTasks).Methods("PUT")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected 400 Bad Request, got %v", status)
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
