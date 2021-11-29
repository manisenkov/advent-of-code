package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`

type Day05TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day05TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day05TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(820, st.sol.Part1())
}

func TestDay05(t *testing.T) {
	st := new(Day05TestSuite)
	suite.Run(t, st)
}
