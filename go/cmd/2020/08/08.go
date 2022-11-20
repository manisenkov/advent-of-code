package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type instruction struct {
	cmd string
	arg int
}

// Solution contains solution for day 8
type Solution struct {
	program []instruction
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.program = make([]instruction, len(input))
	for i, inp := range input {
		ins := strings.Split(inp, " ")
		sol.program[i] = instruction{
			cmd: ins[0],
			arg: common.MustAtoi(ins[1]),
		}
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	acc, _ := runProg(sol.program)
	return acc
}

// Part2 .
func (sol *Solution) Part2() any {
	for i, ins := range sol.program {
		if ins.cmd == "acc" {
			continue
		}
		prog := make([]instruction, len(sol.program))
		copy(prog, sol.program)
		if prog[i].cmd == "nop" {
			prog[i].cmd = "jmp"
		} else {
			prog[i].cmd = "nop"
		}
		acc, isLoop := runProg(prog)
		if !isLoop {
			return acc
		}
	}
	return 0
}

func runProg(prog []instruction) (int, bool) {
	pos := 0
	acc := 0
	isExec := map[int]bool{}
	for !isExec[pos] && pos < len(prog) {
		ins := prog[pos]
		isExec[pos] = true
		switch ins.cmd {
		case "nop":
			pos++
		case "acc":
			acc += ins.arg
			pos++
		case "jmp":
			pos += ins.arg
		}
	}
	return acc, isExec[pos]
}

func main() {
	common.Run(new(Solution))
}
