package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 1
type Solution struct {
	measurements []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	measurements := make([]int, len(input))
	for i, s := range input {
		measurements[i] = common.MustAtoi(s)
	}
	sol.measurements = measurements
}

func calcIncreases(measurements []int) int {
	res := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			res++
		}
	}
	return res
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	return calcIncreases(sol.measurements)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	windows := make([]int, len(sol.measurements)-2)
	for i := 0; i < len(sol.measurements)-2; i++ {
		windows[i] = sol.measurements[i] + sol.measurements[i+1] + sol.measurements[i+2]
	}
	return calcIncreases(windows)
}

func main() {
	common.Run(new(Solution))
}
