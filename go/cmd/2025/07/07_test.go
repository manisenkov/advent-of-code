package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

type Day07TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day07TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day07TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(21, st.sol.Part1())
	st.Equal(40, st.sol.Part2())
}

func TestDay07(t *testing.T) {
	st := new(Day07TestSuite)
	suite.Run(t, st)
}
