package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type operation string

const (
	opAnd operation = "AND"
	opOr  operation = "OR"
	opXor operation = "XOR"
)

type connection struct {
	op    operation
	input [2]string
}

func (c connection) run(inputs map[string]bool) (bool, bool) {
	inp1, ok1 := inputs[c.input[0]]
	inp2, ok2 := inputs[c.input[1]]
	if !ok1 || !ok2 {
		return false, false
	}
	var res bool
	switch c.op {
	case opAnd:
		res = inp1 && inp2
	case opOr:
		res = inp1 || inp2
	case opXor:
		res = inp1 != inp2
	}
	return res, true
}

// Solution contains a solution for day 24
type Solution struct {
	initInput   map[string]bool
	connections map[string]connection
	size        int
}

func findConnection(input [2]string, op operation, connections map[string]connection) string {
	for output, c := range connections {
		if c.op == op && ((c.input[0] == input[0] && c.input[1] == c.input[1]) || (c.input[0] == input[1] && c.input[1] == input[0])) {
			return output
		}
	}
	return ""
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.connections = make(map[string]connection)
	sol.initInput = make(map[string]bool)
	i := 0
	for ; input[i] != ""; i++ {
		parts := strings.Split(input[i], ": ")
		sol.initInput[parts[0]] = (parts[1] == "1")
	}
	i++

	for ; i < len(input); i++ {
		parts := strings.Split(input[i], " -> ")
		inputParts := strings.Split(parts[0], " ")
		sol.connections[parts[1]] = connection{
			op:    operation(inputParts[1]),
			input: [2]string{inputParts[0], inputParts[2]},
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	inputs := maps.Clone(sol.initInput)
	toProcess := sol.connections
	for len(toProcess) > 0 {
		next := make(map[string]connection)
		for output, c := range toProcess {
			r, ok := c.run(inputs)
			if ok {
				inputs[output] = r
			} else {
				next[output] = c
			}
		}
		toProcess = next
	}
	res := 0
	for i := 0; ; i++ {
		key := fmt.Sprintf("z%02d", i)
		if x, ok := inputs[key]; ok {
			if x {
				res += 1 << i
			}
		} else {
			break
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	/*
		# Full adder scheme
		# From https://www.electronics-tutorials.ws/combination/comb_7.html

		t[i] = x[i] ^ y[i]
		d[i] = t[i] ^ carry[i-1]
		a[i] = x[i] & y[i]
		b[i] = t[i] & carry[i-1]
		carry[i] = a[i] | b[i]
	*/

	connections := maps.Clone(sol.connections)
	res := make([]string, 0)
	carry := ""
	for len(res) < 8 {
		for i := 0; ; i++ {
			xIdx := fmt.Sprintf("x%02d", i)
			yIdx := fmt.Sprintf("y%02d", i)
			zIdx := fmt.Sprintf("z%02d", i)

			a := findConnection([2]string{xIdx, yIdx}, opAnd, connections)

			if i == 0 {
				carry = a
				continue
			}

			t := findConnection([2]string{xIdx, yIdx}, opXor, connections)
			d := findConnection([2]string{t, carry}, opXor, connections)

			if d == "" {
				res = append(res, t, a)
				connections[t], connections[a] = connections[a], connections[t]
				break
			}
			if d != zIdx {
				res = append(res, d, zIdx)
				connections[d], connections[zIdx] = connections[zIdx], connections[d]
				break
			}

			b := findConnection([2]string{t, carry}, opAnd, connections)
			carry = findConnection([2]string{a, b}, opOr, connections)
		}
	}
	slices.Sort(res)
	return strings.Join(res, ",")
}

func main() {
	problem.Solve(new(Solution))
}
