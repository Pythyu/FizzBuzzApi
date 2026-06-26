package fizzbuzz

import (
	"FizzBuzzApi/cmd/api/resource/common/helpers"
	"fmt"
	"io"
	"testing"
)

const BigRequestBodyLength = 1190744

var urlPattern = "http://0.0.0.0/v1/fizzbuzz?first_multiple=%d&second_multiple=%d&limit_integer=%d&fizzString=%s&buzzString=%s"
var classicFizzBuzzString = `["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]`

func generateURL(first_multiple, second_multiple, limit_integer int, fizzString, buzzString string) string {
	return fmt.Sprintf(urlPattern, first_multiple, second_multiple, limit_integer, fizzString, buzzString)
}

func TestFizzBuzz_BasicRequest(t *testing.T) {
	f := FizzBuzzApi{}
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
	f := FizzBuzzApi{}
	resp := helpers.MockGetRequest(generateURL(3, 5, 150000, "fuzz", "bizz"), f.ComputeFizzBuzz)

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
