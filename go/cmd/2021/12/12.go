package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 12
type Solution struct {
	edges map[string][]string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.edges = map[string][]string{}
	for _, s := range input {
		xs := strings.Split(s, "-")
		sol.edges[xs[0]] = append(sol.edges[xs[0]], xs[1])
		sol.edges[xs[1]] = append(sol.edges[xs[1]], xs[0])
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	paths := sol.travel("start", map[string]int{}, false)
	return len(paths)
}

// Part2 .
func (sol *Solution) Part2() any {
	paths := sol.travel("start", map[string]int{}, true)
	return len(paths)
}

func (sol *Solution) travel(cur string, visitedSmall map[string]int, doubleVisitAllowed bool) [][]string {
	if cur == "end" {
		return [][]string{{"end"}}
	}

	res := [][]string{}

	if isSmall(cur) {
		visitedSmall[cur]++
		doubleVisitAllowed = doubleVisitAllowed && visitedSmall[cur] != 2
	}

	for _, to := range sol.edges[cur] {
		canVisit := to != "start" && (!isSmall(to) || visitedSmall[to] == 0 || doubleVisitAllowed)
		if canVisit {
			rec := sol.travel(to, visitedSmall, doubleVisitAllowed)
			for _, s := range rec {
				res = append(res, append([]string{cur}, s...))
			}
		}
	}

	if isSmall(cur) {
		visitedSmall[cur]--
	}

	return res
}

func isSmall(s string) bool {
	return strings.ToUpper(s) != s
}

func main() {
	problem.Solve(new(Solution))
}
