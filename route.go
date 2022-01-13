package main

import (
	"github.com/gorilla/mux"

	"selfManager/handler"
)

func routeSetting(mux *mux.Router, taskHandler handler.TaskHandler) {
	mux.HandleFunc("/tasks", taskHandler.Index).Methods("GET")
	mux.HandleFunc("/task/create", taskHandler.New).Methods("GET")
	mux.HandleFunc("/task/create", taskHandler.Create).Methods("POST")
	mux.HandleFunc("/task/{id:[0-9]+}", taskHandler.Show).Methods("GET")
	mux.HandleFunc("/task/{id:[0-9]+}", taskHandler.Update).Methods("POST")
	mux.HandleFunc("/task/delete/{id:[0-9]+}", taskHandler.Delete).Methods("POST")
}
