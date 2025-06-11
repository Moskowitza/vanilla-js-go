package main

import (
	"log"
	"net/http"
)

func main() {
	const address = "localhost:8080"
	http.Handle("/", http.FileServer(http.Dir("public")))
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
