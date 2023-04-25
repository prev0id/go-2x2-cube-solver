package backend

import (
	"encoding/json"
	"go-2x2-solver/pkg/cube"
	"go-2x2-solver/pkg/solver"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Server struct{}
type requestData struct {
	Cube cube.Cube `json:"cube"`
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

	// solve
	startClock := time.Now()
	algorithm, success := solver.Solve(unmarshalled.Cube)
	endClock := time.Now()
	log.Printf("Solved in %d ms", endClock.Sub(startClock).Milliseconds())

	if _, err := writer.Write(makeResponseBody(algorithm, success)); err != nil {
		log.Printf("Error while writing response [%s]", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *Server) Read(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "./internal/frontend/index.html")
}

func makeResponseBody(algorithm solver.Algorithm, success error) []byte {
	if success != nil {
		return []byte("The cube is unsolvable, check the stickers")
	}
	moveToString := []string{
		cube.R:      "R",
		cube.R2:     "R2",
		cube.RPrime: "R'",
		cube.F:      "F",
		cube.F2:     "F2",
		cube.FPrime: "F'",
		cube.U:      "U",
		cube.U2:     "U2",
		cube.UPrime: "U'",
	}

	var resultBuilder strings.Builder
	for _, move := range algorithm {
		resultBuilder.WriteString(moveToString[move])
		resultBuilder.WriteRune(' ')
	}
	return []byte(resultBuilder.String())
}
