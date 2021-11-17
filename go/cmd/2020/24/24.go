package main

import (
	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

// tile represents a tile coord in hexagonal grid
// nice doc here -> https://www.redblobgames.com/grids/hexagons/#coordinates
type tile struct {
	x, y, z int
}

var directions = []string{"e", "se", "sw", "w", "nw", "ne"}

func (t *tile) move(dir string) tile {
	switch dir {
	case "e":
		return tile{t.x + 1, t.y - 1, t.z}
	case "se":
		return tile{t.x, t.y - 1, t.z + 1}
	case "sw":
		return tile{t.x - 1, t.y, t.z + 1}
	case "w":
		return tile{t.x - 1, t.y + 1, t.z}
	case "nw":
		return tile{t.x, t.y + 1, t.z - 1}
	case "ne":
		return tile{t.x + 1, t.y, t.z - 1}
	}
	panic("Wrong direction")
}

// Solution contains solution for day 24
type Solution struct {
	paths     [][]string
	startGrid map[tile]bool
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.paths = make([][]string, len(input))
	for i, inp := range input {
		sol.paths[i] = make([]string, 0)
		j := 0
		sInp := []rune(inp)
		for j < len(inp) {
			if sInp[j] == 'w' || sInp[j] == 'e' {
				sol.paths[i] = append(sol.paths[i], string(sInp[j]))
				j++
			} else {
				sol.paths[i] = append(sol.paths[i], string(sInp[j])+string(sInp[j+1]))
				j += 2
			}
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	sol.startGrid = make(map[tile]bool)
	for _, path := range sol.paths {
		t := tile{0, 0, 0}
		for _, dir := range path {
			t = t.move(dir)
		}
		if _, ok := sol.startGrid[t]; ok {
			delete(sol.startGrid, t)
		} else {
			sol.startGrid[t] = true
		}
	}
	return len(sol.startGrid)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	curGrid := sol.startGrid
	for i := 0; i < 100; i++ {
		nextGrid := make(map[tile]bool)
		checked := make(map[tile]bool)
		for t := range curGrid {
			toCheck := make([]tile, 1, 7)
			toCheck[0] = t
			for _, dir := range directions {
				q := t.move(dir)
				toCheck = append(toCheck, q)
			}
			for _, q := range toCheck {
				if checked[q] {
					continue
				}
				if isNextBlack(q, curGrid) {
					nextGrid[q] = true
				}
				checked[q] = true
			}
		}
		curGrid = nextGrid
	}
	return len(curGrid)
}

func isNextBlack(t tile, grid map[tile]bool) bool {
	curBlack := grid[t]
	neighbors := 0
	for _, dir := range directions {
		if grid[t.move(dir)] {
			neighbors++
		}
	}
	if curBlack {
		return neighbors > 0 && neighbors < 3
	}
	return neighbors == 2
}

func main() {
	common.Run(new(Solution))
}
