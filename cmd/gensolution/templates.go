package main

const dayTemplateStr = `package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains a solution for day {{.Day}}
type Solution struct {
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() common.Any {
	return 0
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() common.Any {
	return 0
}

func main() {
	common.Run(new(Solution))
}
`

const testTemplateStr = `package main

import (
	"strings"
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
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(0, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func TestDay{{.PaddedDay}}(t *testing.T) {
	st := new(Day{{.PaddedDay}}TestSuite)
	suite.Run(t, st)
}
`
