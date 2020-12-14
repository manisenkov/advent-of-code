package main

import (
	"regexp"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

var (
	maskRegex = regexp.MustCompile(`^mask = ((1|0|X){36})$`)
	memRegex  = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

type state struct {
	mem         map[int64]int64
	currentMask mask
}

type command interface {
	Run(st *state)
}

type mask struct {
	sieve int64
	value int64
}

func (m mask) Run(st *state) {
	st.currentMask = m
}

func parseMask(inp string) mask {
	m := mask{}
	for _, c := range inp {
		switch c {
		case '1':
			m.sieve <<= 1
			m.value = (m.value << 1) | 1
		case '0':
			m.sieve <<= 1
			m.value <<= 1
		case 'X':
			m.sieve = (m.sieve << 1) | 1
			m.value <<= 1
		}
	}
	return m
}

type setMemCommand struct {
	address int64
	value   int64
}

func (c setMemCommand) Run(st *state) {
	st.mem[c.address] = st.currentMask.value | (c.value & st.currentMask.sieve)
}

// Solution contains solution for day 14
type Solution struct {
	commands []command
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.commands = make([]command, len(input))
	for i, inp := range input {
		m := maskRegex.FindAllStringSubmatch(inp, -1)
		if len(m) > 0 {
			sol.commands[i] = parseMask(m[0][1])
			continue
		}
		m = memRegex.FindAllStringSubmatch(inp, -1)
		if len(m) > 0 {
			sol.commands[i] = setMemCommand{
				address: common.MustParseInt(m[0][1], 10, 64),
				value:   common.MustParseInt(m[0][2], 10, 64),
			}
			continue
		}
		panic("Don't understand the command")
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	st := state{
		currentMask: mask{},
		mem: make(map[int64]int64),
	}
	for _, cmd := range sol.commands {
		cmd.Run(&st)
	}
	res := int64(0)
	for _, val := range st.mem {
		res += val
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	return 0
}

func main() {
	common.Run(new(Solution))
}
