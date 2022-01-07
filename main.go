package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"selfManager/domain/repository"
)

func main() {
	r := mux.NewRouter()
	routeSetting(r)
	server := http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	if err := repository.SetupDB(".env.dev"); err != nil {
		log.Fatal(err)
	}
	if err := repository.Migrate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting web server...")
	log.Fatal(server.ListenAndServe())
}
