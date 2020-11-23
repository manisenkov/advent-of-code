package day01

import (
	"strconv"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 1
type Solution struct {
	freqChanges []int
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) error {
	freqChanges := make([]int, len(input))
	for i, s := range input {
		freqChanges[i], _ = strconv.Atoi(s)
	}
	s.freqChanges = freqChanges
	return nil
}

// Part1 .
func (s *Solution) Part1() common.Any {
	resultFreq := 0
	for _, f := range s.freqChanges {
		resultFreq += f
	}
	return resultFreq
}

// Part2 .
func (s *Solution) Part2() common.Any {
	buckets := map[int]bool{0: true}
	currentFreq := 0
	for {
		for _, df := range s.freqChanges {
			currentFreq += df
			_, ok := buckets[currentFreq]
			if ok {
				return currentFreq
			}
			buckets[currentFreq] = true
		}
	}
}
