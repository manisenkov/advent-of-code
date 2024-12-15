package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
	"github.com/manisenkov/advent-of-code/pkg/rmat"
)

type game struct {
	pX, pY, xA, xB, yA, yB int
}

// Solution contains a solution for day 13
type Solution struct {
	games []game
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	for i := 0; i < len(input); i += 4 {
		parts := strings.Split(input[i], ": ")
		parts = strings.Split(parts[1], ", ")
		xParts := strings.Split(parts[0], "+")
		yParts := strings.Split(parts[1], "+")
		xA := numbers.MustAtoi[int](xParts[1])
		yA := numbers.MustAtoi[int](yParts[1])

		parts = strings.Split(input[i+1], ": ")
		parts = strings.Split(parts[1], ", ")
		xParts = strings.Split(parts[0], "+")
		yParts = strings.Split(parts[1], "+")
		xB := numbers.MustAtoi[int](xParts[1])
		yB := numbers.MustAtoi[int](yParts[1])

		parts = strings.Split(input[i+2], ": ")
		parts = strings.Split(parts[1], ", ")
		xParts = strings.Split(parts[0], "=")
		yParts = strings.Split(parts[1], "=")
		pX := numbers.MustAtoi[int](xParts[1])
		pY := numbers.MustAtoi[int](yParts[1])

		sol.games = append(sol.games, game{pX, pY, xA, xB, yA, yB})
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, g := range sol.games {
		a := rmat.FromIntTable([][]int{
			{g.xA, g.xB},
			{g.yA, g.yB},
		})
		p := rmat.ColFromIntSlice([]int{g.pX, g.pY})
		x := a.Inverse().Mul(p)
		if !x.At(0, 0).IsInt() || !x.At(1, 0).IsInt() {
			continue
		}
		res += 3*int(x.AtInt64(0, 0)) + int(x.AtInt64(1, 0))
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, g := range sol.games {
		a := rmat.FromIntTable([][]int{
			{g.xA, g.xB},
			{g.yA, g.yB},
		})
		p := rmat.ColFromIntSlice([]int{10000000000000 + g.pX, 10000000000000 + g.pY})
		x := a.Inverse().Mul(p)
		if !x.At(0, 0).IsInt() || !x.At(1, 0).IsInt() {
			continue
		}
		res += 3*int(x.AtInt64(0, 0)) + int(x.AtInt64(1, 0))
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
