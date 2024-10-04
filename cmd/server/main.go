package main

import (
	"fmt"
	"github.com/perkzen/tttm-go/pkg/handlers"
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HandleRoot)
	mux.HandleFunc("GET /move", handlers.HandleGetMove)

	log.Println(fmt.Sprintf("Server is running on : http://localhost:%s", PORT))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}
