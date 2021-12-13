package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`

type Day13TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day13TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day13TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(17, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay13(t *testing.T) {
	st := new(Day13TestSuite)
	suite.Run(t, st)
}
