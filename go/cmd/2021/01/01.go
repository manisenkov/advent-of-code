package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 1
type Solution struct {
	measurements []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	measurements := make([]int, len(input))
	for i, s := range input {
		measurements[i] = numbers.MustAtoi[int](s)
	}
	sol.measurements = measurements
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for i := 0; i < len(sol.measurements)-1; i++ {
		if sol.measurements[i+1] > sol.measurements[i] {
			res++
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	res := 0
	for i := 0; i < len(sol.measurements)-3; i++ {
		if sol.measurements[i+3] > sol.measurements[i] {
			res++
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
