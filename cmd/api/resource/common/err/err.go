package err

import (
	"FizzBuzzApi/cmd/api/resource/common/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const jsonErrorTemplate = `{"error":"%s"}`

type ErrValidatorResponse struct {
	Errors []string `json:"errors"`
}

func ServerError(w http.ResponseWriter, errorMessage string) {
	resp := fmt.Sprintf(jsonErrorTemplate, errorMessage)
	w.WriteHeader(http.StatusInternalServerError)
	helpers.SafeWrite(w, []byte(resp))
}

func BadRequestError(w http.ResponseWriter, errorMessage string) {
	resp := fmt.Sprintf(jsonErrorTemplate, errorMessage)
	w.WriteHeader(http.StatusBadRequest)
	helpers.SafeWrite(w, []byte(resp))
}

func BadRequestValidatorError(w http.ResponseWriter, validatorError validator.ValidationErrors) {
	resp := ErrValidatorResponse{Errors: make([]string, len(validatorError))}

	for idx, errorValue := range validatorError {
		switch errorValue.ActualTag() {
		case "required":
			resp.Errors[idx] = fmt.Sprintf("%s is a required field", errorValue.Field())
		case "gt":
			resp.Errors[idx] = fmt.Sprintf("%s must be greater than %s", errorValue.Field(), errorValue.Param())
		case "gte":
			resp.Errors[idx] = fmt.Sprintf("%s must be greater or equal than %s", errorValue.Field(), errorValue.Param())
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
