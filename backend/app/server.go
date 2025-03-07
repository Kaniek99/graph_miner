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

func handleCORSPreflight(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

func handleRootPathRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("Path %s not found\n", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodOptions {
		handleCORSPreflight(w)
	}

	log.Printf("Received %s request with body: %s\n", r.Method, r.Body)

	response := map[string]string{"message": "Hello World from the server on the container!"}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v\n", err)
	}
}

func NewServer(port int) *HttpServer {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRootPathRequest)
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
