package main

import (
	"log"
	"net/http"

	"github.com/TaufeeqAi/backend/internal/api"
)

func main() {
	http.HandleFunc("/hello", api.HelloHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
