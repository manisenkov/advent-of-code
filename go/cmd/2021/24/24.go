package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 24
type Solution struct {
	args [][3]int
	sols []string
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.args = [][3]int{}
	for i := 0; i < len(input); i += 18 {
		sol.args = append(sol.args, [3]int{
			numbers.MustParseInt[int](strings.Split(input[i+4], " ")[2], 10),
			numbers.MustParseInt[int](strings.Split(input[i+5], " ")[2], 10),
			numbers.MustParseInt[int](strings.Split(input[i+15], " ")[2], 10),
		})
	}
}

func (sol *Solution) Eval(z, w int, step int) int {
	x := 0
	x += z
	x %= 26
	z /= sol.args[step][0]
	x += sol.args[step][1]
	if x == w {
		x = 1
	} else {
		x = 0
	}
	if x == 0 {
		x = 1
	} else {
		x = 0
	}
	y := 0
	y += 25
	y *= x
	y += 1
	z *= y
	y = 0
	y += w
	y += sol.args[step][2]
	y *= x
	z += y
	return z
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	zs := map[int][]string{0: {""}}
	for step := 0; step < 14; step++ {
		nextZs := map[int][]string{}
		var cap int
		if step == 13 {
			cap = 1
		} else {
			cap = 1000000
		}
		for z, nums := range zs {
			for w := 1; w <= 9; w++ {
				nextZ := sol.Eval(z, w, step)
				if nextZ < 0 || nextZ >= cap {
					continue
				}
				for _, n := range nums {
					nextZs[nextZ] = append(nextZs[nextZ], n+fmt.Sprint(w))
				}
			}
		}
		zs = nextZs
	}
	sol.sols = zs[0]
	sort.Strings(sol.sols)
	return sol.sols[len(sol.sols)-1]
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return sol.sols[0]
}

func main() {
	problem.Solve(new(Solution))
}
