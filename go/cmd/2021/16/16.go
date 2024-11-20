package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type packet interface {
	eval() int64
	getVersion() int
}

type operator struct {
	version int
	typeID  int
	args    []packet
}

func (o operator) eval() int64 {
	res := int64(0)
	switch o.typeID {
	case 0:
		for _, arg := range o.args {
			res += arg.eval()
		}
	case 1:
		res = int64(1)
		for _, arg := range o.args {
			res *= arg.eval()
		}
	case 2:
		res = int64(0x7fffffffffffffff)
		for _, arg := range o.args {
			x := arg.eval()
			if x < res {
				res = x
			}
		}
	case 3:
		res = int64(-0x7fffffffffffffff)
		for _, arg := range o.args {
			x := arg.eval()
			if x > res {
				res = x
			}
		}
	case 5:
		if o.args[0].eval() > o.args[1].eval() {
			res = int64(1)
		}
	case 6:
		if o.args[0].eval() < o.args[1].eval() {
			res = int64(1)
		}
	case 7:
		if o.args[0].eval() == o.args[1].eval() {
			res = int64(1)
		}
	}
	return res
}

func (o operator) getVersion() int {
	res := o.version
	for _, p := range o.args {
		res += p.getVersion()
	}
	return res
}

type literal struct {
	version int
	typeID  int
	num     int64
}

func (l literal) eval() int64 {
	return l.num
}

func (l literal) getVersion() int {
	return l.version
}

var hexMap map[rune][]int = map[rune][]int{
	'0': {0, 0, 0, 0},
	'1': {0, 0, 0, 1},
	'2': {0, 0, 1, 0},
	'3': {0, 0, 1, 1},
	'4': {0, 1, 0, 0},
	'5': {0, 1, 0, 1},
	'6': {0, 1, 1, 0},
	'7': {0, 1, 1, 1},
	'8': {1, 0, 0, 0},
	'9': {1, 0, 0, 1},
	'A': {1, 0, 1, 0},
	'B': {1, 0, 1, 1},
	'C': {1, 1, 0, 0},
	'D': {1, 1, 0, 1},
	'E': {1, 1, 1, 0},
	'F': {1, 1, 1, 1},
}

// Solution contains solution for day 16
type Solution struct {
	hexString string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.hexString = input[0]
}

// Part1 .
func (sol *Solution) Part1() any {
	bytes := []int{}
	for _, h := range sol.hexString {
		bytes = append(bytes, hexMap[h]...)
	}
	packet, _ := parse(bytes, 0)
	return packet.getVersion()
}

// Part2 .
func (sol *Solution) Part2() any {
	bytes := []int{}
	for _, h := range sol.hexString {
		bytes = append(bytes, hexMap[h]...)
	}
	packet, _ := parse(bytes, 0)
	return packet.eval()
}

func parse(bytes []int, pos int) (packet, int) {
	version := bytes[pos+0]*4 + bytes[pos+1]*2 + bytes[pos+2]
	typeID := bytes[pos+3]*4 + bytes[pos+4]*2 + bytes[pos+5]
	pos += 6
	if typeID == 4 { // Literal value
		numStr := ""
		for {
			isLastGroup := bytes[pos] == 0
			numStr += bytesToString(bytes[pos+1 : pos+5])
			pos += 5
			if isLastGroup {
				break
			}
		}
		num := numbers.MustParseInt[int64](numStr, 2)
		return literal{version, typeID, num}, pos
	} else {
		lengthTypeID := bytes[pos]
		args := []packet{}
		pos++
		if lengthTypeID == 0 {
			length := numbers.MustParseInt[int](bytesToString(bytes[pos:pos+15]), 2)
			pos += 15
			stopPos := pos + length
			for pos < stopPos {
				arg, nextPos := parse(bytes, pos)
				args = append(args, arg)
				pos = nextPos
			}
		} else {
			left := numbers.MustParseInt[int](bytesToString(bytes[pos:pos+11]), 2)
			pos += 11
			for left > 0 {
				arg, nextPos := parse(bytes, pos)
				args = append(args, arg)
				pos = nextPos
				left--
			}
		}
		return operator{version, typeID, args}, pos
	}
}

func bytesToString(bytes []int) string {
	res := ""
	for _, b := range bytes {
		if b == 1 {
			res += "1"
		} else {
			res += "0"
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
