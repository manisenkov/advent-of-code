package registry

import "github.com/manisenkov/advent-of-code/pkg/common"

// DaySolution contains day solutions (both part 1 and part 2)
type DaySolution interface {
	Init([]string) error
	Part1() common.Any
	Part2() common.Any
}

var solutionRegistry = map[int]map[int]DaySolution{}

// Get returns a solution for a given day
func Get(year, day int) (DaySolution, bool) {
	yearReg, ok := solutionRegistry[year]
	if !ok {
		return nil, false
	}
	solutions, ok := yearReg[day]
	return solutions, ok
}

// Register solution for a day
func Register(year, day int, solution DaySolution) {
	yearReg, ok := solutionRegistry[year]
	if !ok {
		yearReg = map[int]DaySolution{}
		solutionRegistry[year] = yearReg
	}
	yearReg[day] = solution
}
