package main

import (
	"sort"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 9
type Solution struct {
	field [][]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.field = make([][]int, len(input))
	for i, s := range input {
		sol.field[i] = make([]int, len(s))
		for j, t := range s {
			sol.field[i][j] = int(t) - int('0')
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := 0
	for i, row := range sol.field {
		for j, n := range row {
			countHigherAdj := 0
			if i == 0 || (i > 0 && n < sol.field[i-1][j]) {
				countHigherAdj++
			}
			if i == len(sol.field)-1 || (i < len(sol.field)-1 && n < sol.field[i+1][j]) {
				countHigherAdj++
			}
			if j == 0 || (j > 0 && n < row[j-1]) {
				countHigherAdj++
			}
			if j == len(row)-1 || (j < len(row)-1 && n < row[j+1]) {
				countHigherAdj++
			}
			if countHigherAdj == 4 {
				res += n + 1
			}
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	bassinMap := make([][]int, len(sol.field))
	for i, row := range sol.field {
		bassinMap[i] = make([]int, len(row))
	}
	bassinSizes := map[int]int{}
	nextBassin := 1
	for i, row := range sol.field {
		for j, n := range row {
			if bassinMap[i][j] != 0 || n == 9 {
				continue
			}
			queue := [][2]int{{i, j}}
			for len(queue) > 0 {
				r := queue[0][0]
				c := queue[0][1]
				queue = queue[1:]
				if bassinMap[r][c] != 0 {
					continue
				}
				bassinMap[r][c] = nextBassin
				bassinSizes[nextBassin]++
				if r > 0 && sol.field[r-1][c] != 9 && bassinMap[r-1][c] == 0 {
					queue = append(queue, [2]int{r - 1, c})
				}
				if c > 0 && sol.field[r][c-1] != 9 && bassinMap[r][c-1] == 0 {
					queue = append(queue, [2]int{r, c - 1})
				}
				if r < len(sol.field)-1 && sol.field[r+1][c] != 9 && bassinMap[r+1][c] == 0 {
					queue = append(queue, [2]int{r + 1, c})
				}
				if c < len(sol.field[r])-1 && sol.field[r][c+1] != 9 && bassinMap[r][c+1] == 0 {
					queue = append(queue, [2]int{r, c + 1})
				}
			}
			nextBassin++
		}
	}
	sizes := []int{}
	for _, sz := range bassinSizes {
		sizes = append(sizes, sz)
	}
	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {
	common.Run(new(Solution))
}
