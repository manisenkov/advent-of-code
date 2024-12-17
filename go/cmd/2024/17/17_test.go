package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

const testInput2 = `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

type Day17TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day17TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day17TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal("4,6,3,5,6,3,5,2,1,0", st.sol.Part1())
}

func (st *Day17TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(117440), st.sol.Part2())
}

func TestDay17(t *testing.T) {
	st := new(Day17TestSuite)
	suite.Run(t, st)
}
