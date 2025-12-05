package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type idRange struct {
	low  int
	high int
}

func (r idRange) isWithin(id int) bool {
	return id >= r.low && id <= r.high
}

func (r idRange) overlap(t idRange) bool {
	return (r.high >= t.low && r.low <= t.high) || (t.high >= r.low && t.low <= r.high)
}

func (r idRange) merge(t idRange) idRange {
	return idRange{numbers.Min(r.low, t.low), numbers.Max(r.high, t.high)}
}

func (r idRange) size() int {
	return r.high - r.low + 1
}

// Solution contains a solution for day 5
type Solution struct {
	ranges []idRange
	ids    []int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	idx := 0
	sol.ranges = []idRange{}
	for ; strings.Trim(input[idx], " \n") != ""; idx++ {
		parts := strings.Split(input[idx], "-")
		sol.ranges = append(sol.ranges, idRange{numbers.MustAtoi[int](parts[0]), numbers.MustAtoi[int](parts[1])})
	}
	idx++
	sol.ids = []int{}
	for ; idx < len(input); idx++ {
		sol.ids = append(sol.ids, numbers.MustAtoi[int](input[idx]))
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, id := range sol.ids {
		for _, r := range sol.ranges {
			if r.isWithin(id) {
				res++
				break
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	cur := make([]idRange, len(sol.ranges))
	copy(cur, sol.ranges)
	for {
		next := []idRange{}
		for _, r := range cur {
			replaced := false
			for i := 0; i < len(next); i++ {
				if next[i].overlap(r) {
					next[i] = r.merge(next[i])
					replaced = true
					break
				}
			}
			if !replaced {
				next = append(next, r)
			}
		}
		if len(next) == len(cur) {
			break
		}
		cur = next
	}
	res := 0
	for _, r := range cur {
		res += r.size()
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
