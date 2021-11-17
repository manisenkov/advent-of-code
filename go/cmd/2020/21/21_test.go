package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`

type Day21TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day21TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day21TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(5, st.sol.Part1())
	st.Equal("mxmxvkd,sqjhc,fvjkl", st.sol.Part2())
}

func TestDay21(t *testing.T) {
	st := new(Day21TestSuite)
	suite.Run(t, st)
}
