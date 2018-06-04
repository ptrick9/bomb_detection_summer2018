package main


import (
	"fmt"
	"math"
)





type Square struct {
	avg float32
	numEntry int
	values []float32
	maxEntry int
	tot float32
}

func (s *Square) takeMeasurement(x float32) {
	//fmt.Printf("avg: %v num: %v vals: %v max: %v %v", s.avg, s.numEntry, s.values, s.maxEntry, x)

	if s.numEntry < s.maxEntry {
		s.tot += x
		s.values[(s.numEntry) % s.maxEntry] = x
	} else {
		s.tot -= s.values[(s.numEntry) % s.maxEntry]
		s.tot += x
		s.values[(s.numEntry) % s.maxEntry] = x
	}
	s.numEntry += 1
	s.avg = s.tot / float32(math.Min(float64(s.numEntry), float64(s.maxEntry)))

}

func (s *Square) String() string{
	return fmt.Sprintf("avg: %v num: %v vals: %v max: %v", s.avg, s.numEntry, s.values, s.maxEntry)
}
