package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/selfManager/domain/repository"
)

func main() {
	mux := http.NewServeMux()
	routeSetting(mux)
	server := http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	repository.Migrate()
	fmt.Println("Starting web server...")
	log.Fatal(server.ListenAndServe())
}
