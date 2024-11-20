package main

import (
	"sort"

	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type byteStack struct {
	stack []byte
}

func (s *byteStack) size() int {
	return len(s.stack)
}

func (s *byteStack) push(b byte) {
	s.stack = append(s.stack, b)
}

func (s *byteStack) peek() byte {
	return s.stack[len(s.stack)-1]
}

func (s *byteStack) pop() byte {
	el := s.peek()
	s.stack = s.stack[:len(s.stack)-1]
	return el
}

// Solution contains solution for day 10
type Solution struct {
	lines           [][]byte
	incompleteLines [][]byte
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	lines := make([][]byte, len(input))
	for i, s := range input {
		lines[i] = []byte(s)
	}
	sol.lines = lines
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for _, line := range sol.lines {
		switch testLine(line) {
		case ')':
			res += 3
		case ']':
			res += 57
		case '}':
			res += 1197
		case '>':
			res += 25137
		default:
			sol.incompleteLines = append(sol.incompleteLines, line)
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	scores := make([]int, len(sol.incompleteLines))
	for i, line := range sol.incompleteLines {
		completion := completeLine(line)
		lineScore := 0
		for _, b := range completion {
			lineScore *= 5
			switch b {
			case ')':
				lineScore += 1
			case ']':
				lineScore += 2
			case '}':
				lineScore += 3
			case '>':
				lineScore += 4
			}
		}
		scores[i] = lineScore
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func completeLine(line []byte) []byte {
	stack := &byteStack{}
	for _, b := range line {
		switch b {
		case '(', '<', '{', '[':
			stack.push(b)
		case ')', '>', '}', ']':
			stack.pop()
		}
	}
	res := make([]byte, stack.size())
	for i := 0; i < len(res); i++ {
		p := stack.pop()
		switch p {
		case '(':
			res[i] = ')'
		case '[':
			res[i] = ']'
		case '{':
			res[i] = '}'
		case '<':
			res[i] = '>'
		}
	}
	return res
}

func testLine(line []byte) byte {
	stack := &byteStack{}
	for _, b := range line {
		switch b {
		case '(', '<', '{', '[':
			stack.push(b)
		case ')', '>', '}', ']':
			p := stack.peek()
			if p == '(' && b == ')' ||
				p == '[' && b == ']' ||
				p == '{' && b == '}' ||
				p == '<' && b == '>' {
				stack.pop()
			} else {
				return b
			}
		}
	}
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
