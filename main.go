package main

import (
	"fmt"
	"log"
	"net/http"
)

func startServer(address string) {
	log.Printf("Starting HTTP server at address %s", address)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	http.ListenAndServe(address, nil)
}

func main() {
	startServer("localhost:8080")
}
