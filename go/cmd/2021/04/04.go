package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type board [5][5]int

func (b *board) isWinner() bool {
	// Check rows
	for i := 0; i < 5; i++ {
		if b[i][0] < 0 && b[i][1] < 0 && b[i][2] < 0 && b[i][3] < 0 && b[i][4] < 0 {
			return true
		}
	}
	// Check columns
	for j := 0; j < 5; j++ {
		if b[0][j] < 0 && b[1][j] < 0 && b[2][j] < 0 && b[3][j] < 0 && b[4][j] < 0 {
			return true
		}
	}

	return false
}

func (b *board) mark(num int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == num {
				b[i][j] -= 100
				return true
			}
		}
	}
	return false
}

func (b *board) sumUnmarked() int {
	res := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] > 0 {
				res += b[i][j]
			}
		}
	}
	return res
}

type stat struct {
	draw   int
	winner board
}

// Solution contains solution for day 4
type Solution struct {
	draws  []int
	boards []board
	stats  []stat
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	// Parse draws
	xs := strings.Split(input[0], ",")
	draws := make([]int, len(xs))
	for i, s := range xs {
		draws[i] = numbers.MustAtoi[int](s)
	}
	sol.draws = draws

	// Parse boards
	numBoards := (len(input) - 1) / 6
	boards := make([]board, numBoards)
	for nb := range boards {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				s := strings.Trim(input[2+i+nb*6][j*3:j*3+2], " ")
				boards[nb][i][j] = numbers.MustAtoi[int](s)
			}
		}
	}
	sol.boards = boards
}

func (sol *Solution) playBingo() []stat {
	if len(sol.stats) > 0 {
		return sol.stats
	}
	stats := []stat{}
	boards := make([]board, len(sol.boards))
	copy(boards, sol.boards)
	for _, draw := range sol.draws {
		if len(boards) == 0 {
			break
		}
		losers := []board{}
		for nb := range boards {
			if boards[nb].mark(draw) && boards[nb].isWinner() {
				stats = append(stats, stat{draw: draw, winner: boards[nb]})
			} else {
				losers = append(losers, boards[nb])
			}
		}
		boards = losers
	}
	sol.stats = stats
	return stats
}

// Part1 .
func (sol *Solution) Part1() any {
	stats := sol.playBingo()
	return stats[0].draw * stats[0].winner.sumUnmarked()
}

// Part2 .
func (sol *Solution) Part2() any {
	stats := sol.playBingo()
	return stats[len(stats)-1].draw * stats[len(stats)-1].winner.sumUnmarked()
}

func main() {
	problem.Solve(new(Solution))
}
