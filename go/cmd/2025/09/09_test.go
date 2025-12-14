package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

type Day09TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day09TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day09TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(50, st.sol.Part1())
	st.Equal(24, st.sol.Part2())
}

func TestDay09(t *testing.T) {
	st := new(Day09TestSuite)
	suite.Run(t, st)
}
