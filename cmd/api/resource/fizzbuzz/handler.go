package fizzbuzz

import (
	errorHandler "FizzBuzzApi/cmd/api/resource/common/err"
	helper "FizzBuzzApi/cmd/api/resource/common/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

const outputBlockSize = 1024

type FizzBuzzApi struct {
	validate     *validator.Validate
	requestStats *FizzBuzzRequestStats
}

func NewFizzBuzzApi(validate *validator.Validate) *FizzBuzzApi {
	return &FizzBuzzApi{
		validate: validate,
		requestStats: &FizzBuzzRequestStats{
			rqMap: make(map[FizzBuzzRequestParameters]int),
		},
	}
}

type FizzBuzzRequestParameters struct {
	FirstMultiple  int    `schema:"first_multiple" validate:"required,gt=0"`
	SecondMultiple int    `schema:"second_multiple" validate:"required,gt=0"`
	LimitInteger   int    `schema:"limit_integer" validate:"required,gte=1,lte=50000"`
	FizzString     string `schema:"fizzString" validate:"required"`
	BuzzString     string `schema:"buzzString" validate:"required"`
}

// ComputeFizzBuzz generate a custom FizzBuzz sequence based on given parameters
//
// @Summary               Compute the FizzBuzz sequence
// @Description           Returns a fully customizable FizzBuzz sequence
// @Tags                  FizzBuzz
// @Produce               json
// @Param first_multiple  query int true "First multiple for fizz"
// @Param second_multiple query int true "Second multiple for buzz"
// @Param limit_integer   query int true "Limit number of the sequence"
// @Param fizzString      query string true "Fizz replacement"
// @Param buzzString      query string true "Buzz replacement"
// @Success               200 {array} string
// @Failure               400 {object} errorHandler.ErrorResponse
// @Router                /fizzbuzz [get]
func (f *FizzBuzzApi) ComputeFizzBuzz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req FizzBuzzRequestParameters

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&req, r.URL.Query()); err != nil {
		errorHandler.BadRequestError(w, err.Error())
		return
	}

	if err := f.validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorHandler.BadRequestValidatorError(w, validationErrors)
		} else {
			errorHandler.BadRequestError(w, err.Error())
		}
		return
	}

	f.requestStats.RecordRequest(req)

	helper.SafeWrite(w, []byte("["))
	blockNumber := (req.LimitInteger-1)/outputBlockSize + 1
	for i := range blockNumber {
		start := i*outputBlockSize + 1
		var end int
		if i+1 == blockNumber {
			end = req.LimitInteger
		} else {
			end = i*outputBlockSize + outputBlockSize
		}
		output := f.FizzBuzz(req.FirstMultiple, req.SecondMultiple, start, end, req.FizzString, req.BuzzString)
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

type mostPopularJsonResponse struct {
	Parameters FizzBuzzRequestParameters `json:"parameters"`
	CallNumber int                       `json:"call_number"`
}

// GetMostPopularFizzBuzz returns the most frequently requested FizzBuzz parameters.
//
// @Summary      Get FizzBuzz statistics
// @Description  Returns the parameters corresponding to the most frequently requested FizzBuzz computation and the number of times it has been called.
// @Tags         Statistics
// @Produce      json
// @Success      200 {object} mostPopularJsonResponse
// @Failure      500 {object} errorHandler.ErrorResponse
// @Router       /stats [get]
func (f *FizzBuzzApi) GetMostPopularFizzBuzz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mostPopular, callAmount := f.requestStats.GetMostPopular()
	resp := mostPopularJsonResponse{mostPopular, callAmount}
	bytes, err := json.Marshal(resp)
	if err != nil {
		errorHandler.ServerError(w, "Failed to marshal JSON")
		return
	}
	helper.SafeWrite(w, bytes)

}
