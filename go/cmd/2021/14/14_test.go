package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
`

type Day14TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day14TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day14TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(1588), st.sol.Part1())
	st.Equal(int64(2188189693529), st.sol.Part2())
}

func TestDay14(t *testing.T) {
	st := new(Day14TestSuite)
	suite.Run(t, st)
}
