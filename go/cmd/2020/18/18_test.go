package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`

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
	st.Equal(71+51+26+437+12240+13632, st.sol.Part1())
	st.Equal(231+51+46+1445+669060+23340, st.sol.Part2())
}

func TestDay18(t *testing.T) {
	st := new(Day18TestSuite)
	suite.Run(t, st)
}
