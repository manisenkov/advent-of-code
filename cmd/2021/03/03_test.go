package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

type Day03TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day03TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day03TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(198, st.sol.Part1())
	st.Equal(230, st.sol.Part2())
}

func TestDay03(t *testing.T) {
	st := new(Day03TestSuite)
	suite.Run(t, st)
}
