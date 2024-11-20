package main

import (
	"math"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 5
type Solution struct {
	seatIDs   []int
	minSeatID int
	maxSeatID int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.seatIDs = make([]int, len(input))
	sol.minSeatID = math.MaxInt32
	sol.maxSeatID = 0
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
		seatID := numbers.MustParseInt[int](idStr, 2)
		sol.seatIDs[i] = seatID
		if seatID > sol.maxSeatID {
			sol.maxSeatID = seatID
		}
		if seatID < sol.minSeatID {
			sol.minSeatID = seatID
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.maxSeatID
}

// Part2 .
func (sol *Solution) Part2() any {
	taken := make(map[int]bool)
	for _, seatID := range sol.seatIDs {
		taken[seatID] = true
	}
	for id := sol.minSeatID + 1; id < sol.maxSeatID; id++ {
		if !taken[id] {
			return id
		}
	}
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
