package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
Player 1 starting position: 4
Player 2 starting position: 8
`

type Day21TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day21TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day21TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(739785, st.sol.Part1())
	st.Equal(int64(444356092776315), st.sol.Part2())
}

func TestDay21(t *testing.T) {
	st := new(Day21TestSuite)
	suite.Run(t, st)
}
