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
func (sol *Solution) Init(input []string) {
	sol.freqChanges = make([]int, len(input))
	for i, s := range input {
		sol.freqChanges[i], _ = strconv.Atoi(s)
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := 0
	for _, f := range sol.freqChanges {
		res += f
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	buckets := map[int]bool{0: true}
	currentFreq := 0
	for {
		for _, df := range sol.freqChanges {
			currentFreq += df
			_, ok := buckets[currentFreq]
			if ok {
				return currentFreq
			}
			buckets[currentFreq] = true
		}
	}
}
