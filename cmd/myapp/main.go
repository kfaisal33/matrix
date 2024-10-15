package main

import (
	"go-league-challenge/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/process", handlers.HandleProcess).Methods("POST")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Info().Msg("Starting server on :8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")

	}
}
