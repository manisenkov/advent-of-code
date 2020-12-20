package main

import (
	"math"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

const tileSize = 10

type tileBorders struct {
	left   int
	top    int
	right  int
	bottom int
}

type tile struct {
	id      int
	body    [tileSize][tileSize]rune
	borders tileBorders
}

var emptyTile = tile{}

func (t tile) turnCW() tile {
	res := tile{
		id:      t.id,
		body:    [tileSize][tileSize]rune{},
		borders: t.borders,
	}
	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			res.body[i][j] = t.body[tileSize-j-1][i]
		}
	}
	res.borders = tileBorders{
		left:   t.borders.bottom,
		top:    flipInt10(t.borders.left),
		right:  t.borders.top,
		bottom: flipInt10(t.borders.right),
	}
	return res
}

func (t tile) flip() tile {
	return tile{
		id: t.id,
		body: [tileSize][tileSize]rune{
			t.body[9],
			t.body[8],
			t.body[7],
			t.body[6],
			t.body[5],
			t.body[4],
			t.body[3],
			t.body[2],
			t.body[1],
			t.body[0],
		},
		borders: tileBorders{
			left:   flipInt10(t.borders.left),
			top:    t.borders.bottom,
			right:  flipInt10(t.borders.right),
			bottom: t.borders.top,
		},
	}
}

func parseTile(input []string) tile {
	// Parse ID
	body := [tileSize][tileSize]rune{}
	for i := 0; i < tileSize; i++ {
		sInput := []rune(input[i+1])
		body[i] = [tileSize]rune{sInput[0], sInput[1], sInput[2], sInput[3], sInput[4], sInput[5], sInput[6], sInput[7], sInput[8], sInput[9]}
	}
	return tile{
		id:   common.MustAtoi(input[0][5:9]),
		body: body,
		borders: tileBorders{
			left:   calcBorder([tileSize]rune{body[0][0], body[1][0], body[2][0], body[3][0], body[4][0], body[5][0], body[6][0], body[7][0], body[8][0], body[9][0]}),
			top:    calcBorder(body[0]),
			right:  calcBorder([tileSize]rune{body[0][9], body[1][9], body[2][9], body[3][9], body[4][9], body[5][9], body[6][9], body[7][9], body[8][9], body[9][9]}),
			bottom: calcBorder(body[9]),
		},
	}
}

// Solution contains solution for day 20
type Solution struct {
	tiles map[int][]tile
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
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	side := int(math.Sqrt(float64(len(sol.tiles))))
	grid := make([][]tile, side)
	for i := 0; i < side; i++ {
		grid[i] = make([]tile, side)
	}
	grid, ok := sol.solve(side, 0, grid, map[int]bool{})
	if !ok {
		panic("Can't solve :(")
	}
	return grid[0][0].id * grid[0][side-1].id * grid[side-1][0].id * grid[side-1][side-1].id
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	return 0
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

func flipInt10(n int) int {
	return ((n & 1) << 9) +
		((n & 2) << 7) +
		((n & 4) << 5) +
		((n & 8) << 3) +
		((n & 16) << 1) +
		((n & 32) >> 1) +
		((n & 64) >> 3) +
		((n & 128) >> 5) +
		((n & 256) >> 7) +
		((n & 512) >> 9)
}

func calcBorder(row [tileSize]rune) int {
	b := 1
	res := 0
	for i, r := range row {
		if r == '#' {
			res += (b << i)
		}
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
