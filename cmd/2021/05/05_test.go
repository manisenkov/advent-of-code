package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

type Day05TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day05TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day05TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(5, st.sol.Part1())
	st.Equal(12, st.sol.Part2())
}

func TestDay05(t *testing.T) {
	st := new(Day05TestSuite)
	suite.Run(t, st)
}
