package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var testInputs = map[string][2]int{
	"(())":    {0, 0},
	"()()":    {0, 0},
	"(((":     {3, 0},
	"(()(()(": {3, 0},
	"))(((((": {3, 1},
	"())":     {-1, 3},
	"))(":     {-1, 1},
	")))":     {-3, 1},
	")())())": {-3, 1},
}

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) Test1() {
	for testInput, res := range testInputs {
		st.sol.Init([]string{testInput})
		st.Equal(res[0], st.sol.Part1())
		st.Equal(res[1], st.sol.Part2())
	}
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
