package main

import "github.com/manisenkov/advent-of-code/go/pkg/common"

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
