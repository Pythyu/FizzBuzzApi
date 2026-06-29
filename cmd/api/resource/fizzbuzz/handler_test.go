package fizzbuzz

import (
	"FizzBuzzApi/cmd/api/resource/common/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
)

const BigRequestBodyLength = 384076

var urlPattern = "http://0.0.0.0/v1/fizzbuzz?first_multiple=%d&second_multiple=%d&limit_integer=%d&fizzString=%s&buzzString=%s"
var classicFizzBuzzString = `["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]`

func generateURL(first_multiple, second_multiple, limit_integer int, fizzString, buzzString string) string {
	return fmt.Sprintf(urlPattern, first_multiple, second_multiple, limit_integer, fizzString, buzzString)
}

type errorJSON struct {
	Errors []string `json:"errors"`
}

func TestFizzBuzz_BasicRequest(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(3, 5, 15, "fizz", "buzz"), f.ComputeFizzBuzz)

	if resp.StatusCode != 200 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	body, _ := io.ReadAll(resp.Body)

	if string(body) != classicFizzBuzzString {
		t.Fatalf("Unexpected body: %s", string(body))
	}
}

func TestFizzBuzz_BigRequest(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(3, 5, 50000, "fuzz", "bizz"), f.ComputeFizzBuzz)

	if resp.StatusCode != 200 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	body, _ := io.ReadAll(resp.Body)

	if len(body) != BigRequestBodyLength {
		t.Fatalf("Unexpected body length: %d", len(body))
	}
}

func TestFizzBuzz_TooBigRequest(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(3, 5, 5000000, "fuzz", "bizz"), f.ComputeFizzBuzz)

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	var errors errorJSON

	err := json.NewDecoder(resp.Body).Decode(&errors)
	if err != nil {
		t.Fatalf("Unable to parse the request's error: %s", err)
	}
	if len(errors.Errors) != 1 {
		t.Fatalf("Unexpected number of errors: %d", len(errors.Errors))
	}
	if errors.Errors[0] != "LimitInteger must be lower or equal than 50000" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
}

func TestFizzBuzz_BadRequest1(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(-1, 5, 15, "fizz", "buzz"), f.ComputeFizzBuzz)

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	var errors errorJSON

	err := json.NewDecoder(resp.Body).Decode(&errors)
	if err != nil {
		t.Fatalf("Unable to parse the request's error: %s", err)
	}
	if len(errors.Errors) != 1 {
		t.Fatalf("Unexpected number of errors: %d", len(errors.Errors))
	}
	if errors.Errors[0] != "FirstMultiple must be greater than 0" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
}

func TestFizzBuzz_BadRequestAllWrong(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(-1, -4, -2, "", ""), f.ComputeFizzBuzz)

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	var errors errorJSON

	err := json.NewDecoder(resp.Body).Decode(&errors)
	if err != nil {
		t.Fatalf("Unable to parse the request's error: %s", err)
	}
	if len(errors.Errors) != 5 {
		t.Fatalf("Unexpected number of errors: %d", len(errors.Errors))
	}
	if errors.Errors[0] != "FirstMultiple must be greater than 0" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[1] != "SecondMultiple must be greater than 0" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[2] != "LimitInteger must be greater or equal than 1" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[3] != "FizzString is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[4] != "BuzzString is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
}

func TestFizzBuzz_BadRequestAllNull(t *testing.T) {
	v := validator.New()
	f := NewFizzBuzzApi(v)
	resp := helpers.MockGetRequest(generateURL(0, 0, 0, "", ""), f.ComputeFizzBuzz)

	if resp.StatusCode != 400 {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Unexpected content type: %s", resp.Header.Get("Content-Type"))
	}

	var errors errorJSON

	err := json.NewDecoder(resp.Body).Decode(&errors)
	if err != nil {
		t.Fatalf("Unable to parse the request's error: %s", err)
	}
	if len(errors.Errors) != 5 {
		t.Fatalf("Unexpected number of errors: %d", len(errors.Errors))
	}
	if errors.Errors[0] != "FirstMultiple is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[1] != "SecondMultiple is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[2] != "LimitInteger is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[3] != "FizzString is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
	if errors.Errors[4] != "BuzzString is a required field" {
		t.Fatalf("Unexpected error message: %s", errors.Errors[0])
	}
}

func BenchmarkRequests(b *testing.B) {
	if os.Getenv("CI") != "" {
		b.Skip("Skipping because benchmark are not meant to run on CI")
	}
	fullURL := generateURL(3, 5, 5000, "fizz", "buzz")
	v := validator.New()
	f := NewFizzBuzzApi(v)
	b.RunParallel(func(pb *testing.PB) {
		req := httptest.NewRequest("GET", fullURL, nil)
		w := httptest.NewRecorder()

		b.ReportAllocs()
		b.ResetTimer()

		for pb.Next() {
			f.ComputeFizzBuzz(w, req)
		}
	})
}
