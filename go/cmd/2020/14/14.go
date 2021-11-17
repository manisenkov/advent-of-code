package main

import (
	"regexp"

	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

const (
	direct  = iota
	address = iota
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
	run(st *state, mode int)
}

type mask struct {
	sieve int64
	value int64
}

func (m mask) run(st *state, _ int) {
	st.currentMask = m
}

func (m mask) getAddresses(arg int64) []int64 {
	bits := []int64{}
	rest := m.sieve
	for i := 0; i < 36; i++ {
		if rest&1 == 1 {
			bits = append(bits, 1<<i)
		}
		rest >>= 1
	}

	totalCases := int64(1) << len(bits)
	res := make([]int64, totalCases)
	for i := int64(0); i < totalCases; i++ {
		var address int64
		idxs := getBitIndexes(i)
		for idx := range idxs {
			address += bits[idx]
		}
		res[i] = (address | m.value) | (^(m.sieve | m.value) & arg)
	}
	return res
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

func (c setMemCommand) run(st *state, mode int) {
	switch mode {
	case direct:
		st.mem[c.address] = st.currentMask.value | (c.value & st.currentMask.sieve)
	case address:
		addresses := st.currentMask.getAddresses(c.address)
		for _, address := range addresses {
			st.mem[address] = c.value
		}
	}
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
		mem:         make(map[int64]int64),
	}
	for _, cmd := range sol.commands {
		cmd.run(&st, direct)
	}
	res := int64(0)
	for _, val := range st.mem {
		res += val
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	st := state{
		currentMask: mask{},
		mem:         make(map[int64]int64),
	}
	for _, cmd := range sol.commands {
		cmd.run(&st, address)
	}
	res := int64(0)
	for _, val := range st.mem {
		res += val
	}
	return res
}

func getBitIndexes(arg int64) map[int]bool {
	res := make(map[int]bool)
	i := 0
	for arg > 0 {
		if arg&1 == 1 {
			res[i] = true
		}
		i++
		arg >>= 1
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
