package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func turn(grid [10][10]int) ([10][10]int, int) {
	nextGrid := grid
	flashesQueue := [][2]int{}

	for row := range nextGrid {
		for col := range nextGrid {
			nextGrid[row][col] = (nextGrid[row][col] + 1) % 10
			if nextGrid[row][col] == 0 {
				flashesQueue = append(flashesQueue, [2]int{row, col})
			}
		}
	}

	totalFlashes := 0
	for len(flashesQueue) > 0 {
		row := flashesQueue[0][0]
		col := flashesQueue[0][1]
		flashesQueue = flashesQueue[1:]
		totalFlashes++
		adj := [][2]int{}
		if row > 0 && col > 0 {
			adj = append(adj, [2]int{row - 1, col - 1})
		}
		if row > 0 {
			adj = append(adj, [2]int{row - 1, col})
		}
		if row > 0 && col < 9 {
			adj = append(adj, [2]int{row - 1, col + 1})
		}
		if col > 0 {
			adj = append(adj, [2]int{row, col - 1})
		}
		if col < 9 {
			adj = append(adj, [2]int{row, col + 1})
		}
		if row < 9 && col > 0 {
			adj = append(adj, [2]int{row + 1, col - 1})
		}
		if row < 9 {
			adj = append(adj, [2]int{row + 1, col})
		}
		if row < 9 && col < 9 {
			adj = append(adj, [2]int{row + 1, col + 1})
		}
		for _, x := range adj {
			adjRow := x[0]
			adjCol := x[1]
			if nextGrid[adjRow][adjCol] == 0 {
				continue
			}
			nextGrid[adjRow][adjCol] = (nextGrid[adjRow][adjCol] + 1) % 10
			if nextGrid[adjRow][adjCol] == 0 {
				flashesQueue = append(flashesQueue, [2]int{adjRow, adjCol})
			}
		}
	}
	return nextGrid, totalFlashes
}

// Solution contains solution for day 11
type Solution struct {
	grid [10][10]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	for row, s := range input {
		for col, r := range s {
			sol.grid[row][col] = int(r - '0')
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	grid := sol.grid
	for i := 0; i < 100; i++ {
		var flashCount int
		grid, flashCount = turn(grid)
		res += flashCount
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	grid := sol.grid
	step := 1
	for {
		var flashCount int
		grid, flashCount = turn(grid)
		if flashCount == 100 {
			return step
		}
		step++
	}
}

func main() {
	problem.Solve(new(Solution))
}
