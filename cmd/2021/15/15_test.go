package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`

type Day15TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day15TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day15TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(40, st.sol.Part1())
	st.Equal(315, st.sol.Part2())
}

func TestDay15(t *testing.T) {
	st := new(Day15TestSuite)
	suite.Run(t, st)
}
