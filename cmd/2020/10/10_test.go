package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
16
10
15
5
1
11
7
19
6
12
4
`

const testInput2 = `
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

type Day10TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day10TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day10TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(35, st.sol.Part1())
	st.Equal(8, st.sol.Part2())
}

func (st *Day10TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(220, st.sol.Part1())
	st.Equal(19208, st.sol.Part2())
}

func TestDay10(t *testing.T) {
	st := new(Day10TestSuite)
	suite.Run(t, st)
}
