package main

import (
	"go-2x2-solver/internal/backend"
	"log"
	"net/http"
)

const (
	root = "/"
	port = ":8001"
)

func main() {
	http.HandleFunc(root, backend.Handler(backend.Server{}))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err.Error())
	}
}
