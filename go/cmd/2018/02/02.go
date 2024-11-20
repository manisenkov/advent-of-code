package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 2
type Solution struct {
	ids []string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.ids = make([]string, len(input))
	copy(sol.ids, input)
}

// Part1 .
func (sol *Solution) Part1() any {
	twoLet := 0
	threeLet := 0
	for _, id := range sol.ids {
		if contains2Letters(id) {
			twoLet++
		}
		if contains3Letters(id) {
			threeLet++
		}
	}
	return twoLet * threeLet
}

// Part2 .
func (sol *Solution) Part2() any {
	res := ""
	for i := 0; i < len(sol.ids)-1; i++ {
		for j := i + 1; j < len(sol.ids); j++ {
			diff := calcDiff(sol.ids[i], sol.ids[j])
			if len(diff) > len(res) {
				res = diff
			}
		}
	}
	return res
}

func calcDiff(id1, id2 string) string {
	rid1 := []rune(id1)
	rid2 := []rune(id2)
	res := ""
	for i, r1 := range rid1 {
		r2 := rid2[i]
		if r1 == r2 {
			res += string(r1)
		}
	}
	return res
}

func contains2Letters(s string) bool {
	buckets := map[rune]int{}
	for _, c := range s {
		buckets[c]++
	}
	for _, n := range buckets {
		if n == 2 {
			return true
		}
	}
	return false
}

func contains3Letters(s string) bool {
	buckets := map[rune]int{}
	for _, c := range s {
		buckets[c]++
	}
	for _, n := range buckets {
		if n == 3 {
			return true
		}
	}
	return false
}

func main() {
	problem.Solve(new(Solution))
}
