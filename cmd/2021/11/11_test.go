package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
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
	st.Equal(1656, st.sol.Part1())
	st.Equal(195, st.sol.Part2())
}

func TestDay11(t *testing.T) {
	st := new(Day11TestSuite)
	suite.Run(t, st)
}
