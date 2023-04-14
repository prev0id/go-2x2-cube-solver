package main

import (
	"github.com/joho/godotenv"
	"go-2x2-solver/internal/backend"
	"log"
	"net/http"
	"os"
)

func main() {
	// loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// handle server
	http.HandleFunc("/", backend.Handler(backend.Server{}))
	//load static files
	fs := http.FileServer(http.Dir("internal/frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err.Error())
	}
}
