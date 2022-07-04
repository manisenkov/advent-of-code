package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 20
type Solution struct {
	enhancer      map[int]bool
	cache         map[[3]int]bool
	width, height int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.enhancer = make(map[int]bool)
	sol.cache = map[[3]int]bool{}
	for i, b := range input[0] {
		if b == '#' {
			sol.enhancer[i] = true
		}
	}
	for row, s := range input[2:] {
		for col, t := range s {
			if t == '#' {
				sol.cache[[3]int{0, row, col}] = true
			}
		}
	}
	sol.height = len(input[2:])
	sol.width = len(input[2:][0])
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.enhance(2)
}

// Part2 .
func (sol *Solution) Part2() any {
	return sol.enhance(50)
}

func (sol *Solution) enhance(steps int) int {
	minRow := -100
	maxRow := sol.height + 100
	minCol := -100
	maxCol := sol.width + 100
	res := 0
	for row := minRow; row < maxRow; row++ {
		for col := minCol; col < maxCol; col++ {
			if sol.calcPoint(steps, row, col) {
				res++
			}
		}
	}
	return res
}

func (sol *Solution) calcPoint(step, row, col int) bool {
	val, ok := sol.cache[[3]int{step, row, col}]
	if ok {
		return val
	} else if step == 0 {
		return false
	}

	enhancerIndex := (boolToBit(sol.calcPoint(step-1, row-1, col-1)) << 8) +
		(boolToBit(sol.calcPoint(step-1, row-1, col)) << 7) +
		(boolToBit(sol.calcPoint(step-1, row-1, col+1)) << 6) +
		(boolToBit(sol.calcPoint(step-1, row, col-1)) << 5) +
		(boolToBit(sol.calcPoint(step-1, row, col)) << 4) +
		(boolToBit(sol.calcPoint(step-1, row, col+1)) << 3) +
		(boolToBit(sol.calcPoint(step-1, row+1, col-1)) << 2) +
		(boolToBit(sol.calcPoint(step-1, row+1, col)) << 1) +
		(boolToBit(sol.calcPoint(step-1, row+1, col+1)) << 0)
	val = sol.enhancer[enhancerIndex]
	sol.cache[[3]int{step, row, col}] = val
	return val
}

func boolToBit(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func main() {
	common.Run(new(Solution))
}
