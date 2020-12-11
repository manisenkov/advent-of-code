package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

type Day11TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day11TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day11TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(37, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay11(t *testing.T) {
	st := new(Day11TestSuite)
	suite.Run(t, st)
}
