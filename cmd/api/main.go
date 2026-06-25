package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", sayHello)

	s := &http.Server{
		Addr:         ":8989",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	log.Println("Starting server :8989")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World o/")
	if err != nil {
		return
	}
}
