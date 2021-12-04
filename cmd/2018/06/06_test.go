package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

type Day06TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day06TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day06TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.sol.maxDist = 32 // for test purposes
	st.Equal(17, st.sol.Part1())
	st.Equal(16, st.sol.Part2())
}

func TestDay06(t *testing.T) {
	st := new(Day06TestSuite)
	suite.Run(t, st)
}
