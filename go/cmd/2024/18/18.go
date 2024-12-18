package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func simulate(fallenBytes [][2]int, size [2]int) (int, bool) {
	bytes := make(collections.Set[[2]int])
	for _, b := range fallenBytes {
		bytes[b] = true
	}
	start := [2]int{0, 0}
	cost := map[[2]int]int{start: 0}
	queue := [][2]int{start}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		options := [][2]int{
			{pos[0] - 1, pos[1]},
			{pos[0], pos[1] - 1},
			{pos[0] + 1, pos[1]},
			{pos[0], pos[1] + 1},
		}
		for _, opt := range options {
			if opt[0] < 0 || opt[0] >= size[0] || opt[1] < 0 || opt[1] >= size[1] || bytes[opt] {
				continue
			}
			if c, ok := cost[opt]; !ok || cost[pos]+1 < c {
				cost[opt] = cost[pos] + 1
				queue = append(queue, opt)
			}
		}
	}
	res, ok := cost[[2]int{size[0] - 1, size[1] - 1}]
	return res, ok
}

// Solution contains a solution for day 18
type Solution struct {
	fallenBytes [][2]int
	size        [2]int
	numBytes    int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.size = [2]int{71, 71}
	sol.numBytes = 1024
	sol.fallenBytes = make([][2]int, len(input))
	for i, line := range input {
		parts := strings.Split(line, ",")
		sol.fallenBytes[i] = [2]int{
			numbers.MustAtoi[int](parts[0]),
			numbers.MustAtoi[int](parts[1]),
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res, _ := simulate(sol.fallenBytes[:sol.numBytes], sol.size)
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	var res [2]int
	for i := sol.numBytes; i < len(sol.fallenBytes); i++ {
		_, ok := simulate(sol.fallenBytes[:i], sol.size)
		if !ok {
			res = sol.fallenBytes[i-1]
			break
		}
	}
	return strings.Join([]string{numbers.Itoa(res[0]), numbers.Itoa(res[1])}, ",")
}

func main() {
	problem.Solve(new(Solution))
}
