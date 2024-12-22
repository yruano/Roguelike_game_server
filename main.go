package main

import (
	"fmt"
	"log"
	"net/http"

	"server/server"
)

func main() {
	http.HandleFunc("/ws", server.HandleConnections)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
