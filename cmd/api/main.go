package main

import (
	"FizzBuzzApi/cmd/api/router"
	"FizzBuzzApi/config"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewServerConfig()
	apiRouter := router.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      apiRouter,
		ReadTimeout:  cfg.TimeoutRead,
		WriteTimeout: cfg.TimeoutWrite,
		IdleTimeout:  cfg.TimeoutIdle,
	}

	log.Printf("Starting server :%d\n", cfg.Port)
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server startup failed")
	}
}
