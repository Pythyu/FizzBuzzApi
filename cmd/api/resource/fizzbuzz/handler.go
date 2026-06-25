package fizzbuzz

import (
	errorHandler "FizzBuzzApi/cmd/api/resource/common/err"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type FizzBuzzApi struct{}

func (f *FizzBuzzApi) ComputeFizzBuzz(w http.ResponseWriter, r *http.Request) {
	first_multiple, err := strconv.Atoi(r.URL.Query().Get("first_multiple"))
	if err != nil {
		errorHandler.BadRequestError(w, "Invalid field first_multiple")
		return
	}
	second_multiple, err := strconv.Atoi(r.URL.Query().Get("second_multiple"))
	if err != nil {
		errorHandler.BadRequestError(w, "Invalid field second_multiple")
		return
	}
	limit_integer, err := strconv.Atoi(r.URL.Query().Get("limit_integer"))
	if err != nil {
		errorHandler.BadRequestError(w, "Invalid field limit_integer")
		return
	}
	fizzString := r.URL.Query().Get("fizzString")
	if fizzString == "" {
		errorHandler.BadRequestError(w, "Invalid field fizzString")
		return
	}
	buzzString := r.URL.Query().Get("buzzString")
	if buzzString == "" {
		errorHandler.BadRequestError(w, "Invalid field buzzString")
		return
	}

	output := f.FizzBuzz(first_multiple, second_multiple, limit_integer, fizzString, buzzString)

	if err := json.NewEncoder(w).Encode(output); err != nil {
		errorHandler.ServerError(w, "json encode failure")
		return
	}

}

func (f *FizzBuzzApi) GetMostPopularFizzBuzz(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Not Implemented Yet")
}
