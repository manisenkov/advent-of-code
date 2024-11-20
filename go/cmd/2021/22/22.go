package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 22
type Solution struct {
	cuboids [][3][2]int
	states  []bool
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.cuboids = make([][3][2]int, len(input))
	sol.states = make([]bool, len(input))
	for i, s := range input {
		xs := strings.Split(s, " ")
		if xs[0] == "on" {
			sol.states[i] = true
		}
		xs = strings.Split(xs[1], ",")
		for j := 0; j < 3; j++ {
			cs := strings.Split(xs[j][2:], "..")
			sol.cuboids[i][j][0] = numbers.MustAtoi[int](cs[0])
			sol.cuboids[i][j][1] = numbers.MustAtoi[int](cs[1])
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	reactor := map[[3]int]bool{}
	for i, cuboid := range sol.cuboids {
		st := sol.states[i]
		for x := numbers.Max(cuboid[0][0], -50); x <= numbers.Min(cuboid[0][1], 50); x++ {
			for y := numbers.Max(cuboid[1][0], -50); y <= numbers.Min(cuboid[1][1], 50); y++ {
				for z := numbers.Max(cuboid[2][0], -50); z <= numbers.Min(cuboid[2][1], 50); z++ {
					reactor[[3]int{x, y, z}] = st
				}
			}
		}
	}
	res := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if reactor[[3]int{x, y, z}] {
					res++
				}
			}
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	// Find borders
	borderMap := [3]map[int]bool{{}, {}, {}}
	for _, cuboid := range sol.cuboids {
		for i, p := range cuboid {
			borderMap[i][p[0]] = true
			borderMap[i][p[1]+1] = true
		}
	}
	borders := [3][]int{{}, {}, {}}
	for i, m := range borderMap {
		for p := range m {
			borders[i] = append(borders[i], p)
		}
		sort.Ints(borders[i])
	}

	// Calculate size of blocks that are turned on
	res := 0
	for x := 0; x < len(borders[0])-1; x++ {
		fmt.Printf("%v of %v\n", x+1, len(borders[0])-1)
		for y := 0; y < len(borders[1])-1; y++ {
			for z := 0; z < len(borders[2])-1; z++ {
				pos := [3]int{borders[0][x], borders[1][y], borders[2][z]}
				for i := len(sol.cuboids) - 1; i >= 0; i-- {
					inCuboid := sol.cuboids[i][0][0] <= pos[0] && pos[0] <= sol.cuboids[i][0][1] &&
						sol.cuboids[i][1][0] <= pos[1] && pos[1] <= sol.cuboids[i][1][1] &&
						sol.cuboids[i][2][0] <= pos[2] && pos[2] <= sol.cuboids[i][2][1]
					if inCuboid {
						if sol.states[i] {
							res += (borders[0][x+1] - borders[0][x]) *
								(borders[1][y+1] - borders[1][y]) *
								(borders[2][z+1] - borders[2][z])
						}
						break
					}
				}
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
