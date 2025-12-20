package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

type Day10TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day10TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day10TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(7, st.sol.Part1())
	st.Equal(33, st.sol.Part2())
}

func TestDay10(t *testing.T) {
	st := new(Day10TestSuite)
	suite.Run(t, st)
}
