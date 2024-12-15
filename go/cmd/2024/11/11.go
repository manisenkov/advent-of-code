package main

import (
	"strconv"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

var solutions map[[2]int]int = map[[2]int]int{}

func blink(stone int, count int) int {
	if count == 0 {
		return 1
	}

	if res, ok := solutions[[2]int{stone, count}]; ok {
		return res
	}

	nextStones := make([]int, 0)
	if stone == 0 {
		nextStones = append(nextStones, 1)
	} else if s := strconv.Itoa(stone); len(s)%2 == 0 {
		nextStones = append(nextStones, numbers.MustAtoi[int](s[:len(s)/2]), numbers.MustAtoi[int](s[len(s)/2:]))
	} else {
		nextStones = append(nextStones, stone*2024)
	}
	res := numbers.Sum(collections.MapTo(nextStones, func(s int) int {
		return blink(s, count-1)
	}))
	solutions[[2]int{stone, count}] = res
	return res
}

// Solution contains a solution for day 11
type Solution struct {
	stones []int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.stones = collections.MapTo(strings.Split(input[0], " "), numbers.MustAtoi[int])
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	return numbers.Sum(collections.MapTo(sol.stones, func(s int) int {
		return blink(s, 25)
	}))
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return numbers.Sum(collections.MapTo(sol.stones, func(s int) int {
		return blink(s, 75)
	}))
}

func main() {
	problem.Solve(new(Solution))
}
