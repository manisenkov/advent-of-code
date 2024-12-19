package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

type Day19TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day19TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day19TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(6, st.sol.Part1())
	st.Equal(16, st.sol.Part2())
}

func TestDay19(t *testing.T) {
	st := new(Day19TestSuite)
	suite.Run(t, st)
}
