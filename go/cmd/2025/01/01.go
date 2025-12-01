package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type direction string

const (
	left  direction = "L"
	right direction = "R"
)

type rotation struct {
	dir   direction
	count int
}

func (r rotation) move(pos int, withPasses bool) (int, int) {
	passCounter := 0
	count := r.count
	if withPasses {
		passCounter = count / 100
		count %= 100
		if (r.dir == left && pos != 0 && pos < count) || (r.dir == right && withPasses && pos+count > 100) {
			passCounter++
		}
	}
	if r.dir == left {
		pos -= count
	} else {
		pos += count
	}
	// Normalize position
	if pos < 0 {
		pos = (100 + pos) % 100
	} else {
		pos %= 100
	}
	if pos == 0 {
		passCounter++
	}
	return pos, passCounter
}

func parseRotation(input string) rotation {
	var dir direction
	var count int
	if input[0] == 'L' {
		dir = left
	} else {
		dir = right
	}
	count = numbers.MustAtoi[int](input[1:])
	return rotation{dir, count}
}

// Solution contains a solution for day 1
type Solution struct {
	instructions []rotation
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.instructions = make([]rotation, len(input))
	for i, s := range input {
		sol.instructions[i] = parseRotation(s)
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	pos := 50
	res := 0
	var counter int
	for _, s := range sol.instructions {
		pos, counter = s.move(pos, false)
		res += counter
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	pos := 50
	res := 0
	var counter int
	for _, s := range sol.instructions {
		pos, counter = s.move(pos, true)
		res += counter
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
