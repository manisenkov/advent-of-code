package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type group struct {
	size   int
	counts map[rune]int
}

// Solution contains solution for day 6
type Solution struct {
	groups []group
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.groups = make([]group, 0)
	currentGroup := group{size: 0, counts: map[rune]int{}}
	newGroup := true
	for _, inp := range input {
		if inp == "" {
			sol.groups = append(sol.groups, currentGroup)
			currentGroup = group{size: 0, counts: map[rune]int{}}
			newGroup = true
			continue
		}
		newGroup = false
		for _, c := range inp {
			currentGroup.counts[c]++
		}
		currentGroup.size++
	}
	if !newGroup {
		sol.groups = append(sol.groups, currentGroup)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for _, g := range sol.groups {
		res += len(g.counts)
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	res := 0
	for _, g := range sol.groups {
		for _, c := range g.counts {
			if c == g.size {
				res++
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
