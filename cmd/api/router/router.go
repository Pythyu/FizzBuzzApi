package router

import (
	"FizzBuzzApi/cmd/api/resource/fizzbuzz"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		FBApi := &fizzbuzz.FizzBuzzApi{}
		r.Get("/fizzbuzz", FBApi.ComputeFizzBuzz)
		r.Get("/stats", FBApi.GetMostPopularFizzBuzz)
	})

	return r
}
