package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

type Server struct {
	port      uint16
	endpoints []Endpoint
}

func (s *Server) AddEndpoint(method Method, path string, action func(w http.ResponseWriter, r *http.Request)) {
	s.endpoints = append(s.endpoints, Endpoint{method, path, action})
}

func (s *Server) Start() {
	log.Printf("Starting server")
	r := chi.NewRouter()

	// Enable logging
	r.Use(middleware.Logger)

	// Load endpoints
	for _, endpoint := range s.endpoints {
		log.Printf("Loading %s endpoint at %s", endpoint.method, endpoint.path)
		switch endpoint.method {
		case GET:
			r.Get(endpoint.path, endpoint.action)
			break
		case POST:
			r.Post(endpoint.path, endpoint.action)
			break
		case DELETE:
			r.Delete(endpoint.path, endpoint.action)
			break
		case PUT:
			r.Put(endpoint.path, endpoint.action)
			break
		}
	}

	log.Printf("Connected to port %d. Access at http://localhost:%d", s.port, s.port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
}

func New(port uint16) *Server {
	return &Server{port, make([]Endpoint, 0)}
}
