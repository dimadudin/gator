package main

import (
	"log"
	"net/http"
)

func showWelcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Welcome to Chadder!`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", showWelcome)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
