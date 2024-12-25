package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func parseLock(input []string) [5]int {
	res := [5]int{}
	for _, line := range input[1:] {
		for j := 0; j < 5; j++ {
			if line[j] == '#' {
				res[j]++
			}
		}
	}
	return res
}

func parseKey(input []string) [5]int {
	res := [5]int{}
	for _, line := range input[0:6] {
		for j := 0; j < 5; j++ {
			if line[j] == '#' {
				res[j]++
			}
		}
	}
	return res
}

// Solution contains a solution for day 25
type Solution struct {
	locks [][5]int
	keys  [][5]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	i := 0
	for i < len(input) {
		if input[i] == "#####" {
			sol.locks = append(sol.locks, parseLock(input[i:i+6]))
		} else {
			sol.keys = append(sol.keys, parseKey(input[i:i+6]))
		}
		i += 8
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, lock := range sol.locks {
		for _, key := range sol.keys {
			fit := true
			for j := 0; j < 5; j++ {
				if lock[j]+key[j] > 5 {
					fit = false
					break
				}
			}
			if fit {
				res++
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return "HO HO HO"
}

func main() {
	problem.Solve(new(Solution))
}
