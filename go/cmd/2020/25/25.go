package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

const factor = 20201227

// Solution contains solution for day 25
type Solution struct {
	publicKeys [2]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.publicKeys[0] = numbers.MustAtoi[int](input[0])
	sol.publicKeys[1] = numbers.MustAtoi[int](input[1])
}

// Part1 .
func (sol *Solution) Part1() any {
	lc := 1
	p := 1
	for {
		p = (p * 7) % factor
		if p == sol.publicKeys[0] {
			break
		}
		lc++
	}
	p = 1
	for i := 0; i < lc; i++ {
		p = (p * sol.publicKeys[1]) % factor
	}
	return p
}

// Part2 .
func (sol *Solution) Part2() any {
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
