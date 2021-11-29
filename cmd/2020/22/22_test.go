package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
`

type Day22TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day22TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day22TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(306, st.sol.Part1())
	st.Equal(291, st.sol.Part2())
}

func TestDay22(t *testing.T) {
	st := new(Day22TestSuite)
	suite.Run(t, st)
}
