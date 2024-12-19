package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func match(towelIndex map[rune][][]rune, design []rune, cache map[string]int) int {
	if res, ok := cache[string(design)]; ok {
		return res
	}
	if len(design) == 0 {
		return 1
	}
	options := towelIndex[design[0]]
	if len(options) == 0 {
		return 0
	}
	res := 0
	for _, opt := range options {
		if len(design) >= len(opt) && collections.IsEqualSlices(opt, design[:len(opt)]) {
			numMatches := match(towelIndex, design[len(opt):], cache)
			res += numMatches
		}
	}
	cache[string(design)] = res
	return res
}

// Solution contains a solution for day 19
type Solution struct {
	towelIndex map[rune][][]rune
	designs    [][]rune
	cache      map[string]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	towels := collections.MapTo(strings.Split(input[0], ", "), func(t string) []rune { return []rune(t) })
	sol.towelIndex = make(map[rune][][]rune)
	for _, towel := range towels {
		sol.towelIndex[towel[0]] = append(sol.towelIndex[towel[0]], towel)
	}
	sol.designs = collections.MapTo(input[2:], func(t string) []rune { return []rune(t) })
	sol.cache = make(map[string]int)
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	return len(collections.Filter(sol.designs, func(d []rune) bool {
		return match(sol.towelIndex, d, sol.cache) > 0
	}))
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return collections.Reduce(collections.MapTo(sol.designs, func(d []rune) int {
		return match(sol.towelIndex, d, sol.cache)
	}), func(r1, r2 int) int { return r1 + r2 })
}

func main() {
	problem.Solve(new(Solution))
}
