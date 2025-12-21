package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
16x7: 1 0 1 0 3 2
`

type Day12TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day12TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day12TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(1, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay12(t *testing.T) {
	st := new(Day12TestSuite)
	suite.Run(t, st)
}
