package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

type Day25TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day25TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day25TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(3, st.sol.Part1())
	st.Equal("HO HO HO", st.sol.Part2())
}

func TestDay25(t *testing.T) {
	st := new(Day25TestSuite)
	suite.Run(t, st)
}
