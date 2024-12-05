package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

type Day05TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day05TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day05TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(143, st.sol.Part1())
	st.Equal(123, st.sol.Part2())
}

func TestDay05(t *testing.T) {
	st := new(Day05TestSuite)
	suite.Run(t, st)
}
