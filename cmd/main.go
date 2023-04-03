package main

import (
	"go-2x2-solver/internal/backend"
	"log"
	"net/http"
)

const (
	root = "/"
	port = ":8080"
)

func main() {
	http.HandleFunc(root, backend.Handler(backend.Server{}))

	fs := http.FileServer(http.Dir("internal/frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
