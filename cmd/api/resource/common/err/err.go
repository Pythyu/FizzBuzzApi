package err

import (
	"FizzBuzzApi/cmd/api/resource/common/helpers"
	"fmt"
	"net/http"
)

const jsonErrorTemplate = `{"error":"%s"}`

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
