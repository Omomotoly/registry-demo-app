package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := fmt.Sprintf("Hello from Go Web App Demo for DockerHub and GHCR %s\n", hostname)
		fmt.Fprintf(w, response)
		log.Printf("Request received from %s", r.RemoteAddr)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
		log.Printf("Health check request received from %s", r.RemoteAddr)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "2222"
	}

	log.Printf("Starting server on port %s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}