package main

import (
	"FizzBuzzApi/config"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewServerConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  cfg.TimeoutRead,
		WriteTimeout: cfg.TimeoutWrite,
		IdleTimeout:  cfg.TimeoutIdle,
	}

	log.Printf("Starting server :%d\n", cfg.Port)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World o/")
	if err != nil {
		return
	}
}
