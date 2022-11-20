package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 5
type Solution struct {
	polymer []rune
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.polymer = []rune(input[0])
}

// Part1 .
func (sol *Solution) Part1() any {
	return len(collapse(sol.polymer))
}

// Part2 .
func (sol *Solution) Part2() any {
	shortest := len(sol.polymer)
	units := "abcdefghijklmnopqrstuvwxyz"
	for _, u := range units {
		pol := []rune(strings.ReplaceAll(strings.ReplaceAll(string(sol.polymer), string(u), ""), strings.ToUpper(string(u)), ""))
		collapsed := collapse(pol)
		if len(collapsed) < shortest {
			shortest = len(collapsed)
		}
	}
	return shortest
}

func collapse(pol []rune) []rune {
	i := 0
	for i < len(pol)-1 {
		c1 := pol[i]
		c2 := pol[i+1]
		if c1 != c2 && strings.ToLower(string(c1)) == strings.ToLower(string(c2)) {
			nPol := make([]rune, len(pol)-2)
			copy(nPol, pol[0:i])
			copy(nPol[i:], pol[i+2:])
			pol = nPol
			i = i - 1
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return pol
}

func main() {
	common.Run(new(Solution))
}
