package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func getPlots(plan [][]rune) ([]collections.Set[[2]int], map[[2]int]int, []rune) {
	plots := []collections.Set[[2]int]{}
	names := []rune{}
	visited := make(map[[2]int]int)
	cur := [][2]int{}
	next := [][2]int{{0, 0}}
	for len(cur) > 0 || len(next) > 0 {
		if len(cur) == 0 {
			for len(next) > 0 {
				nextPos := next[0]
				next = next[1:]
				if _, ok := visited[nextPos]; !ok {
					cur = append(cur, nextPos)
					plots = append(plots, make(collections.Set[[2]int]))
					names = append(names, plan[nextPos[0]][nextPos[1]])
					break
				}
			}
			continue
		}
		pos := cur[0]
		cur = cur[1:]
		if _, ok := visited[pos]; ok {
			continue
		}
		val := plan[pos[0]][pos[1]]
		plotIdx := len(plots) - 1
		visited[pos] = plotIdx
		plots[plotIdx][pos] = true
		toCheck := [][2]int{
			{pos[0] - 1, pos[1]},
			{pos[0], pos[1] - 1},
			{pos[0] + 1, pos[1]},
			{pos[0], pos[1] + 1},
		}
		for _, p := range toCheck {
			if p[0] >= 0 && p[0] < len(plan) && p[1] >= 0 && p[1] < len(plan[0]) {
				if plan[p[0]][p[1]] == val {
					cur = append(cur, p)
				} else {
					next = append(next, p)
				}
			}
		}
	}
	return plots, visited, names
}

// Solution contains a solution for day 12
type Solution struct {
	plan [][]rune
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.plan = collections.Map(input, func(s string) []rune {
		return []rune(s)
	})
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	plots, _, _ := getPlots(sol.plan)
	res := 0
	for _, plot := range plots {
		perimeter := 0
		for pos := range plot {
			toCheck := [][2]int{
				{pos[0] - 1, pos[1]},
				{pos[0], pos[1] - 1},
				{pos[0] + 1, pos[1]},
				{pos[0], pos[1] + 1},
			}
			for _, x := range toCheck {
				if !plot[x] {
					perimeter += 1
				}
			}
		}
		res += perimeter * len(plot)
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	plots, _, _ := getPlots(sol.plan)
	res := 0
	for _, plot := range plots {
		perimeter := 0
		checkCorner := func(toCheck [][2]int) bool {
			bits := 0
			for i, x := range toCheck {
				if !plot[x] {
					bits += 1 << i
				}
			}
			if bits == 0b111 || bits == 0b001 || bits == 0b110 {
				return true
			}
			return false
		}
		for pos := range plot {
			// Left up corner
			if checkCorner([][2]int{
				{pos[0] - 1, pos[1] - 1},
				{pos[0], pos[1] - 1},
				{pos[0] - 1, pos[1]},
			}) {
				perimeter++
			}

			// Right up corner
			if checkCorner([][2]int{
				{pos[0] + 1, pos[1] - 1},
				{pos[0], pos[1] - 1},
				{pos[0] + 1, pos[1]},
			}) {
				perimeter++
			}

			// Left down corner
			if checkCorner([][2]int{
				{pos[0] - 1, pos[1] + 1},
				{pos[0], pos[1] + 1},
				{pos[0] - 1, pos[1]},
			}) {
				perimeter++
			}

			// Right down corner
			if checkCorner([][2]int{
				{pos[0] + 1, pos[1] + 1},
				{pos[0], pos[1] + 1},
				{pos[0] + 1, pos[1]},
			}) {
				perimeter++
			}
		}
		res += perimeter * len(plot)
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
