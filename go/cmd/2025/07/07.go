package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 7
type Solution struct {
	width     int
	height    int
	startPos  int
	splitters collections.Set[[2]int]
	cache     map[[2]int]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.width = len(input[0])
	sol.height = len(input)
	sol.splitters = make(collections.Set[[2]int])
	for row, s := range input {
		for col, c := range s {
			switch c {
			case 'S':
				sol.startPos = col
			case '^':
				sol.splitters[[2]int{row, col}] = true
			}
		}
	}
	sol.cache = make(map[[2]int]int)
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	splitCount := 0
	startBeam := [2]int{0, sol.startPos}
	beamed := collections.Set[[2]int]{startBeam: true}
	beams := [][2]int{startBeam}
	for len(beams) > 0 {
		curBeam := beams[0]
		beams = beams[1:]
		for {
			curBeam[0]++
			if curBeam[0] >= sol.height || beamed[curBeam] {
				break
			}
			beamed[curBeam] = true
			if sol.splitters[curBeam] {
				beamLeft := [2]int{curBeam[0], curBeam[1] - 1}
				beamRight := [2]int{curBeam[0], curBeam[1] + 1}
				if !beamed[beamLeft] {
					beams = append(beams, beamLeft)
					beamed[beamLeft] = true
				}
				if !beamed[beamRight] {
					beams = append(beams, beamRight)
					beamed[beamRight] = true
				}
				splitCount++
				break
			}
		}
	}
	return splitCount
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return sol.solve([2]int{0, sol.startPos})
}

func (sol *Solution) solve(beam [2]int) int {
	if r := sol.cache[beam]; r > 0 {
		return r
	}
	if beam[0] >= sol.height {
		return 1
	}
	if sol.splitters[beam] {
		beamLeft := [2]int{beam[0], beam[1] - 1}
		beamRight := [2]int{beam[0], beam[1] + 1}
		r := sol.solve(beamLeft) + sol.solve(beamRight)
		sol.cache[beam] = r
		return r
	}
	r := sol.solve([2]int{beam[0] + 1, beam[1]})
	sol.cache[beam] = r
	return r
}

func main() {
	problem.Solve(new(Solution))
}
