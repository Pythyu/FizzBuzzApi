package fizzbuzz

import (
	"reflect"
	"testing"
)

var classicFizzBuzz = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}

func TestFizzBuzz_ClassicFizzBuzz(t *testing.T) {
	f := FizzBuzzApi{}
	output := f.FizzBuzz(3, 5, 1, 15, "Fizz", "Buzz")
	if !reflect.DeepEqual(output, classicFizzBuzz) {
		t.Fatalf("Invalid ClassicFizzBuzz, returned %s expected %s", output, classicFizzBuzz)
	}
}
