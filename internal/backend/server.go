package backend

import (
	"log"
	"net/http"
)

type Server struct{}

func Handler(server Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			server.Create(writer, request)
		case http.MethodGet:
			server.Read(writer, request)
		default:
			log.Println("Unsupported request")
		}
	}
}

func (s *Server) Create(writer http.ResponseWriter, request *http.Request) {
	log.Println("POST request")
}

func (s *Server) Read(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get shit")
	//http.StripPrefix()
	http.ServeFile(writer, request, "./internal/frontend/index.html")
}
