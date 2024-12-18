package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

type Day18TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day18TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day18TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.sol.numBytes = 12
	st.sol.size = [2]int{7, 7}
	st.Equal(22, st.sol.Part1())
	st.Equal("6,1", st.sol.Part2())
}

func TestDay18(t *testing.T) {
	st := new(Day18TestSuite)
	suite.Run(t, st)
}
