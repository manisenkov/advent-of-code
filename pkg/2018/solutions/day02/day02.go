package day02

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 2
type Solution struct {
	ids []string
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) error {
	s.ids = make([]string, len(input))
	copy(s.ids, input)
	return nil
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

// Part1 .
func (s *Solution) Part1() common.Any {
	twoLet := 0
	threeLet := 0
	for _, id := range s.ids {
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
func (s *Solution) Part2() common.Any {
	res := ""
	for i := 0; i < len(s.ids)-1; i++ {
		for j := i + 1; j < len(s.ids); j++ {
			diff := calcDiff(s.ids[i], s.ids[j])
			if len(diff) > len(res) {
				res = diff
			}
		}
	}
	return res
}
