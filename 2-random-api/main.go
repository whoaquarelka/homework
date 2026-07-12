package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/random", randomNumber)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("server is running on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func randomNumber(w http.ResponseWriter, r *http.Request) {
	rndm := rand.Intn(6) + 1
	result := strconv.Itoa(rndm)
	log.Printf("generated number: %s", result)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, result)
}
