package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

const (
	port = 8080
)

func main() {
	log.Print("Starting server")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("welcome"))
	})
	log.Printf("Connected to port %d. Access at http://localhost:%d", port, port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
