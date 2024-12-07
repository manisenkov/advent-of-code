package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

type Day07TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day07TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day07TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(3749, st.sol.Part1())
	st.Equal(11387, st.sol.Part2())
}

func TestDay07(t *testing.T) {
	st := new(Day07TestSuite)
	suite.Run(t, st)
}
