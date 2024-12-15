package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func isSafe(r []int) bool {
	isInc := r[1] > r[0]
	for i := 1; i < len(r); i++ {
		d := numbers.Abs(r[i] - r[i-1])
		if d < 1 || d > 3 || (isInc && r[i] < r[i-1]) || (!isInc && r[i] > r[i-1]) {
			return false
		}
	}
	return true
}

// Solution contains a solution for day 2
type Solution struct {
	reports [][]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.reports = collections.MapTo(input, func(s string) []int {
		return collections.MapTo(strings.Split(s, " "), numbers.MustAtoi[int])
	})
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, r := range sol.reports {
		if isSafe(r) {
			res += 1
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, r := range sol.reports {
		if isSafe(r) {
			res += 1
			continue
		}
		for i := range r {
			q := make([]int, len(r)-1)
			copy(q, r[:i])
			copy(q[i:], r[i+1:])
			if isSafe(q) {
				res += 1
				break
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
