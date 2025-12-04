package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 4
type Solution struct {
	plan [][]bool
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.plan = make([][]bool, len(input))
	for i, s := range input {
		sol.plan[i] = make([]bool, len(s))
		for j, c := range s {
			if c == '@' {
				sol.plan[i][j] = true
			}
		}
	}
}

func countNeighbors(plan [][]bool, row, col int) int {
	res := 0
	if row >= 1 {
		if col >= 1 && plan[row-1][col-1] {
			res++
		}
		if plan[row-1][col] {
			res++
		}
		if col < len(plan[0])-1 && plan[row-1][col+1] {
			res++
		}
	}
	if col >= 1 && plan[row][col-1] {
		res++
	}
	if col < len(plan[0])-1 && plan[row][col+1] {
		res++
	}
	if row < len(plan)-1 {
		if col >= 1 && plan[row+1][col-1] {
			res++
		}
		if plan[row+1][col] {
			res++
		}
		if col < len(plan[0])-1 && plan[row+1][col+1] {
			res++
		}
	}

	return res
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for row := range len(sol.plan) {
		for col := range len(sol.plan[0]) {
			if sol.plan[row][col] && countNeighbors(sol.plan, row, col) < 4 {
				res++
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	plan := make([][]bool, len(sol.plan))
	for i, r := range sol.plan {
		plan[i] = make([]bool, len(r))
		copy(plan[i], r)
	}
	res := 0
	toRemove := [][2]int{}
	for {
		for row := range len(sol.plan) {
			for col := range len(sol.plan[0]) {
				if plan[row][col] && countNeighbors(plan, row, col) < 4 {
					toRemove = append(toRemove, [2]int{row, col})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, pos := range toRemove {
			plan[pos[0]][pos[1]] = false
		}
		res += len(toRemove)
		toRemove = [][2]int{}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
