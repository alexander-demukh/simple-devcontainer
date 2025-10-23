package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	port := ":8000"
	if len(os.Args) > 1 {
		port = ":" + os.Args[1]
	}

	http.HandleFunc("/ping", pingHandler)

	// Start background goroutine to print message every 5 seconds
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			fmt.Printf("App %s: I'm working\n", port[1:])
		}
	}()

	log.Printf("Server starting on port %s", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
