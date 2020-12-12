package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
F10
N3
W3
S3
E3
N3
F7
R90
L90
L180
R180
L270
R270
L270
F11
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
	st.Equal(25, st.sol.Part1())
	st.Equal(286, st.sol.Part2())
}

func TestDay12(t *testing.T) {
	st := new(Day12TestSuite)
	suite.Run(t, st)
}
