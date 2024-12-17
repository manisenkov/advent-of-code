package main

import (
	"maps"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type instruction int64

const (
	adv instruction = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func combo(op int64, registers map[string]int64) int64 {
	if op >= 0 && op <= 3 {
		return op
	}
	if op == 4 {
		return registers["A"]
	}
	if op == 5 {
		return registers["B"]
	}
	if op == 6 {
		return registers["C"]
	}
	return -1
}

func run(ptr int, program []int64, regs map[string]int64) (int, []int64) {
	ins := instruction(program[ptr])
	op := program[ptr+1]
	res := []int64{}
	switch ins {
	case adv:
		regs["A"] = regs["A"] / (1 << combo(op, regs))
		ptr += 2
	case bxl:
		regs["B"] ^= op
		ptr += 2
	case bst:
		regs["B"] = combo(op, regs) % 8
		ptr += 2
	case jnz:
		if regs["A"] == 0 {
			ptr += 2
		} else {
			ptr = int(op)
		}
	case bxc:
		regs["B"] ^= regs["C"]
		ptr += 2
	case out:
		res = append(res, combo(op, regs)%8)
		ptr += 2
	case bdv:
		regs["B"] = regs["A"] / (1 << combo(op, regs))
		ptr += 2
	case cdv:
		regs["C"] = regs["A"] / (1 << combo(op, regs))
		ptr += 2
	}
	return ptr, res
}

func execute(regs map[string]int64, program []int64) []int64 {
	output := []int64{}
	ptr := 0
	for ptr < len(program) {
		var res []int64
		ptr, res = run(ptr, program, regs)
		output = append(output, res...)
	}
	return output
}

func solve(input int64, pos int, conns map[int][]int, eval func(int64) int) (int64, bool) {
	if pos == -1 {
		return input, true
	}
	options := []int64{}
	for _, cmb := range collections.AllCombinations(conns[pos]) {
		diff := int64(0)
		for _, shift := range cmb {
			diff += 1 << shift
		}
		diffPos := eval(input + diff)
		if diffPos < pos {
			options = append(options, diff)
		}
	}
	for _, opt := range options {
		res, ok := solve(input+opt, pos-1, conns, eval)
		if ok {
			return res, true
		}
	}
	return -1, false
}

// Solution contains a solution for day 17
type Solution struct {
	registers map[string]int64
	program   []int64
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.registers = map[string]int64{
		"A": numbers.MustAtoi[int64](strings.Split(input[0], ": ")[1]),
		"B": numbers.MustAtoi[int64](strings.Split(input[1], ": ")[1]),
		"C": numbers.MustAtoi[int64](strings.Split(input[2], ": ")[1]),
	}
	sol.program = collections.MapTo(strings.Split(strings.Split(input[4], ": ")[1], ","), numbers.MustAtoi[int64])
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	output := execute(maps.Clone(sol.registers), sol.program)
	return strings.Join(collections.MapTo(output, numbers.Itoa), ",")
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	// Direct solution
	if len(sol.program) < 10 {
		target := strings.Join(collections.MapTo(sol.program, numbers.Itoa), ",")
		res := int64(0)
		for {
			output := execute(map[string]int64{"A": res}, sol.program)
			if strings.Join(collections.MapTo(output, numbers.Itoa), ",") == target {
				return res
			} else {
				res++
			}
		}
	}

	// Complex solution for bigger case
	conns := map[int][]int{}
	for i := 0; i < len(sol.program); i++ {
		conns[i] = []int{}
	}
	for i := 0; i < len(sol.program)*3; i++ {
		output := execute(map[string]int64{"A": 1 << i}, sol.program)
		for j, n := range output {
			if n != 0 {
				conns[j] = append(conns[j], i)
			}
		}
	}
	eval := func(a int64) int {
		output := execute(map[string]int64{"A": a}, sol.program)
		if len(output) != len(sol.program) {
			return len(sol.program) - 1
		}
		for i := len(sol.program) - 1; i >= 0; i-- {
			if sol.program[i] != output[i] {
				return i
			}
		}
		return -1
	}
	res, ok := solve(0, len(sol.program)-1, conns, eval)
	if !ok {
		panic("solution failed")
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
