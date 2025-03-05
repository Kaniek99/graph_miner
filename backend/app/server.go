package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	port   int
	server *http.Server
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := map[string]string{"message": "Hello World from the server on the container!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func NewServer(port int) *HttpServer {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return &HttpServer{
		port:   port,
		server: srv,
	}
}

func (s *HttpServer) Start() error {
	log.Printf("Server is running on port %d\n", s.port)
	return s.server.ListenAndServe()
}

func (s *HttpServer) Stop() error {
	log.Printf("Server is shutting down\n")
	return s.server.Close()
}
