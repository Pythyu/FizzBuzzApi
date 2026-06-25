package err

import (
	"fmt"
	"net/http"
)

const jsonErrorTemplate = `{"error":"%s"}`

func ServerError(w http.ResponseWriter, errorMessage string) {
	resp := fmt.Sprintf(jsonErrorTemplate, errorMessage)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(resp))
}

func BadRequestError(w http.ResponseWriter, errorMessage string) {
	resp := fmt.Sprintf(jsonErrorTemplate, errorMessage)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(resp))
}
