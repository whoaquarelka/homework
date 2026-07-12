package main

import (
	"log"
	"net/http"
	"time"
	"verify/3-validation-api/configs"
	"verify/3-validation-api/internal/email"
)

func main() {
	mux := http.NewServeMux()
	cfg := configs.LoadConfig()

	email.NewEmailHandlers(mux, cfg)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("server start and listening on port %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
