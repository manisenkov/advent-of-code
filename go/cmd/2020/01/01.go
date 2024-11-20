package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 1
type Solution struct {
	entries []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.entries = make([]int, len(input))
	for i, inp := range input {
		sol.entries[i] = numbers.MustAtoi[int](inp)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	for i := 0; i < len(sol.entries)-1; i++ {
		for j := i + 1; j < len(sol.entries); j++ {
			if sol.entries[i]+sol.entries[j] == 2020 {
				return sol.entries[i] * sol.entries[j]
			}
		}
	}
	return 0
}

// Part2 .
func (sol *Solution) Part2() any {
	for i := 0; i < len(sol.entries)-2; i++ {
		for j := i + 1; j < len(sol.entries)-1; j++ {
			for k := j + 1; k < len(sol.entries); k++ {
				if sol.entries[i]+sol.entries[j]+sol.entries[k] == 2020 {
					return sol.entries[i] * sol.entries[j] * sol.entries[k]
				}
			}
		}
	}
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
