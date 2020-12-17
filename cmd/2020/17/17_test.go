package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
.#.
..#
###
`

type Day17TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day17TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day17TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(112, st.sol.Part1())
	st.Equal(848, st.sol.Part2())
}

func TestDay17(t *testing.T) {
	st := new(Day17TestSuite)
	suite.Run(t, st)
}
