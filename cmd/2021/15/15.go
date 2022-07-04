package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 15
type Solution struct {
	field [][]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.field = make([][]int, len(input))
	for i, s := range input {
		sol.field[i] = make([]int, len(s))
		for j, t := range []byte(s) {
			sol.field[i][j] = int(t - '0')
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	return findPath(sol.field)
}

// Part2 .
func (sol *Solution) Part2() any {
	height := len(sol.field)
	width := len(sol.field[0])
	bigField := make([][]int, height*5)
	for i := 0; i < height*5; i++ {
		bigField[i] = make([]int, width*5)
		for j := 0; j < width*5; j++ {
			bigField[i][j] = sol.field[i%height][j%width] + i/height + j/height
			if bigField[i][j] > 9 {
				bigField[i][j] -= 9
			}
		}
	}
	return findPath(bigField)
}

func findPath(field [][]int) int {
	height := len(field)
	width := len(field[0])
	dist := make([][]int, height)
	for i := 0; i < height; i++ {
		dist[i] = make([]int, width)
		for j := 0; j < width; j++ {
			dist[i][j] = 0x7fffffff
		}
	}
	dist[0][0] = 0
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		row := queue[0][0]
		col := queue[0][1]
		queue = queue[1:]
		cur := dist[row][col]
		adj := [][2]int{
			{row - 1, col},
			{row + 1, col},
			{row, col - 1},
			{row, col + 1},
		}
		for _, adjP := range adj {
			adjRow := adjP[0]
			adjCol := adjP[1]
			if adjRow < 0 || adjCol < 0 || adjRow >= height || adjCol >= width {
				continue
			}
			if cur+field[adjRow][adjCol] < dist[adjRow][adjCol] {
				dist[adjRow][adjCol] = cur + field[adjRow][adjCol]
				queue = append(queue, adjP)
			}
		}
	}
	return dist[height-1][width-1]
}

func main() {
	common.Run(new(Solution))
}
