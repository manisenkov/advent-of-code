package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 1
type Solution struct {
	input string
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.input = input[0]
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	sz := len(sol.input)
	for i := 0; i < sz; i++ {
		x := sol.input[i]
		t := sol.input[(i+1)%sz]
		if x == t {
			res += int(x - '0')
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	sz := len(sol.input)
	for i := 0; i < len(sol.input); i++ {
		x := sol.input[i]
		t := sol.input[(i+sz/2)%sz]
		if x == t {
			res += int(x - '0')
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
