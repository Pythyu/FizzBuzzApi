package fizzbuzz

import "sync"

type FizzBuzzRequestStats struct {
	mu    sync.Mutex
	rqMap map[FizzBuzzRequestParameters]int

	mostPopular FizzBuzzRequestParameters
	maxRqAmount int
}

func (s *FizzBuzzRequestStats) RecordRequest(req FizzBuzzRequestParameters) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rqMap[req] += 1

	if s.rqMap[req] > s.maxRqAmount {
		s.maxRqAmount = s.rqMap[req]
		s.mostPopular = req
	}
}

func (s *FizzBuzzRequestStats) GetMostPopular() (FizzBuzzRequestParameters, int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.mostPopular, s.maxRqAmount
}
