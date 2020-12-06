package main

const dayTemplateStr = `package day{{.PaddedDay}}

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day {{.Day}}
type Solution struct {
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
}

// Part1 .
func (sol *Solution) Part1() common.Any {
  return 0
}

// Part2 .
func (sol *Solution) Part2() common.Any {
  return 0
}
`

const testTemplateStr = `package day{{.PaddedDay}}

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = ` + "`" + `
hello
world
` + "`" + `

type Day{{.PaddedDay}}TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day{{.PaddedDay}}TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day{{.PaddedDay}}TestSuite) Test1() {
	strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(0, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay{{.PaddedDay}}(t *testing.T) {
	st := new(Day{{.PaddedDay}}TestSuite)
	suite.Run(t, st)
}
`
