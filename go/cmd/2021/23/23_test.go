package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
`

type Day23TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day23TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day23TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(12521, st.sol.Part1())
	st.Equal(44169, st.sol.Part2())
}

func TestDay23(t *testing.T) {
	st := new(Day23TestSuite)
	suite.Run(t, st)
}
