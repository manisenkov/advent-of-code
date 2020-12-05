package day05

import (
	"math"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 5
type Solution struct {
	seatIDs   []int
	minSeatID int
	maxSeatID int
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
	seatIDs := make([]int, len(input))
	minSeatID := math.MaxInt32
	maxSeatID := 0
	for i, inp := range input {
		idStr := ""
		for _, c := range inp {
			switch c {
			case 'F':
				idStr += "0"
			case 'B':
				idStr += "1"
			case 'L':
				idStr += "0"
			case 'R':
				idStr += "1"
			}
		}
		seatID := int(common.MustParseInt(idStr, 2, 32))
		seatIDs[i] = seatID
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
		if seatID < minSeatID {
			minSeatID = seatID
		}
	}
	s.seatIDs = seatIDs
	s.minSeatID = minSeatID
	s.maxSeatID = maxSeatID
}

// Part1 .
func (s *Solution) Part1() common.Any {
	return s.maxSeatID
}

// Part2 .
func (s *Solution) Part2() common.Any {
	taken := make(map[int]bool)
	for _, seatID := range s.seatIDs {
		taken[seatID] = true
	}
	for id := s.minSeatID + 1; id < s.maxSeatID; id++ {
		if !taken[id] {
			return id
		}
	}
	return 0
}
