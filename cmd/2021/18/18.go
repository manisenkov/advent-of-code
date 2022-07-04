package main

import (
	"fmt"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type node struct {
	value int
	level int
}

func (n node) String() string {
	return fmt.Sprintf("[\"%v\" L:%v]", n.value, n.level)
}

type number []node

func parseNumber(input string) number {
	level := 0
	res := number{}
	for _, b := range input {
		if b == '[' {
			level++
		} else if b == ']' {
			level--
		} else if b == ',' {
			continue
		} else {
			res = append(res, node{int(b - '0'), level})
		}
	}
	return res
}

func (num number) add(other number) number {
	res := make(number, len(num)+len(other))
	copy(res, num)
	copy(res[len(num):], other)
	for i := range res {
		res[i].level++
	}
	return res.reduce()
}

func (num number) explode() (number, bool) {
	for i, n := range num {
		if n.level == 5 {
			res := make(number, len(num)-1)
			copy(res, num[:i])
			res[i] = node{value: 0, level: 4}
			copy(res[i+1:], num[i+2:])
			if i > 0 {
				res[i-1].value += num[i].value
			}
			if i < len(res)-1 {
				res[i+1].value += num[i+1].value
			}
			return res, true
		}
	}
	return num, false
}

func (num number) magnitude() int {
	res := make(number, len(num))
	copy(res, num)
	for len(res) > 1 {
		for i := range res {
			if i < len(res)-1 && res[i].level == res[i+1].level {
				res[i] = node{
					value: 3*res[i].value + 2*res[i+1].value,
					level: res[i].level - 1,
				}
				res = append(res[:i+1], res[i+2:]...)
				break
			}
		}
	}
	return res[0].value
}

func (num number) reduce() number {
	res := num
	for {
		var okExplode, okSplit bool
		res, okExplode = res.explode()
		if !okExplode {
			res, okSplit = res.split()
		}
		if !okExplode && !okSplit {
			break
		}
	}
	return res
}

func (num number) split() (number, bool) {
	for i, n := range num {
		if n.value > 9 {
			res := make(number, len(num)+1)
			copy(res, num[:i])
			res[i] = node{num[i].value / 2, num[i].level + 1}
			res[i+1] = node{num[i].value - num[i].value/2, num[i].level + 1}
			copy(res[i+2:], num[i+1:])
			return res, true
		}
	}
	return num, false
}

// Solution contains solution for day 18
type Solution struct {
	numbers []number
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.numbers = make([]number, len(input))
	for i, s := range input {
		sol.numbers[i] = parseNumber(s)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := sol.numbers[0]
	for _, num := range sol.numbers[1:] {
		res = res.add(num)
	}
	return res.magnitude()
}

// Part2 .
func (sol *Solution) Part2() any {
	maxMagnitude := -0x7fffffff
	for i := range sol.numbers {
		for j := range sol.numbers {
			if i == j {
				continue
			}
			m := sol.numbers[i].add(sol.numbers[j]).magnitude()
			if m > maxMagnitude {
				maxMagnitude = m
			}
		}
	}
	return maxMagnitude
}

func main() {
	common.Run(new(Solution))
}
