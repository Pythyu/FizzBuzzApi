package fizzbuzz

import "strconv"

func isMultipleOf(number int, mult1 int, mult2 int) int {
	idx := 0
	if (number % mult1) == 0 {
		idx += 1
	}
	if (number % mult2) == 0 {
		idx += 2
	}
	return idx
}

func (f *FizzBuzzApi) FizzBuzz(firstMult int, secondMult int, start int, limit int, fizzStr string, buzzStr string) []string {
	var output []string
	var pattern = [4]string{"", fizzStr, buzzStr, fizzStr + buzzStr}
	for i := start; i <= limit; i += 1 {
		patternIndex := isMultipleOf(i, firstMult, secondMult)
		if patternIndex > 0 {
			output = append(output, pattern[patternIndex])
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}

	return output
}
