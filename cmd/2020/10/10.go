package main

import (
	"sort"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 10
type Solution struct {
	jolts []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.jolts = make([]int, len(input)+2)
	sol.jolts[0] = 0
	maxJolt := 0
	for i, inp := range input {
		jolt := common.MustAtoi(inp)
		if jolt > maxJolt {
			maxJolt = jolt
		}
		sol.jolts[i+1] = jolt
	}
	sol.jolts[len(sol.jolts)-1] = maxJolt + 3
	sort.Slice(sol.jolts, func(i, j int) bool {
		return sol.jolts[i] < sol.jolts[j]
	})
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	curJolt := 0
	diff1Count := 0
	diff3Count := 0
	for _, jolt := range sol.jolts {
		if jolt-curJolt == 1 {
			diff1Count++
		}
		if jolt-curJolt == 3 {
			diff3Count++
		}
		curJolt = jolt
	}
	return diff1Count * diff3Count
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	i := 0
	j := i + 1
	res := 1
	for j < len(sol.jolts) {
		if sol.jolts[j]-sol.jolts[j-1] == 3 {
			res *= sol.countArrangements(i, j)
			i = j
			j = i + 1
			continue
		}
		j++
	}
	return res
}

func (sol *Solution) countArrangements(startIndex, endIndex int) int {
	if startIndex == endIndex {
		return 1
	}
	res := 0
	for i := startIndex + 1; i <= endIndex && sol.jolts[i]-sol.jolts[startIndex] <= 3; i++ {
		res += sol.countArrangements(i, endIndex)
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
