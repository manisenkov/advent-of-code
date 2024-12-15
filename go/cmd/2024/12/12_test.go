package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

const testInput2 = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`

type Day12TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day12TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day12TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(1930, st.sol.Part1())
	st.Equal(1206, st.sol.Part2())
}

func (st *Day12TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(1184, st.sol.Part1())
	st.Equal(368, st.sol.Part2())
}

func TestDay12(t *testing.T) {
	st := new(Day12TestSuite)
	suite.Run(t, st)
}
