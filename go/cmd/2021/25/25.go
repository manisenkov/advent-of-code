package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 25
type Solution struct {
	initMap [][]byte
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.initMap = make([][]byte, len(input))
	for i, s := range input {
		sol.initMap[i] = []byte(s)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	cur := sol.initMap
	numSteps := 1
	for {
		nextMap := sol.emptyMap()

		// First: east-movers
		for i := 0; i < len(cur); i++ {
			for j := 0; j < len(cur[i]); j++ {
				jRight := (j + 1) % len(cur[i])
				if cur[i][j] == '>' {
					if cur[i][jRight] == '.' {
						nextMap[i][jRight] = '>'
					} else {
						nextMap[i][j] = '>'
					}
				}
			}
		}

		// Then: south-movers
		for i := 0; i < len(nextMap); i++ {
			for j := 0; j < len(nextMap[i]); j++ {
				if cur[i][j] == 'v' {
					iBottom := (i + 1) % len(nextMap)
					if nextMap[iBottom][j] == '.' && cur[iBottom][j] != 'v' {
						nextMap[iBottom][j] = 'v'
					} else {
						nextMap[i][j] = 'v'
					}
				}
			}
		}

		if isMapsEqual(cur, nextMap) {
			break
		}
		cur = nextMap
		numSteps++
	}
	return numSteps
}

// Part2 .
func (sol *Solution) Part2() any {
	return 0
}

func (sol *Solution) emptyMap() [][]byte {
	res := make([][]byte, len(sol.initMap))
	for i, s := range sol.initMap {
		res[i] = make([]byte, len(s))
		for j := range s {
			res[i][j] = '.'
		}
	}
	return res
}

func isMapsEqual(map1, map2 [][]byte) bool {
	if len(map1) != len(map2) {
		return false
	}
	for i, s := range map1 {
		for j, b := range s {
			if b != map2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	problem.Solve(new(Solution))
}
