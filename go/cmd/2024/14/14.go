package main

import (
	"fmt"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type robot struct {
	pos [2]int
	vel [2]int
}

func (r *robot) move(steps int, size [2]int) [2]int {
	x := ((r.pos[0] + steps*r.vel[0]) + (size[0] * steps)) % size[0]
	y := ((r.pos[1] + steps*r.vel[1]) + (size[1] * steps)) % size[1]
	return [2]int{x, y}
}

func checkImage(robotPos [][2]int) bool {
	posSet := collections.SetFromSlice(robotPos)
	if len(posSet) != len(robotPos) {
		return false
	}
	return true
}

func print(robotPos [][2]int, size [2]int) {
	lines := make([]string, size[1])
	for i := 0; i < size[1]; i++ {
		for j := 0; j < size[0]; j++ {
			lines[i] += " "
		}
	}
	for _, r := range robotPos {
		runes := []rune(lines[r[1]])
		runes[r[0]] = '#'
		lines[r[1]] = string(runes)
	}
	for _, s := range lines {
		fmt.Println(s)
	}
}

// Solution contains a solution for day 14
type Solution struct {
	size   [2]int
	robots []robot
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.size = [2]int{101, 103}
	sol.robots = make([]robot, len(input))
	for i, line := range input {
		parts := strings.Split(line, " ")
		posParts := strings.Split(strings.Split(parts[0], "=")[1], ",")
		velParts := strings.Split(strings.Split(parts[1], "=")[1], ",")
		sol.robots[i] = robot{
			pos: [2]int{numbers.MustAtoi[int](posParts[0]), numbers.MustAtoi[int](posParts[1])},
			vel: [2]int{numbers.MustAtoi[int](velParts[0]), numbers.MustAtoi[int](velParts[1])},
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	robotPos := make([][2]int, len(sol.robots))
	for i, robot := range sol.robots {
		robotPos[i] = robot.move(100, sol.size)
	}
	var quadrants [4]int
	for _, p := range robotPos {
		if p[0] < sol.size[0]/2 && p[1] < sol.size[1]/2 {
			quadrants[0]++
		} else if p[0] > sol.size[0]/2 && p[1] < sol.size[1]/2 {
			quadrants[1]++
		} else if p[0] < sol.size[0]/2 && p[1] > sol.size[1]/2 {
			quadrants[2]++
		} else if p[0] > sol.size[0]/2 && p[1] > sol.size[1]/2 {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	robotPos := make([][2]int, len(sol.robots))
	for step := 0; step < 10000; step++ {
		for i, robot := range sol.robots {
			robotPos[i] = robot.move(step, sol.size)
		}
		posSet := collections.SetFromSlice(robotPos)
		if len(posSet) == len(robotPos) {
			print(robotPos, sol.size)
			return step
		}
	}
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
