package main

import (
	"net/http"

	"github.com/selfManager/services/task"
)

func routeSetting(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", task.TaskListHandler)
}
