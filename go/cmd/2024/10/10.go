package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func nextSteps(plan [][]int, size [2]int, pos [2]int) [][2]int {
	res := [][2]int{}
	cur := plan[pos[0]][pos[1]]
	if pos[0] > 0 && plan[pos[0]-1][pos[1]] == cur+1 {
		res = append(res, [2]int{pos[0] - 1, pos[1]})
	}
	if pos[0] < size[0]-1 && plan[pos[0]+1][pos[1]] == cur+1 {
		res = append(res, [2]int{pos[0] + 1, pos[1]})
	}
	if pos[1] > 0 && plan[pos[0]][pos[1]-1] == cur+1 {
		res = append(res, [2]int{pos[0], pos[1] - 1})
	}
	if pos[1] < size[1]-1 && plan[pos[0]][pos[1]+1] == cur+1 {
		res = append(res, [2]int{pos[0], pos[1] + 1})
	}
	return res
}

// Solution contains a solution for day 10
type Solution struct {
	plan       [][]int
	size       [2]int
	trailheads collections.Set[[2]int]
	reachable  map[[2]int]map[[2]int]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.size = [2]int{len(input), len(input[0])}
	sol.trailheads = make(collections.Set[[2]int])
	sol.plan = collections.Map(input, func(s string) []int {
		return collections.Map([]rune(s), func(t rune) int {
			return int(t - '0')
		})
	})
	for i, row := range sol.plan {
		for j, x := range row {
			if x == 0 {
				sol.trailheads[[2]int{i, j}] = true
			}
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	sol.reachable = make(map[[2]int]map[[2]int]int)
	for trailhead := range sol.trailheads {
		queue := [][2]int{trailhead}
		sol.reachable[trailhead] = make(map[[2]int]int)
		for len(queue) > 0 {
			pos := queue[0]
			queue = queue[1:]
			if sol.plan[pos[0]][pos[1]] == 9 {
				sol.reachable[trailhead][pos] += 1
			} else {
				next := nextSteps(sol.plan, sol.size, pos)
				queue = append(queue, next...)
			}
		}
		res += len(sol.reachable[trailhead])
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for trailhead := range sol.trailheads {
		for _, rank := range sol.reachable[trailhead] {
			res += rank
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
