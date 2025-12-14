package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 11
type Solution struct {
	paths       map[string][]string
	invertPaths map[string][]string
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.paths = make(map[string][]string)
	sol.invertPaths = make(map[string][]string)
	for _, s := range input {
		parts := strings.Split(s, ": ")
		source := parts[0]
		targets := strings.Split(parts[1], " ")
		sol.paths[source] = targets
		for _, target := range targets {
			sol.invertPaths[target] = append(sol.invertPaths[target], source)
		}
	}
}

func (sol *Solution) calcTotalNumberOfPaths(target string, distances map[string]int, exclude collections.Set[string]) int {
	if d, ok := distances[target]; ok {
		return d
	}
	res := 0
	for _, source := range sol.invertPaths[target] {
		if !exclude[source] {
			res += sol.calcTotalNumberOfPaths(source, distances, exclude)
		}
	}
	distances[target] = res
	return res
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	return sol.calcTotalNumberOfPaths("out", map[string]int{"you": 1}, collections.Set[string]{})
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	r1 := sol.calcTotalNumberOfPaths("fft", map[string]int{"svr": 1}, collections.SetFromSlice([]string{"dac"}))
	r2 := sol.calcTotalNumberOfPaths("dac", map[string]int{"fft": 1}, collections.SetFromSlice([]string{"svr"}))
	r3 := sol.calcTotalNumberOfPaths("out", map[string]int{"dac": 1}, collections.SetFromSlice([]string{"fft", "svr"}))
	r4 := sol.calcTotalNumberOfPaths("dac", map[string]int{"svr": 1}, collections.SetFromSlice([]string{"fft"}))
	r5 := sol.calcTotalNumberOfPaths("fft", map[string]int{"dac": 1}, collections.SetFromSlice([]string{"svr"}))
	r6 := sol.calcTotalNumberOfPaths("out", map[string]int{"fft": 1}, collections.SetFromSlice([]string{"dac", "svr"}))
	return (r1 * r2 * r3) + (r4 * r5 * r6)
}

func main() {
	problem.Solve(new(Solution))
}
