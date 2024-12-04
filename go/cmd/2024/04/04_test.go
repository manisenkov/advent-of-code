package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

type Day04TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day04TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day04TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(18, st.sol.Part1())
	st.Equal(9, st.sol.Part2())
}

func TestDay04(t *testing.T) {
	st := new(Day04TestSuite)
	suite.Run(t, st)
}
