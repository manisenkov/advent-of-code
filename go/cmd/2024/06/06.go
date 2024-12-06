package main

import (
	"maps"
	"slices"

	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type direction int

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}
	return ""
}

const (
	up direction = iota
	down
	left
	right
)

type move struct {
	pos [2]int
	dir direction
}

func simulate(size [2]int, obstacles map[[2]int]bool, startPos [2]int) ([]move, int, bool) {
	cur := startPos
	dir := up
	visited := make(map[[2]int][]direction)
	steps := make([]move, 0)
	var looped = false
	for cur[0] >= 0 && cur[0] < size[0] && cur[1] > 0 && cur[1] < size[1] {
		if slices.Index(visited[cur], dir) >= 0 {
			looped = true
			break
		}
		visited[cur] = append(visited[cur], dir)
		steps = append(steps, move{cur, dir})
		var nextPos [2]int
		switch dir {
		case up:
			nextPos = [2]int{cur[0] - 1, cur[1]}
		case down:
			nextPos = [2]int{cur[0] + 1, cur[1]}
		case left:
			nextPos = [2]int{cur[0], cur[1] - 1}
		case right:
			nextPos = [2]int{cur[0], cur[1] + 1}
		}
		if obstacles[nextPos] {
			switch dir {
			case up:
				dir = right
			case down:
				dir = left
			case left:
				dir = up
			case right:
				dir = down
			}
		} else {
			cur = nextPos
		}
	}
	return steps, len(visited), looped
}

// Solution contains a solution for day 6
type Solution struct {
	size      [2]int
	obstacles map[[2]int]bool
	startPos  [2]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.size = [2]int{len(input), len(input[0])}
	sol.obstacles = make(map[[2]int]bool)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			switch input[i][j] {
			case '#':
				sol.obstacles[[2]int{i, j}] = true
			case '^':
				sol.startPos = [2]int{i, j}
			}
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	_, res, _ := simulate(sol.size, sol.obstacles, sol.startPos)
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	moves, _, _ := simulate(sol.size, sol.obstacles, sol.startPos)
	placesToPutObstacles := map[[2]int]bool{}
	for _, step := range moves {
		var toPlaceObstacle [2]int
		switch step.dir {
		case up:
			toPlaceObstacle = [2]int{step.pos[0] - 1, step.pos[1]}
		case down:
			toPlaceObstacle = [2]int{step.pos[0] + 1, step.pos[1]}
		case left:
			toPlaceObstacle = [2]int{step.pos[0], step.pos[1] - 1}
		case right:
			toPlaceObstacle = [2]int{step.pos[0], step.pos[1] + 1}
		}
		if (toPlaceObstacle[0] < 0 && toPlaceObstacle[0] >= sol.size[0] &&
			toPlaceObstacle[1] < 0 && toPlaceObstacle[1] >= sol.size[1]) ||
			sol.obstacles[toPlaceObstacle] ||
			toPlaceObstacle == sol.startPos {
			continue
		}
		newObstacles := maps.Clone(sol.obstacles)
		newObstacles[toPlaceObstacle] = true
		_, _, looped := simulate(sol.size, newObstacles, sol.startPos)
		if looped {
			placesToPutObstacles[toPlaceObstacle] = true
		}
	}
	return len(placesToPutObstacles)
}

func main() {
	problem.Solve(new(Solution))
}
