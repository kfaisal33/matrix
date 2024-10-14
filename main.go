package main

import (
	"go-league-challenge/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/process", handlers.HandleProcess).Methods("POST") // Specify method if applicable

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
