package fizzbuzz

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type FizzBuzzApi struct{}

func (f *FizzBuzzApi) ComputeFizzBuzz(w http.ResponseWriter, r *http.Request) {
	first_multiple, err := strconv.Atoi(r.URL.Query().Get("first_multiple"))
	if err != nil {
		log.Printf("Error parsing first_multiple: %s", err)
		// Error handling
		return
	}
	second_multiple, err := strconv.Atoi(r.URL.Query().Get("second_multiple"))
	if err != nil {
		log.Printf("Error parsing second_multiple: %s", err)
		// Error handling
		return
	}
	limit_integer, err := strconv.Atoi(r.URL.Query().Get("limit_integer"))
	if err != nil {
		log.Printf("Error parsing limit_integer: %s", err)
		// Error handling
		return
	}
	fizzString := r.URL.Query().Get("fizzString")
	buzzString := r.URL.Query().Get("buzzString")

	output := f.FizzBuzz(first_multiple, second_multiple, limit_integer, fizzString, buzzString)

	if err := json.NewEncoder(w).Encode(output); err != nil {
		return
	}

}

func (f *FizzBuzzApi) GetMostPopularFizzBuzz(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Not Implemented Yet")
}
