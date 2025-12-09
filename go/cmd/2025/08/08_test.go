package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
`

type Day08TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day08TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day08TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.sol.numSteps = 10
	st.Equal(40, st.sol.Part1())
	st.Equal(25272, st.sol.Part2())
}

func TestDay08(t *testing.T) {
	st := new(Day08TestSuite)
	suite.Run(t, st)
}
