package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
3,4,3,1,2
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
	st.Equal(int64(5934), st.sol.Part1())
	st.Equal(int64(26984457539), st.sol.Part2())
}

func TestDay06(t *testing.T) {
	st := new(Day06TestSuite)
	suite.Run(t, st)
}
