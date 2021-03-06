package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
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
	st.Equal(4, st.sol.Part1())
	st.Equal("3", st.sol.Part2())
}

func TestDay03(t *testing.T) {
	st := new(Day03TestSuite)
	suite.Run(t, st)
}
