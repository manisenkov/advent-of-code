package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type task struct {
	minDiff, cheatSize int
}

// Solution contains a solution for day 20
type Solution struct {
	start [2]int
	end   [2]int
	walls collections.Set[[2]int]
	size  [2]int
	tasks [2]task
	cost  map[[2]int]int
}

func (sol *Solution) run() map[[2]int]int {
	cost := map[[2]int]int{sol.start: 0}
	queue := [][2]int{sol.start}
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
			if sol.walls[opt] {
				continue
			}
			if c, ok := cost[opt]; !ok || cost[pos]+1 < c {
				cost[opt] = cost[pos] + 1
				queue = append(queue, opt)
			}
		}
	}
	return cost
}

func (sol *Solution) solveTask(input task) int {
	res := 0
	for pos, curCost := range sol.cost {
		for i := -input.cheatSize; i <= input.cheatSize; i++ {
			for j := numbers.Abs(i) - input.cheatSize; j <= input.cheatSize-numbers.Abs(i); j++ {
				dist := numbers.Abs(i) + numbers.Abs(j)
				if dist <= 1 {
					continue
				}
				if otherCost, ok := sol.cost[[2]int{pos[0] + i, pos[1] + j}]; ok {
					cheat := otherCost - curCost - dist
					if cheat >= input.minDiff {
						res++
					}
				}
			}
		}
	}
	return res
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.tasks = [2]task{
		{minDiff: 100, cheatSize: 2},
		{minDiff: 100, cheatSize: 20},
	}
	sol.walls = make(collections.Set[[2]int])
	sol.size = [2]int{len(input), len(input[0])}
	for i, line := range input {
		for j, c := range line {
			pos := [2]int{i, j}
			switch c {
			case '#':
				sol.walls[pos] = true
			case 'S':
				sol.start = pos
			case 'E':
				sol.end = pos
			}
		}
	}
	sol.cost = sol.run()
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	return sol.solveTask(sol.tasks[0])
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return sol.solveTask(sol.tasks[1])
}

func main() {
	problem.Solve(new(Solution))
}
