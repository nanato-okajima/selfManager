package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"selfManager/config"
	"selfManager/handler"
	"selfManager/infrastructure/persistence"
	"selfManager/usecase/task"
)

func main() {
	config.SetEnv(".env.dev")
	taskPersistence := persistence.NewTaskPersistence(config.Connect())
	taskUseCase := task.NewTaskUseCase(taskPersistence)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	r := mux.NewRouter()
	routeSetting(r, taskHandler)
	server := http.Server{
		Addr:         ":80",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	fmt.Println("Starting web server...")
	log.Fatal(server.ListenAndServe())
}
