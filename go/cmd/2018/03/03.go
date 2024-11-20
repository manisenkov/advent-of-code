package main

import (
	"fmt"
	"regexp"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

var claimRegex = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

// Solution contains solution for day 3
type Solution struct {
	ids    []string
	fabric map[string][]string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.fabric = make(map[string][]string)
	sol.ids = make([]string, len(input))
	for i, inp := range input {
		m := claimRegex.FindAllStringSubmatch(inp, -1)
		id := m[0][1]
		left := numbers.MustAtoi[int](m[0][2])
		top := numbers.MustAtoi[int](m[0][3])
		width := numbers.MustAtoi[int](m[0][4])
		height := numbers.MustAtoi[int](m[0][5])
		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				k := fmt.Sprintf("%v,%v", x, y)
				sol.fabric[k] = append(sol.fabric[k], id)
			}
		}
		sol.ids[i] = id
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for _, ids := range sol.fabric {
		if len(ids) > 1 {
			res++
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	idsLeft := map[string]bool{}
	for _, id := range sol.ids {
		idsLeft[id] = true
	}
	for _, ids := range sol.fabric {
		if len(ids) > 1 {
			for _, id := range ids {
				delete(idsLeft, id)
			}
		}
	}
	for id := range idsLeft {
		return id
	}
	return ""
}

func main() {
	problem.Solve(new(Solution))
}
