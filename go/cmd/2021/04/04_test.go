package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`

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
	st.Equal(4512, st.sol.Part1())
	st.Equal(1924, st.sol.Part2())
}

func TestDay04(t *testing.T) {
	st := new(Day04TestSuite)
	suite.Run(t, st)
}
