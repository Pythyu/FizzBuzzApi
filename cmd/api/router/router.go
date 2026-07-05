package router

import (
	_ "FizzBuzzApi/cmd/api/docs"
	"FizzBuzzApi/cmd/api/resource/fizzbuzz"
	"reflect"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	v := validator.New()

	// Report field names using the schema struct tag
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		if name := fld.Tag.Get("schema"); name != "" {
			return name
		}
		return fld.Name
	})

	r.Use(httprate.LimitByIP(60, time.Minute))

	r.Route("/v1", func(r chi.Router) {
		FBApi := fizzbuzz.NewFizzBuzzApi(v)
		r.Get("/fizzbuzz", FBApi.ComputeFizzBuzz)
		r.Get("/stats", FBApi.GetMostPopularFizzBuzz)
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
