package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 4
type Solution struct {
	board [][]rune
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.board = collections.Map(input, func(s string) []rune {
		return []rune(s)
	})
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	xPositions := [][2]int{}
	b := sol.board

	// Find occurences of "X" first
	for i, row := range b {
		for j, char := range row {
			if char == 'X' {
				xPositions = append(xPositions, [2]int{i, j})
			}
		}
	}

	// Search!
	res := 0
	for _, xPos := range xPositions {
		i := xPos[0]
		j := xPos[1]

		// Up
		if i >= 3 {
			// Up-left
			if j >= 3 {
				if b[i-1][j-1] == 'M' && b[i-2][j-2] == 'A' && b[i-3][j-3] == 'S' {
					res += 1
				}
			}
			// Up
			if b[i-1][j] == 'M' && b[i-2][j] == 'A' && b[i-3][j] == 'S' {
				res += 1
			}
			// Up-right
			if j < len(b[0])-3 {
				if b[i-1][j+1] == 'M' && b[i-2][j+2] == 'A' && b[i-3][j+3] == 'S' {
					res += 1
				}
			}
		}
		// Left
		if j >= 3 {
			if b[i][j-1] == 'M' && b[i][j-2] == 'A' && b[i][j-3] == 'S' {
				res += 1
			}
		}
		// Right
		if j < len(b[0])-3 {
			if b[i][j+1] == 'M' && b[i][j+2] == 'A' && b[i][j+3] == 'S' {
				res += 1
			}
		}
		// Down
		if i < len(b)-3 {
			// Down-left
			if j >= 3 {
				if b[i+1][j-1] == 'M' && b[i+2][j-2] == 'A' && b[i+3][j-3] == 'S' {
					res += 1
				}
			}
			// Down
			if b[i+1][j] == 'M' && b[i+2][j] == 'A' && b[i+3][j] == 'S' {
				res += 1
			}
			// Down-right
			if j < len(b[0])-3 {
				if b[i+1][j+1] == 'M' && b[i+2][j+2] == 'A' && b[i+3][j+3] == 'S' {
					res += 1
				}
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	b := sol.board
	res := 0
	for i := 0; i < len(b)-2; i++ {
		for j := 0; j < len(b[0])-2; j++ {
			checkDiag1 := (b[i][j] == 'M' && b[i+1][j+1] == 'A' && b[i+2][j+2] == 'S') || (b[i][j] == 'S' && b[i+1][j+1] == 'A' && b[i+2][j+2] == 'M')
			checkDiag2 := (b[i][j+2] == 'M' && b[i+1][j+1] == 'A' && b[i+2][j] == 'S') || (b[i][j+2] == 'S' && b[i+1][j+1] == 'A' && b[i+2][j] == 'M')
			if checkDiag1 && checkDiag2 {
				res += 1
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
