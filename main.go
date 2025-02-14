package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imran4u/web-socket/handler"
)

func main() {
	r := mux.NewRouter()
	// Define WebSocket route
	r.HandleFunc("/ws", handler.HandleWebSocket)

	// start server
	http.Handle("/", r)
	fmt.Println("Server listening on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Unable to start server")
	}
}
