package task

import "github.com/gorilla/mux"

func NewRouter(taskHandler *Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")

	return router
}
