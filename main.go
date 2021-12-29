package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func tasksHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	mux := http.NewServeMux()
	routeSetting(mux)
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	fmt.Println("Starting web server...")
	log.Fatal(server.ListenAndServe())
}
