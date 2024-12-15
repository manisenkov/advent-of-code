package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

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
	st.sol.size = [2]int{11, 7}
	st.Equal(12, st.sol.Part1())
	st.Equal(1, st.sol.Part2())
}

func TestDay14(t *testing.T) {
	st := new(Day14TestSuite)
	suite.Run(t, st)
}
