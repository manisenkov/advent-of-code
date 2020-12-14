package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

type Day14TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day14TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day14TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(165), st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay14(t *testing.T) {
	st := new(Day14TestSuite)
	suite.Run(t, st)
}
