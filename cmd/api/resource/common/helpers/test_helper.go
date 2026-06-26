package helpers

import (
	"net/http"
	"net/http/httptest"
)

func MockGetRequest(fullURL string, handler func(w http.ResponseWriter, r *http.Request)) *http.Response {
	req := httptest.NewRequest("GET", fullURL, nil)
	w := httptest.NewRecorder()

	handler(w, req)
	return w.Result()
}
