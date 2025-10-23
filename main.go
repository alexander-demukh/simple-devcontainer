package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", pingHandler)

	// Start background goroutine to print message every 5 seconds
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			fmt.Println("I'm working")
		}
	}()

	port := ":8000"
	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
