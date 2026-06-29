package main

import (
	"FizzBuzzApi/cmd/api/config"
	"FizzBuzzApi/cmd/api/router"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// @title FizzBuzz API
// @version 1.0
// @description An API to generate FizzBuzz that will Fizz and Buzz however you want

// @contact.name Fizz Buzzer
// @contact.email some.fake@gmail.com

// @BasePath /v1
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
