package main

import (
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 1
type Solution struct {
	left  []int
	right []int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, s := range input {
		xs := strings.Split(s, "   ")
		left[i] = numbers.MustAtoi[int](xs[0])
		right[i] = numbers.MustAtoi[int](xs[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	sol.left = left
	sol.right = right
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	dist := 0
	for i := range sol.left {
		dist += numbers.Abs(sol.left[i] - sol.right[i])
	}
	return dist
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	counts := make(map[int]int)
	for _, item := range sol.right {
		counts[item] += 1
	}
	similarity := 0
	for _, item := range sol.left {
		similarity += item * counts[item]
	}
	return similarity
}

func main() {
	problem.Solve(new(Solution))
}
