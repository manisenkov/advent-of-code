package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day02TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day02TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(2, st.sol.Part1())
	st.Equal(1, st.sol.Part2())
}

func TestDay02(t *testing.T) {
	st := new(Day02TestSuite)
	suite.Run(t, st)
}
