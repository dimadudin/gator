package main

import (
	"log"
	"net/http"
	"os"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var wsUpgrader = ws.Upgrader{
	HandshakeTimeout: 20 * time.Second,
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
}

func serveWelcomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Welcome to Chadder!`))
}

func serveWebSocketPage(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	go writePump(conn)
	go readPump(conn)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn't open .env file")
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", serveWelcomePage)
	mux.HandleFunc("GET /ws", serveWebSocketPage)

	server := http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
