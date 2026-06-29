package fizzbuzz

import "testing"

func TestFizzBuzzRequestStats_RecordRequest(t *testing.T) {
	s := FizzBuzzRequestStats{
		rqMap: make(map[FizzBuzzRequestParameters]int),
	}

	req := FizzBuzzRequestParameters{
		FirstMultiple:  3,
		SecondMultiple: 5,
		LimitInteger:   15,
		FizzString:     "Fizz",
		BuzzString:     "Buzz",
	}

	req2 := FizzBuzzRequestParameters{
		FirstMultiple:  4,
		SecondMultiple: 6,
		LimitInteger:   1500,
		FizzString:     "Fuzz",
		BuzzString:     "Bozz",
	}
	for range 3 {
		s.RecordRequest(req)
	}
	for range 2 {
		s.RecordRequest(req2)
	}

	if len(s.rqMap) != 2 {
		t.Errorf("FizzBuzzRequestStats record request map should contain 2 keys not %d", len(s.rqMap))
	}

	if s.maxRqAmount != 3 {
		t.Errorf("FizzBuzzRequestStats most used request number should be 3 not %d", s.maxRqAmount)
	}

	if s.mostPopular != req {
		t.Errorf("FizzBuzzRequestStats most popular request should be %v not %v", req, s.mostPopular)
	}
}

func TestFizzBuzzRequestStats_GetMostPopular(t *testing.T) {
	s := FizzBuzzRequestStats{
		rqMap: make(map[FizzBuzzRequestParameters]int),
	}

	req := FizzBuzzRequestParameters{
		FirstMultiple:  3,
		SecondMultiple: 5,
		LimitInteger:   15,
		FizzString:     "Fizz",
		BuzzString:     "Buzz",
	}

	req2 := FizzBuzzRequestParameters{
		FirstMultiple:  4,
		SecondMultiple: 6,
		LimitInteger:   1500,
		FizzString:     "Fuzz",
		BuzzString:     "Bozz",
	}
	for range 3 {
		s.RecordRequest(req2)
	}
	for range 2 {
		s.RecordRequest(req)
	}

	popReq, callNb := s.GetMostPopular()

	if popReq != req2 {
		t.Errorf("FizzBuzzRequestStats most popular request should be %v not %v", req2, popReq)
	}

	if callNb != 3 {
		t.Errorf("FizzBuzzRequestStats most popular request should be called 3 times not %d", callNb)
	}
}
