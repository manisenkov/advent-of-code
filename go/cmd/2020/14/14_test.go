package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

const testInput2 = `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`

type Day14TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day14TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day14TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(165), st.sol.Part1())
}

func (st *Day14TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(208), st.sol.Part2())
}

func TestDay14(t *testing.T) {
	st := new(Day14TestSuite)
	suite.Run(t, st)
}
