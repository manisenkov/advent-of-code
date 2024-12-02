package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day02TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day02TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(2, st.sol.Part1())
	st.Equal(4, st.sol.Part2())
}

func TestDay02(t *testing.T) {
	st := new(Day02TestSuite)
	suite.Run(t, st)
}
