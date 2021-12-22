package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
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
			sol.cuboids[i][j][0] = common.MustAtoi(cs[0])
			sol.cuboids[i][j][1] = common.MustAtoi(cs[1])
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	reactor := map[[3]int]bool{}
	for i, cuboid := range sol.cuboids {
		st := sol.states[i]
		for x := common.Max(cuboid[0][0], -50); x <= common.Min(cuboid[0][1], 50); x++ {
			for y := common.Max(cuboid[1][0], -50); y <= common.Min(cuboid[1][1], 50); y++ {
				for z := common.Max(cuboid[2][0], -50); z <= common.Min(cuboid[2][1], 50); z++ {
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
func (sol *Solution) Part2() common.Any {
	return 0
}

func main() {
	common.Run(new(Solution))
}
