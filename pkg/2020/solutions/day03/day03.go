package day03

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 3
type Solution struct {
	treeMap [][]bool
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) error {
	treeMap := make([][]bool, len(input))
	for i, inp := range input {
		treeMap[i] = make([]bool, len(inp))
		for j, c := range inp {
			if c == '#' {
				treeMap[i][j] = true
			}
		}
	}
	s.treeMap = treeMap
	return nil
}

// Part1 .
func (s *Solution) Part1() common.Any {
	return s.checkSlope(3, 1)
}

// Part2 .
func (s *Solution) Part2() common.Any {
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
		res *= s.checkSlope(dx, dy)
	}
	return res
}

func (s *Solution) checkSlope(dx, dy int) int {
	res := 0
	w := len(s.treeMap[0])
	x := dx
	for y := dy; y < len(s.treeMap); y += dy {
		if s.treeMap[y][x] {
			res++
		}
		x = (x + dx) % w
	}
	return res
}
