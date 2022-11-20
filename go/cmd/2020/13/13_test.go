package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testInput1 = `
939
7,13,x,x,59,x,31,19
`
	testInput2 = `
0
17,x,13,19
`
	testInput3 = `
0
67,7,59,61
`
	testInput4 = `
0
67,x,7,59,61
`
	testInput5 = `
0
67,7,x,59,61
`
	testInput6 = `
0
1789,37,47,1889
`
)

type Day13TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day13TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day13TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(295), st.sol.Part1())
	st.Equal(int64(1068781), st.sol.Part2())
}

func (st *Day13TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(3417), st.sol.Part2())
}

func (st *Day13TestSuite) Test3() {
	input := strings.Split(strings.Trim(testInput3, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(754018), st.sol.Part2())
}

func (st *Day13TestSuite) Test4() {
	input := strings.Split(strings.Trim(testInput4, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(779210), st.sol.Part2())
}

func (st *Day13TestSuite) Test5() {
	input := strings.Split(strings.Trim(testInput5, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(1261476), st.sol.Part2())
}

func (st *Day13TestSuite) Test6() {
	input := strings.Split(strings.Trim(testInput6, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(1202161486), st.sol.Part2())
}

func TestDay13(t *testing.T) {
	st := new(Day13TestSuite)
	suite.Run(t, st)
}
