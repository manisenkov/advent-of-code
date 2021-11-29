package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
abc

a
b
c

ab
ac

a
a
a
a

b
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
	st.Equal(11, st.sol.Part1())
	st.Equal(6, st.sol.Part2())
}

func TestDay06(t *testing.T) {
	st := new(Day06TestSuite)
	suite.Run(t, st)
}
