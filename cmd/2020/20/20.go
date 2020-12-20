package main

import (
	"math"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 20
type Solution struct {
	tiles     map[int][]tile
	monster   monsterType
	grid      [][]rune
	roughness int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.tiles = make(map[int][]tile, 0)
	for i := 0; i < len(input); i += 12 {
		t := parseTile(input[i : i+11])
		sol.tiles[t.id] = []tile{
			t,
			t.turnCW(),
			t.turnCW().turnCW(),
			t.turnCW().turnCW().turnCW(),
			t.flip(),
			t.flip().turnCW(),
			t.flip().turnCW().turnCW(),
			t.flip().turnCW().turnCW().turnCW(),
		}
	}
	sol.monster = parseMonster()
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	side := int(math.Sqrt(float64(len(sol.tiles))))
	tileGrid := make([][]tile, side)
	for i := 0; i < side; i++ {
		tileGrid[i] = make([]tile, side)
	}
	tileGrid, ok := sol.solve(side, 0, tileGrid, map[int]bool{})
	if !ok {
		panic("Can't solve :(")
	}

	// Parse grid for part 2
	grid := make([][]rune, side*8)
	for i := 0; i < side*8; i++ {
		grid[i] = make([]rune, side*8)
	}
	for i, row := range tileGrid {
		for j, t := range row {
			for y := 1; y < 9; y++ {
				for x := 1; x < 9; x++ {
					c := t.body[y][x]
					if c == '#' {
						sol.roughness++
					}
					grid[i*8+y-1][j*8+x-1] = c
				}
			}
		}
	}
	sol.grid = grid

	return tileGrid[0][0].id * tileGrid[0][side-1].id * tileGrid[side-1][0].id * tileGrid[side-1][side-1].id
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	variants := [][][]rune{
		sol.grid,
		turnGrid(sol.grid),
		turnGrid(turnGrid(sol.grid)),
		turnGrid(turnGrid(turnGrid(sol.grid))),
		flipGrid(sol.grid),
		turnGrid(flipGrid(sol.grid)),
		turnGrid(turnGrid(flipGrid(sol.grid))),
		turnGrid(turnGrid(turnGrid(flipGrid(sol.grid)))),
	}

	numMonsters := 0
	for _, v := range variants {
		n := sol.monster.numMonsters(v)
		if n > numMonsters {
			numMonsters = n
		}
	}
	return sol.roughness - numMonsters*sol.monster.size()
}

func (sol *Solution) solve(side int, pos int, grid [][]tile, used map[int]bool) ([][]tile, bool) {
	if pos == side*side {
		return grid, true
	}
	i := pos / side
	j := pos % side
	for tileID, variants := range sol.tiles {
		if used[tileID] {
			continue
		}
		for _, tile := range variants {
			// Try to match with top tile
			if ((i > 0 && grid[i-1][j].borders.bottom == tile.borders.top) || i == 0) &&
				((j > 0 && grid[i][j-1].borders.right == tile.borders.left) || j == 0) {
				grid[i][j] = tile
				used[tileID] = true
				if _, ok := sol.solve(side, pos+1, grid, used); ok {
					return grid, ok
				}
				grid[i][j] = emptyTile
				used[tileID] = false
			}
		}
	}
	return grid, false
}

func main() {
	common.Run(new(Solution))
}
