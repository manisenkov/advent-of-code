package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
029A
980A
179A
456A
379A`

type Day21TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day21TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day21TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(126384, st.sol.Part1())
	st.Equal(154115708116294, st.sol.Part2())
}

func TestDay21(t *testing.T) {
	st := new(Day21TestSuite)
	suite.Run(t, st)
}
