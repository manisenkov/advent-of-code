package main

import (
	"math"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 17
type Solution struct {
	minX, maxX, minY, maxY int
	velocities             map[[2]int]bool
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	s := strings.Split(strings.Split(input[0], ": ")[1], ", ")
	xs := strings.Split(string([]byte(s[0])[2:]), "..")
	sol.minX = numbers.MustAtoi[int](xs[0])
	sol.maxX = numbers.MustAtoi[int](xs[1])
	ys := strings.Split(string([]byte(s[1])[2:]), "..")
	sol.minY = numbers.MustAtoi[int](ys[0])
	sol.maxY = numbers.MustAtoi[int](ys[1])
}

// Part1 .
func (sol *Solution) Part1() any {
	minVX := int(-1.0 + math.Sqrt(float64(1+8*sol.minX))/2.0)
	if minVX*(minVX+1)/2 < sol.minX {
		minVX++
	}
	maxVX := sol.maxX

	maxVY := -0x7FFFFFFF
	sol.velocities = map[[2]int]bool{}
	for vx := minVX; vx <= maxVX; vx++ {
		for vy := sol.minY; vy < 1000; vy++ {
			if sol.shot(vx, vy) {
				if vy > maxVY {
					maxVY = vy
				}
				sol.velocities[[2]int{vx, vy}] = true
			}
		}
	}
	return maxVY * (maxVY + 1) / 2
}

// Part2 .
func (sol *Solution) Part2() any {
	return len(sol.velocities)
}

func (sol *Solution) shot(vx, vy int) bool {
	x := 0
	y := 0
	for y > sol.minY && (x < sol.minX || y > sol.maxY) {
		x += vx
		y += vy
		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		vy--
	}
	return x >= sol.minX && x <= sol.maxX && y >= sol.minY && y <= sol.maxY
}

func main() {
	problem.Solve(new(Solution))
}
