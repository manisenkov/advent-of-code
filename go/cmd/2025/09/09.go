package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func checkPair(covered collections.Set[[2]int], p, q [2]int) bool {
	lt := [2]int{numbers.Min(p[0], q[0]), numbers.Min(p[1], q[1])}
	rt := [2]int{numbers.Min(p[0], q[0]), numbers.Max(p[1], q[1])}
	lb := [2]int{numbers.Max(p[0], q[0]), numbers.Min(p[1], q[1])}
	rb := [2]int{numbers.Max(p[0], q[0]), numbers.Max(p[1], q[1])}

	pairs := [4][2][2]int{
		{lt, rt},
		{rt, rb},
		{lb, rb},
		{lt, lb},
	}
	for _, pair := range pairs {
		x := pair[0]
		y := pair[1]
		if x == y {
			continue
		}
		dir := [2]int{y[0] - x[0], y[1] - x[1]}
		dir = [2]int{dir[0] / numbers.Abs(dir[0]+dir[1]), dir[1] / numbers.Abs(dir[0]+dir[1])}
		cur := x
		for cur != y {
			cur = [2]int{cur[0] + dir[0], cur[1] + dir[1]}
			if !covered[cur] {
				return false
			}
		}
	}

	return true
}

// Solution contains a solution for day 9
type Solution struct {
	tiles [][2]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.tiles = make([][2]int, len(input))
	for i, s := range input {
		parts := strings.Split(s, ",")
		sol.tiles[i] = [2]int{numbers.MustAtoi[int](parts[0]), numbers.MustAtoi[int](parts[1])}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	maxArea := 0
	for i := range len(sol.tiles) - 1 {
		for j := i + 1; j < len(sol.tiles); j++ {
			p := sol.tiles[i]
			q := sol.tiles[j]
			area := (numbers.Abs(p[0]-q[0]) + 1) * (numbers.Abs(p[1]-q[1]) + 1)
			maxArea = numbers.Max(area, maxArea)
		}
	}
	return maxArea
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	// Find borders of the polygon
	covered := make(collections.Set[[2]int])
	borders := make(collections.Set[[2]int])
	rowTiles := make(collections.Set[int])
	colTiles := make(collections.Set[int])
	for i := range len(sol.tiles) {
		j := (i + 1) % len(sol.tiles)
		p := sol.tiles[i]
		q := sol.tiles[j]
		dir := [2]int{q[0] - p[0], q[1] - p[1]}
		dir = [2]int{dir[0] / numbers.Abs(dir[0]+dir[1]), dir[1] / numbers.Abs(dir[0]+dir[1])}
		cur := p
		rowTiles[p[0]] = true
		colTiles[p[1]] = true
		covered[cur] = true
		borders[cur] = true
		for cur != q {
			cur = [2]int{cur[0] + dir[0], cur[1] + dir[1]}
			covered[cur] = true
			borders[cur] = true
		}
	}

	// Find first inner cell
	leftRightRow := make(map[int][]int)
	for c := range covered {
		leftRightRow[c[0]] = append(leftRightRow[c[0]], c[1])
		slices.Sort(leftRightRow[c[0]])
	}
	var cur [2]int
	for row, cols := range leftRightRow {
		if len(cols) == 2 && cols[1]-cols[0] > 1 {
			cur = [2]int{row, cols[0] + 1}
			break
		}
	}

	// Find inner coverage of the polygon
	queue := [][2]int{cur}
	steps := 0
	for len(queue) > 0 {
		steps++
		if steps%100000 == 0 {
			fmt.Printf(" -- len(queue) = %v, len(covered) = %v\n", len(queue), len(covered))
		}
		cur = queue[0]
		queue = queue[1:]

		if !(rowTiles[cur[0]] || colTiles[cur[1]] ||
			borders[[2]int{cur[0] - 1, cur[1] - 1}] ||
			borders[[2]int{cur[0] - 1, cur[1]}] ||
			borders[[2]int{cur[0] - 1, cur[1] + 1}] ||
			borders[[2]int{cur[0], cur[1] - 1}] ||
			borders[[2]int{cur[0], cur[1] + 1}] ||
			borders[[2]int{cur[0] + 1, cur[1] - 1}] ||
			borders[[2]int{cur[0] + 1, cur[1]}] ||
			borders[[2]int{cur[0] + 1, cur[1] + 1}]) {
			continue
		}

		neighbors := [][2]int{
			{cur[0] - 1, cur[1]},
			{cur[0], cur[1] - 1},
			{cur[0] + 1, cur[1]},
			{cur[0], cur[1] + 1},
		}
		for _, n := range neighbors {
			if !covered[n] {
				covered[n] = true
				queue = append(queue, n)
			}
		}
	}

	// Checking areas
	maxArea := 0
	for i := range len(sol.tiles) - 1 {
		fmt.Printf(" == tile %v of %v\n", i+1, len(sol.tiles))
		for j := i + 1; j < len(sol.tiles); j++ {
			p := sol.tiles[i]
			q := sol.tiles[j]
			if checkPair(covered, p, q) {
				area := (numbers.Abs(p[0]-q[0]) + 1) * (numbers.Abs(p[1]-q[1]) + 1)
				maxArea = numbers.Max(area, maxArea)
			}
		}
	}
	return maxArea
}

func main() {
	problem.Solve(new(Solution))
}
