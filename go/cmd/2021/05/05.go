package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type point struct {
	x, y int
}

type lineSector struct {
	points [2]point
}

// Solution contains solution for day 5
type Solution struct {
	sectors []lineSector
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sectors := make([]lineSector, len(input))
	for i, s := range input {
		xs := strings.Split(s, " -> ")
		ps1 := strings.Split(xs[0], ",")
		ps2 := strings.Split(xs[1], ",")
		sectors[i] = lineSector{
			points: [2]point{
				{
					x: numbers.MustAtoi[int](ps1[0]),
					y: numbers.MustAtoi[int](ps1[1]),
				},
				{
					x: numbers.MustAtoi[int](ps2[0]),
					y: numbers.MustAtoi[int](ps2[1]),
				},
			},
		}
	}
	sol.sectors = sectors
}

func (sol *Solution) solve(includeDiagonals bool) int {
	markedPoints := map[point]int{}
	for _, sector := range sol.sectors {
		dx := calcStep(sector.points[0].x, sector.points[1].x)
		dy := calcStep(sector.points[0].y, sector.points[1].y)
		if !includeDiagonals && dx != 0 && dy != 0 {
			continue
		}
		for p := sector.points[0]; p != sector.points[1]; {
			markedPoints[p]++
			p.x += dx
			p.y += dy
		}
		markedPoints[sector.points[1]]++
	}

	res := 0
	for _, count := range markedPoints {
		if count > 1 {
			res++
		}
	}
	return res
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.solve(false)
}

// Part2 .
func (sol *Solution) Part2() any {
	return sol.solve(true)
}

func calcStep(p0, p1 int) int {
	if p1 > p0 {
		return 1
	}
	if p1 == p0 {
		return 0
	}
	return -1
}

func main() {
	problem.Solve(new(Solution))
}
