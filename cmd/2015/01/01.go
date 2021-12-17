package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 1
type Solution struct {
	instructions string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.instructions = input[0]
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := 0
	for _, s := range sol.instructions {
		if s == '(' {
			res++
		} else {
			res--
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	res := 0
	for i, s := range sol.instructions {
		if s == '(' {
			res++
		} else {
			res--
		}
		if res == -1 {
			return i + 1
		}
	}
	return 0
}

func main() {
	common.Run(new(Solution))
}
