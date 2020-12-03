package main

const dayTemplateStr = `package day{{.PaddedDay}}

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day {{.Day}}
type Solution struct {
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
}

// Part1 .
func (s *Solution) Part1() common.Any {
  return 0
}

// Part2 .
func (s *Solution) Part2() common.Any {
  return 0
}
`

const testTemplateStr = `package day{{.PaddedDay}}

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day{{.PaddedDay}}TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day{{.PaddedDay}}TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day{{.PaddedDay}}TestSuite) Test1() {
	s.sol.Init([]string{"foo"})
	s.Equal(0, s.sol.Part1())
	s.Equal(0, s.sol.Part2())
}

func TestDay{{.PaddedDay}}(t *testing.T) {
	s := new(Day{{.PaddedDay}}TestSuite)
	suite.Run(t, s)
}
`
