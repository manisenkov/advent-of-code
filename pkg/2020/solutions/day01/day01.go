package day01

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 1
type Solution struct {
	entries []int
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) error {
	entries := make([]int, len(input))
	for i, inp := range input {
		entries[i] = common.MustAtoi(inp)
	}
	s.entries = entries
	return nil
}

// Part1 .
func (s *Solution) Part1() common.Any {
	for i := 0; i < len(s.entries)-1; i++ {
		for j := i + 1; j < len(s.entries); j++ {
			if s.entries[i]+s.entries[j] == 2020 {
				return s.entries[i] * s.entries[j]
			}
		}
	}
	return 0
}

// Part2 .
func (s *Solution) Part2() common.Any {
	for i := 0; i < len(s.entries)-2; i++ {
		for j := i + 1; j < len(s.entries)-1; j++ {
			for k := j + 1; k < len(s.entries); k++ {
				if s.entries[i]+s.entries[j]+s.entries[k] == 2020 {
					return s.entries[i] * s.entries[j] * s.entries[k]
				}
			}
		}
	}
	return 0
}
