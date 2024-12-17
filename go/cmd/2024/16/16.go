package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (d direction) move(pos [2]int) [2]int {
	switch d {
	case north:
		return [2]int{pos[0] - 1, pos[1]}
	case east:
		return [2]int{pos[0], pos[1] + 1}
	case south:
		return [2]int{pos[0] + 1, pos[1]}
	case west:
		return [2]int{pos[0], pos[1] - 1}
	}
	panic("wrong direction")
}

func (d direction) rotLeft() direction {
	switch d {
	case north:
		return west
	case east:
		return north
	case south:
		return east
	case west:
		return south
	}
	panic("wrong direction")
}

func (d direction) rotRight() direction {
	switch d {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	panic("wrong direction")
}

type step struct {
	pos [2]int
	dir direction
}

func (s step) options(plan [][]rune) map[step]int {
	res := map[step]int{}
	d := s.dir
	t := d.move(s.pos)
	if plan[t[0]][t[1]] == '.' {
		res[step{pos: t, dir: d}] = 1
	}
	d = s.dir.rotLeft()
	t = d.move(s.pos)
	if plan[t[0]][t[1]] == '.' {
		res[step{pos: t, dir: d}] = 1001
	}
	d = s.dir.rotRight()
	t = d.move(s.pos)
	if plan[t[0]][t[1]] == '.' {
		res[step{pos: t, dir: d}] = 1001
	}
	return res
}

type record struct {
	points int
	paths  [][][2]int
}

// Solution contains a solution for day 16
type Solution struct {
	startPos [2]int
	endPos   [2]int
	size     [2]int
	plan     [][]rune
	minPts   int
	minPaths [][][2]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.plan = make([][]rune, len(input))
	sol.size = [2]int{len(input), len(input[0])}
	for i, line := range input {
		sol.plan[i] = make([]rune, sol.size[0])
		for j, c := range line {
			switch c {
			case '#':
				sol.plan[i][j] = '#'
			default:
				sol.plan[i][j] = '.'
			}
			if c == 'S' {
				sol.startPos = [2]int{i, j}
			} else if c == 'E' {
				sol.endPos = [2]int{i, j}
			}
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	pos := sol.startPos
	dir := east
	points := map[step]record{
		{pos, dir}: {0, [][][2]int{{sol.startPos}}},
	}
	queue := []step{{pos, dir}}
	sol.minPts = 0xFFFFFFFF
	sol.minPaths = [][][2]int{}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		curRec := points[s]
		if s.pos == sol.endPos {
			if curRec.points == sol.minPts {
				sol.minPaths = append(
					sol.minPaths,
					curRec.paths...)
			}
			if curRec.points < sol.minPts {
				sol.minPts = curRec.points
				sol.minPaths = curRec.paths
			}
			continue
		}
		opts := s.options(sol.plan)
		for opt, pts := range opts {
			p, ok := points[opt]
			if ok && curRec.points+pts == p.points {
				points[opt] = record{
					points: curRec.points + pts,
					paths: append(
						p.paths,
						collections.MapTo(curRec.paths, func(p [][2]int) [][2]int {
							return append(append([][2]int{}, p...), opt.pos)
						})...),
				}
			}
			if !ok || curRec.points+pts < p.points {
				points[opt] = record{
					points: curRec.points + pts,
					paths: append(
						[][][2]int{},
						collections.MapTo(curRec.paths, func(p [][2]int) [][2]int {
							return append(append([][2]int{}, p...), opt.pos)
						})...),
				}
				queue = append(queue, opt)
			}
		}
	}
	return sol.minPts
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := collections.Set[[2]int]{}
	for _, path := range sol.minPaths {
		for _, pos := range path {
			res[pos] = true
		}
	}
	return len(res)
}

func main() {
	problem.Solve(new(Solution))
}
