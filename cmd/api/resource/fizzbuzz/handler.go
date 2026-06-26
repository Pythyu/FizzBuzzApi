package fizzbuzz

import (
	errorHandler "FizzBuzzApi/cmd/api/resource/common/err"
	helper "FizzBuzzApi/cmd/api/resource/common/helpers"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

const outputBlockSize = 1024

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
	w.Header().Set("Content-Type", "application/json")
	helper.SafeWrite(w, []byte("["))
	blockNumber := (limit_integer-1)/outputBlockSize + 1
	for i := range blockNumber {
		start := i*outputBlockSize + 1
		var end int
		if i+1 == blockNumber {
			end = limit_integer
		} else {
			end = i*outputBlockSize + outputBlockSize
		}
		output := f.FizzBuzz(first_multiple, second_multiple, start, end, fizzString, buzzString)
		bytes, err := json.Marshal(output)
		if err != nil {
			errorHandler.ServerError(w, "Failed to marshal JSON")
			return
		}
		if start != 1 {
			helper.SafeWrite(w, []byte(","))
		}
		helper.SafeWrite(w, bytes[1:len(bytes)-1])
	}
	helper.SafeWrite(w, []byte("]"))

}

func (f *FizzBuzzApi) GetMostPopularFizzBuzz(w http.ResponseWriter, r *http.Request) {

	_, err := io.WriteString(w, "Not Implemented Yet")
	if err != nil {
		return
	}
}
