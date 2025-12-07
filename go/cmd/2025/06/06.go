package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 6
type Solution struct {
	numbers   [][]string
	operators []string
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	// Finding vertical spaces first
	spaceCols := []int{}
	maxLineLength := 0
	for _, s := range input {
		maxLineLength = numbers.Max(maxLineLength, len(s))
	}
	for col := range maxLineLength {
		isEmptyCol := true
		for _, s := range input {
			if len(s)-1 >= col && s[col] != ' ' {
				isEmptyCol = false
				break
			}
		}
		if isEmptyCol {
			spaceCols = append(spaceCols, col)
		}
	}
	spaceCols = append(spaceCols, maxLineLength)

	sol.numbers = make([][]string, len(input)-1)
	for row, s := range input[:len(input)-1] {
		for len(s) < maxLineLength {
			s += " "
		}
		j := 0
		for _, col := range spaceCols {
			sol.numbers[row] = append(sol.numbers[row], s[j:col])
			j = col + 1
		}
	}

	sol.operators = []string{}
	j := 0
	for _, col := range spaceCols {
		s := input[len(input)-1]
		op := strings.Trim(s[j:numbers.Min(col, len(s))], " ")
		sol.operators = append(sol.operators, op)
		j = col + 1
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for i, op := range sol.operators {
		var s int
		if op == "*" {
			s = 1
		}
		for j := range len(sol.numbers) {
			n := numbers.MustAtoi[int](sol.numbers[j][i])
			if op == "+" {
				s += n
			} else {
				s *= n
			}
		}
		res += s
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for i, op := range sol.operators {
		var s int
		if op == "*" {
			s = 1
		}
		sz := len(sol.numbers[0][i])
		nums := make([]string, sz)
		for j := range sz {
			for k := range len(sol.numbers) {
				nums[j] += string([]byte{sol.numbers[k][i][j]})
			}
		}
		for _, ns := range nums {
			n := numbers.MustAtoi[int](ns)
			if op == "+" {
				s += n
			} else {
				s *= n
			}
		}
		res += s
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
