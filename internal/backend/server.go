package backend

import (
	"encoding/json"
	"go-2x2-solver/pkg/solver"
	"io"
	"log"
	"net/http"
)

type Server struct{}
type requestData struct {
	Cube [24]solver.Sticker `json:"cube"`
}

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
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Broken request body [%s]\n", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("request body: [%s]\n", requestBody)
	var unmarshalled requestData
	if err := json.Unmarshal(requestBody, &unmarshalled); err != nil {
		log.Printf("Can not unmarshal reqest body [%s]\n", err.Error())
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// set up cube
	cube := solver.NewEmptyCube()
	cube.SetStickers(unmarshalled.Cube)
	if !cube.IsValid() {
		//log.Printf("Non-valid cube [%v]\n", cube)
		_, err := writer.Write([]byte("Your cube is non-valid, try again"))
		if err != nil {
			log.Printf("Error while writing response [%s]", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// solve
	solution := "This is your solution:\nR U R' F\n"

	if _, err := writer.Write([]byte(solution)); err != nil {
		log.Printf("Error while writing response [%s]", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) Read(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./internal/frontend/index.html")
}
