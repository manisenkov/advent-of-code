package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
2199943210
3987894921
9856789892
8767896789
9899965678
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
	st.Equal(15, st.sol.Part1())
	st.Equal(1134, st.sol.Part2())
}

func TestDay09(t *testing.T) {
	st := new(Day09TestSuite)
	suite.Run(t, st)
}
