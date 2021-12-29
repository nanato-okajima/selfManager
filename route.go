package main

import (
	"net/http"
	"selfManager/services/task"
)

func routeSetting(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", task.TaskListHandler)
}
