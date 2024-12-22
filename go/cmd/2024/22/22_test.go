package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
1
10
100
2024`

const testInput2 = `
1
2
3
2024`

type Day22TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day22TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day22TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(37327623), st.sol.Part1())
	st.Equal(int64(24), st.sol.Part2())
}

func (st *Day22TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.sol.Part1()
	st.Equal(int64(37990510), st.sol.Part1())
	st.Equal(int64(23), st.sol.Part2())
}

func TestDay22(t *testing.T) {
	st := new(Day22TestSuite)
	suite.Run(t, st)
}
