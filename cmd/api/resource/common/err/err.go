package err

import (
	"FizzBuzzApi/cmd/api/resource/common/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidatorResponse struct {
	Errors []string `json:"errors"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Error: message,
	})
}

func ServerError(w http.ResponseWriter, errorMessage string) {
	WriteError(w, http.StatusInternalServerError, errorMessage)
}

func BadRequestError(w http.ResponseWriter, errorMessage string) {
	WriteError(w, http.StatusBadRequest, errorMessage)
}

func BadRequestValidatorError(w http.ResponseWriter, validatorError validator.ValidationErrors) {
	resp := ValidatorResponse{Errors: make([]string, len(validatorError))}

	for idx, errorValue := range validatorError {
		switch errorValue.ActualTag() {
		case "required":
			resp.Errors[idx] = fmt.Sprintf("%s is a required field", errorValue.Field())
		case "gt":
			resp.Errors[idx] = fmt.Sprintf("%s must be greater than %s", errorValue.Field(), errorValue.Param())
		case "gte":
			resp.Errors[idx] = fmt.Sprintf("%s must be greater or equal than %s", errorValue.Field(), errorValue.Param())
		case "lte":
			resp.Errors[idx] = fmt.Sprintf("%s must be lower or equal than %s", errorValue.Field(), errorValue.Param())
		default:
			resp.Errors[idx] = fmt.Sprintf("something wrong on %s; %s", errorValue.Field(), errorValue.Tag())
		}
	}
	respBody, err := json.Marshal(resp)
	if err != nil {
		ServerError(w, "Error serializing validator errors")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	helpers.SafeWrite(w, respBody)
}
