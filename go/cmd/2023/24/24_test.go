package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
`

type Day24TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day24TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day24TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.sol.area = [2]int64{7, 27}
	st.Equal(2, st.sol.Part1())
	st.Equal(int64(47), st.sol.Part2())
}

func TestDay24(t *testing.T) {
	st := new(Day24TestSuite)
	suite.Run(t, st)
}
