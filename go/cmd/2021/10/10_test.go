package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`

type Day10TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day10TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day10TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(26397, st.sol.Part1())
	st.Equal(288957, st.sol.Part2())
}

func TestDay10(t *testing.T) {
	st := new(Day10TestSuite)
	suite.Run(t, st)
}
