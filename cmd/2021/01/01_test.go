package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
199
200
208
210
200
207
240
269
260
263
`

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(7, st.sol.Part1())
	st.Equal(5, st.sol.Part2())
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
