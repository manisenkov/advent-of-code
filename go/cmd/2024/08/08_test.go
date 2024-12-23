package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

type Day08TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day08TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day08TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(14, st.sol.Part1())
	st.Equal(34, st.sol.Part2())
}

func TestDay08(t *testing.T) {
	st := new(Day08TestSuite)
	suite.Run(t, st)
}
