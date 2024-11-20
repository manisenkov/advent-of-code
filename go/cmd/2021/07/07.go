package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 7
type Solution struct {
	initPos []int
	maxPos  int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	xs := strings.Split(input[0], ",")
	sol.initPos = make([]int, len(xs))
	sol.maxPos = -0x7FFFFFFF
	for i, s := range xs {
		p := numbers.MustAtoi[int](s)
		sol.initPos[i] = p
		if p > sol.maxPos {
			sol.maxPos = p
		}
	}
}

func (sol *Solution) solve(distFn func(curPos, initPos int) int) int {
	minSum := 0x7FFFFFFF
	for i := 0; i <= sol.maxPos; i++ {
		sum := 0
		for _, p := range sol.initPos {
			sum += distFn(i, p)
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return minSum
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.solve(func(curPos, initPos int) int {
		return numbers.Abs(curPos - initPos)
	})
}

// Part2 .
func (sol *Solution) Part2() any {
	return sol.solve(func(curPos, initPos int) int {
		n := numbers.Abs(curPos - initPos)
		return n * (n + 1) / 2
	})
}

func main() {
	problem.Solve(new(Solution))
}
