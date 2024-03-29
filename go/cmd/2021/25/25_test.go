package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
`

type Day25TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day25TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day25TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(58, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay25(t *testing.T) {
	st := new(Day25TestSuite)
	suite.Run(t, st)
}
