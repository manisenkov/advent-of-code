package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

type Day20TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day20TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day20TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.sol.tasks = [2]task{
		{1, 2},
		{50, 20},
	}
	st.Equal(44, st.sol.Part1())
	st.Equal(285, st.sol.Part2())
}

func TestDay20(t *testing.T) {
	st := new(Day20TestSuite)
	suite.Run(t, st)
}
