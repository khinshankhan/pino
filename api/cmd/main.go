package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kkhan01/pino/api/pkg/database"
	"github.com/kkhan01/pino/api/pkg/server"
)

const (
	port = 3000
)

var (
	dbUser = getUser()
	dbPass = getPass()
	dbAddr = getHost()
	dbPort = getPort()
)

func getUser() string {
	u := os.Getenv("DB_USER")
	if u == "" {
		log.Fatalf("No user provided")
	}
	return u
}

func getPass() string {
	p := os.Getenv("DB_PASS")
	if p == "" {
		log.Fatalf("No pass provided")
	}
	return p
}

func getPort() uint16 {
	p := os.Getenv("DB_PORT")
	if p == "" {
		return 27017
	}
	port, err := strconv.Atoi(p)

	if err != nil {
		log.Fatalf("Invalid port")
	}

	return uint16(port)
}

func getHost() string {
	h := os.Getenv("DB_HOST")
	if h == "" {
		return "localhost"
	}

	return h
}

func ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Pong!"))
}

func main() {
	// Connect to MongoDB database
	db := database.New(dbUser, dbPass, dbAddr, dbPort)
	err := db.Connect()

	if err != nil {
		log.Fatalf("error occurred, %v", err)
	}

	// Listen to program stop signals to properly close the database resources
	ch := make(chan os.Signal, 3)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		_ = <-ch
		signal.Stop(ch)
		log.Println("Shutting down server")
		db.Close()
		os.Exit(0)
	}()

	s := server.New(port)
	s.AddEndpoint(server.GET, "/", ping)
	s.Start()
}
