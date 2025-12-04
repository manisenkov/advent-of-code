package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 3
type Solution struct {
	banks [][]int
}

func findMax(ns []int) (int, int) {
	idx := 0
	res := ns[0]
	for i := 1; i < len(ns); i++ {
		if ns[i] > res {
			res = ns[i]
			idx = i
		}
	}
	return res, idx
}

func calcJolts(ns []int) int {
	res := 0
	for i := range len(ns) {
		res += ns[i] * numbers.PowInt(10, len(ns)-i-1)
	}
	return res
}

func solve(bank []int, sz int) int {
	res := 0
	nums := []int{}
	idx := 0
	for sz > 0 {
		sz--
		x, i := findMax(bank[idx : len(bank)-sz])
		nums = append(nums, x)
		idx = idx + i + 1
	}
	res += calcJolts(nums)
	return res
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.banks = [][]int{}
	for _, s := range input {
		s = strings.Trim(s, " \n")
		if s == "" {
			continue
		}
		bank := make([]int, len(s))
		for j := 0; j < len(s); j++ {
			bank[j] = numbers.MustAtoi[int](s[j : j+1])
		}
		sol.banks = append(sol.banks, bank)
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, bank := range sol.banks {
		res += solve(bank, 2)
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, bank := range sol.banks {
		res += solve(bank, 12)
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
