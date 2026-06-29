package router

import (
	"FizzBuzzApi/cmd/api/resource/fizzbuzz"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/go-playground/validator/v10"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	v := validator.New()

	r.Use(httprate.LimitByIP(60, time.Minute))

	r.Route("/v1", func(r chi.Router) {
		FBApi := fizzbuzz.NewFizzBuzzApi(v)
		r.Get("/fizzbuzz", FBApi.ComputeFizzBuzz)
		r.Get("/stats", FBApi.GetMostPopularFizzBuzz)
	})

	return r
}
