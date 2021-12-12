package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
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
	st.Equal("CABDFE", st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay07(t *testing.T) {
	st := new(Day07TestSuite)
	suite.Run(t, st)
}
