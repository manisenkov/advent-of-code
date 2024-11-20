package main

import (
	"strconv"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

var simpleTurns = map[rune]map[int]rune{
	'E': {0: 'E', 90: 'N', 180: 'W', 270: 'S'},
	'N': {0: 'N', 90: 'W', 180: 'S', 270: 'E'},
	'W': {0: 'W', 90: 'S', 180: 'E', 270: 'N'},
	'S': {0: 'S', 90: 'E', 180: 'N', 270: 'W'},
}

type instruction struct {
	command rune
	arg     int
}

func (ins instruction) String() string {
	return string(ins.command) + strconv.Itoa(ins.arg)
}

// Solution contains solution for day 12
type Solution struct {
	instructions []instruction
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.instructions = make([]instruction, len(input))
	for i, inp := range input {
		sol.instructions[i] = instruction{
			command: []rune(inp)[0],
			arg:     numbers.MustAtoi[int](inp[1:]),
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	dir := 'E'
	pos := [2]int{0, 0}
	for _, ins := range sol.instructions {
		pos, dir = moveSimple(pos, dir, ins)
	}
	return numbers.Abs(pos[0]) + numbers.Abs(pos[1])
}

// Part2 .
func (sol *Solution) Part2() any {
	waypoint := [2]int{10, 1}
	pos := [2]int{0, 0}
	for _, ins := range sol.instructions {
		pos, waypoint = moveWaypoint(pos, waypoint, ins)
	}
	return numbers.Abs(pos[0]) + numbers.Abs(pos[1])
}

func moveSimple(pos [2]int, dir rune, ins instruction) ([2]int, rune) {
	switch ins.command {
	case 'L':
		return pos, simpleTurns[dir][ins.arg%360]
	case 'R':
		return pos, simpleTurns[dir][360-(ins.arg%360)]
	case 'F':
		return moveSimple(pos, dir, instruction{command: dir, arg: ins.arg})
	case 'E':
		return [2]int{pos[0] + ins.arg, pos[1]}, dir
	case 'N':
		return [2]int{pos[0], pos[1] + ins.arg}, dir
	case 'W':
		return [2]int{pos[0] - ins.arg, pos[1]}, dir
	case 'S':
		return [2]int{pos[0], pos[1] - ins.arg}, dir
	}
	return pos, dir
}

func moveWaypoint(pos [2]int, waypoint [2]int, ins instruction) ([2]int, [2]int) {
	switch ins.command {
	case 'F':
		return [2]int{pos[0] + ins.arg*waypoint[0], pos[1] + ins.arg*waypoint[1]}, waypoint
	case 'L':
		switch ins.arg % 360 {
		case 0:
			return pos, waypoint
		case 90:
			return pos, [2]int{-waypoint[1], waypoint[0]}
		case 180:
			return pos, [2]int{-waypoint[0], -waypoint[1]}
		case 270:
			return pos, [2]int{waypoint[1], -waypoint[0]}
		}
	case 'R':
		switch ins.arg % 360 {
		case 0:
			return pos, waypoint
		case 90:
			return pos, [2]int{waypoint[1], -waypoint[0]}
		case 180:
			return pos, [2]int{-waypoint[0], -waypoint[1]}
		case 270:
			return pos, [2]int{-waypoint[1], waypoint[0]}
		}
	case 'E':
		return pos, [2]int{waypoint[0] + ins.arg, waypoint[1]}
	case 'N':
		return pos, [2]int{waypoint[0], waypoint[1] + ins.arg}
	case 'W':
		return pos, [2]int{waypoint[0] - ins.arg, waypoint[1]}
	case 'S':
		return pos, [2]int{waypoint[0], waypoint[1] - ins.arg}
	}
	return pos, waypoint
}

func main() {
	problem.Solve(new(Solution))
}
