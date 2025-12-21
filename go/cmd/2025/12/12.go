package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type region struct {
	width        int
	height       int
	presentCount []int
}

// Solution contains a solution for day 12
type Solution struct {
	presents []int
	regions  []region
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.presents = make([]int, 6)
	for i := range 6 {
		present := 0
		for r := i*5 + 1; r < i*5+4; r++ {
			for _, t := range input[r] {
				if t == '#' {
					present++
				}
			}
		}
		sol.presents[i] = present
	}

	for i := 30; i < len(input); i++ {
		parts := strings.Split(input[i], ": ")
		size := strings.Split(parts[0], "x")
		counts := strings.Split(parts[1], " ")
		sol.regions = append(sol.regions, region{
			width:        numbers.MustAtoi[int](size[0]),
			height:       numbers.MustAtoi[int](size[1]),
			presentCount: collections.MapTo(counts, numbers.MustAtoi[int]),
		})
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	solvableCount := 0
	nonSolvableCount := 0
	potentiallySolvable := []region{}
	for _, reg := range sol.regions {
		nW := reg.width - (reg.width % 3)
		nH := reg.height - (reg.height % 3)
		if nW*nH >= numbers.Sum(reg.presentCount)*9 {
			solvableCount++
		} else {
			sum := 0
			for i, count := range reg.presentCount {
				sum += count * sol.presents[i]
			}
			if sum > reg.width*reg.height {
				nonSolvableCount++
			} else {
				potentiallySolvable = append(potentiallySolvable, reg)
			}
		}
	}
	return solvableCount
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
