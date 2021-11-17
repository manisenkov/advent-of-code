package main

import (
	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

const factor = 20201227

// Solution contains solution for day 25
type Solution struct {
	publicKeys [2]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.publicKeys[0] = common.MustAtoi(input[0])
	sol.publicKeys[1] = common.MustAtoi(input[1])
}

// Part1 .
func (sol *Solution) Part1() common.Any {
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
func (sol *Solution) Part2() common.Any {
	return 0
}

func main() {
	common.Run(new(Solution))
}
