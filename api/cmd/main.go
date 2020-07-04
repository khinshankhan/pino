package main

import (
	"github.com/woojiahao/pino/api/pkg/server"
	"net/http"
)

const (
	port = 8080
)

func ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Pong!"))
}

func main() {
	s := server.New(3000)
	s.AddEndpoint(server.GET, "/", ping)
	s.Start()
}
