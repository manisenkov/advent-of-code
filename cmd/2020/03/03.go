package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 3
type Solution struct {
	treeMap [][]bool
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.treeMap = make([][]bool, len(input))
	for i, inp := range input {
		sol.treeMap[i] = make([]bool, len(inp))
		for j, c := range inp {
			if c == '#' {
				sol.treeMap[i][j] = true
			}
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	return sol.checkSlope(3, 1)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	res := 1
	for _, slope := range slopes {
		dx := slope[0]
		dy := slope[1]
		res *= sol.checkSlope(dx, dy)
	}
	return res
}

func (sol *Solution) checkSlope(dx, dy int) int {
	res := 0
	w := len(sol.treeMap[0])
	x := dx
	for y := dy; y < len(sol.treeMap); y += dy {
		if sol.treeMap[y][x] {
			res++
		}
		x = (x + dx) % w
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
