package main

import (
	"slices"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func steps(a, b, size [2]int, maxSteps int) collections.Set[[2]int] {
	res := make(collections.Set[[2]int])
	d := [2]int{b[0] - a[0], b[1] - a[1]}

	// Left direction
	x := a[0] - d[0]
	y := a[1] - d[1]
	n := 0
	for n < maxSteps && x >= 0 && x < size[0] && y >= 0 && y < size[1] {
		res[[2]int{x, y}] = true
		x -= d[0]
		y -= d[1]
		n++
	}

	// Right direction
	x = b[0] + d[0]
	y = b[1] + d[1]
	n = 0
	for n < maxSteps && x >= 0 && x < size[0] && y >= 0 && y < size[1] {
		res[[2]int{x, y}] = true
		x += d[0]
		y += d[1]
		n++
	}

	return res
}

// Solution contains a solution for day 8
type Solution struct {
	antennas map[rune][][2]int
	size     [2]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	antennas := make(map[rune][][2]int)
	size := [2]int{len(input), len(input[0])}
	for i, row := range input {
		for j, c := range row {
			if c != '.' {
				antennas[c] = append(antennas[c], [2]int{i, j})
			}
		}
	}
	// Sort antenna places from left to right
	for _, places := range antennas {
		slices.SortFunc(places, func(a, b [2]int) int {
			return a[0] - b[0]
		})
	}
	sol.antennas = antennas
	sol.size = size
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	antinodes := make(map[[2]int]bool)
	for _, places := range sol.antennas {
		for i := 0; i < len(places)-1; i++ {
			for j := i + 1; j < len(places); j++ {
				a := places[i]
				b := places[j]
				antinodes = collections.MergeSets(steps(a, b, sol.size, 1), antinodes)
			}
		}
	}
	return len(antinodes)
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	antinodes := make(collections.Set[[2]int])
	for _, places := range sol.antennas {
		collections.AppendSliceToSet(places, antinodes)
	}
	for _, places := range sol.antennas {
		for i := 0; i < len(places)-1; i++ {
			for j := i + 1; j < len(places); j++ {
				a := places[i]
				b := places[j]
				antinodes = collections.MergeSets(steps(a, b, sol.size, sol.size[0]), antinodes)
			}
		}
	}
	return len(antinodes)
}

func main() {
	problem.Solve(new(Solution))
}
