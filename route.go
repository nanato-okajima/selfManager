package main

import (
	"github.com/gorilla/mux"

	"selfManager/services/task"
)

func routeSetting(mux *mux.Router) {
	mux.HandleFunc("/tasks", task.TaskListHandler).Methods("GET")
	mux.HandleFunc("/task/create", task.NewTaskHandler).Methods("GET")
	mux.HandleFunc("/task/create", task.CreateTaskHandler).Methods("POST")
	mux.HandleFunc("/task/{id:[0-9]+}", task.TaskHandler).Methods("GET")
	mux.HandleFunc("/task/{id:[0-9]+}", task.UpdateTaskHandler).Methods("POST")
	mux.HandleFunc("/task/delete/{id:[0-9]+}", task.TaskDeleteHandler).Methods("POST")
}
