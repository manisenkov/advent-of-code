package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

type Day16TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day16TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day16TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(11048, st.sol.Part1())
	st.Equal(64, st.sol.Part2())
}

func TestDay16(t *testing.T) {
	st := new(Day16TestSuite)
	suite.Run(t, st)
}
