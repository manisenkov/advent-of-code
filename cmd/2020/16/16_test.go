package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testInput1 = `
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`
	testInput2 = `
departure class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`
)

type Day16TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day16TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day16TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(71, st.sol.Part1())
}

func (st *Day16TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(0, st.sol.Part1())
	st.Equal(12, st.sol.Part2())
}

func TestDay16(t *testing.T) {
	st := new(Day16TestSuite)
	suite.Run(t, st)
}
