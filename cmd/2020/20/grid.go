package main

import "fmt"

func flipGrid(grid [][]rune) [][]rune {
	h := len(grid)
	res := make([][]rune, h)
	for i := 0; i < h; i++ {
		res[i] = make([]rune, h)
		copy(res[i], grid[h-i-1])
	}
	return res
}

func turnGrid(grid [][]rune) [][]rune {
	w := len(grid[0])
	h := len(grid)
	res := make([][]rune, w)
	for i := 0; i < w; i++ {
		res[i] = make([]rune, h)
	}
	for i, row := range grid {
		for j, c := range row {
			res[j][h-i-1] = c
		}
	}
	return res
}

func printGrid(grid [][]rune) {
	s := ""
	for _, r := range grid {
		s += string(r) + "\n"
	}
	fmt.Println(s)
}
