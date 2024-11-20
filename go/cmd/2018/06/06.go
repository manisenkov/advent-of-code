package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type coord struct {
	row, col int
}

func (c coord) dist(other coord) int {
	return numbers.Abs(c.row-other.row) + numbers.Abs(c.col-other.col)
}

type distance struct {
	equal   bool
	nearest coord
	value   int
}

// Solution contains solution for day 6
type Solution struct {
	maxDist       int
	width, height int
	points        []coord
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	points := make([]coord, len(input))
	boundRect := [2]coord{
		{row: 0x7fffffff, col: 0x7fffffff},
		{row: -0x7fffffff, col: -0x7fffffff},
	}
	for i, c := range input {
		xs := strings.Split(c, ", ")
		p := coord{col: numbers.MustAtoi[int](xs[0]), row: numbers.MustAtoi[int](xs[1])}
		if p.row < boundRect[0].row {
			boundRect[0].row = p.row
		}
		if p.col < boundRect[0].col {
			boundRect[0].col = p.col
		}
		if p.row > boundRect[1].row {
			boundRect[1].row = p.row
		}
		if p.col > boundRect[1].col {
			boundRect[1].col = p.col
		}
		points[i] = p
	}

	// Extend the bound rectange to 1 cell to each direction
	boundRect[0].row--
	boundRect[0].col--
	boundRect[1].row++
	boundRect[1].col++

	// Normalize coordinate system by shifting it to (0,0)
	for i := range points {
		points[i].row -= boundRect[0].row
		points[i].col -= boundRect[0].col
	}

	sol.points = points
	sol.width = boundRect[1].col - boundRect[0].col + 1
	sol.height = boundRect[1].row - boundRect[0].row + 1
	sol.maxDist = 10000
}

// Part1 .
func (sol *Solution) Part1() any {
	// Calculate distances
	grid := make([][]distance, sol.height)
	for row := range grid {
		grid[row] = make([]distance, sol.width)
		for col := range grid[row] {
			gridPoint := coord{row, col}
			minDist := distance{
				equal:   false,
				nearest: sol.points[0],
				value:   sol.points[0].dist(gridPoint),
			}
			for _, p := range sol.points[1:] {
				distValue := p.dist(gridPoint)
				if distValue < minDist.value {
					minDist = distance{
						equal:   false,
						nearest: p,
						value:   distValue,
					}
				} else if distValue == minDist.value {
					minDist.equal = true
				}
			}
			grid[row][col] = minDist
		}
	}

	// Exclude points on edge
	pointsToExclude := map[coord]bool{}
	for col := range grid[0] {
		pointsToExclude[grid[0][col].nearest] = true
		pointsToExclude[grid[sol.height-1][col].nearest] = true
	}
	for row := range grid {
		pointsToExclude[grid[row][0].nearest] = true
		pointsToExclude[grid[row][sol.width-1].nearest] = true
	}

	// Calculate biggest area
	areas := map[coord]int{}
	for row := range grid {
		for col := range grid[row] {
			dist := grid[row][col]
			if dist.equal || pointsToExclude[dist.nearest] {
				continue
			}
			areas[dist.nearest]++
		}
	}
	maxArea := 0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// Part2 .
func (sol *Solution) Part2() any {
	// Calculate distances
	grid := make([][][]distance, sol.height)
	for row := range grid {
		grid[row] = make([][]distance, sol.width)
		for col := range grid[row] {
			gridPoint := coord{row, col}
			dists := make([]distance, len(sol.points))
			for i, point := range sol.points {
				dists[i] = distance{
					nearest: point,
					value:   point.dist(gridPoint),
				}
			}
			grid[row][col] = dists
		}
	}

	// Calculate area that less than maxDist
	res := 0
	for row := range grid {
		for col := range grid[row] {
			sumDist := 0
			for _, dist := range grid[row][col] {
				sumDist += dist.value
			}
			if sumDist < sol.maxDist {
				res++
			}
		}
	}

	return res
}

func main() {
	problem.Solve(new(Solution))
}
