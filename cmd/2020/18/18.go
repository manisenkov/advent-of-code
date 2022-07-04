package main

import (
	"fmt"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

const (
	tlpar = "("
	trpar = ")"
	tplus = "+"
	tmult = "*"
	tnum  = "N"
	tend  = "$"
)

type token struct {
	value string
	typ   string
}

// Solution contains solution for day 18
type Solution struct {
	exprs [][]token
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.exprs = make([][]token, len(input))
	for i, inp := range input {
		sol.exprs[i] = parseTokens([]rune(inp + tend))
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for _, expr := range sol.exprs {
		r, _ := eval(expr, 0)
		res += r
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	res := 0
	for _, expr := range sol.exprs {
		r, _ := evalOrder(expr, 0, 0)
		res += r
	}
	return res
}

func detTokenType(r rune) string {
	switch r {
	case '+':
		return tplus
	case '*':
		return tmult
	case '(':
		return tlpar
	case ')':
		return trpar
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return tnum
	case '$':
		return tend
	}
	panic("Unknown character")
}

func parseTokens(input []rune) []token {
	tokens := make([]token, 0)
	value := ""
	curTokenType := ""
	i := 0
	for i < len(input) {
		r := input[i]
		if r == ' ' {
			i++
			continue
		}
		tokenType := detTokenType(r)
		if curTokenType == "" {
			curTokenType = tokenType
		}
		if value != "" && ((curTokenType == tnum && tokenType != curTokenType) || curTokenType != tnum) {
			tok := token{
				value: value,
				typ:   curTokenType,
			}
			tokens = append(tokens, tok)
			value = ""
			curTokenType = tokenType
			continue
		}
		value += string(r)
		i++
	}
	return tokens
}

func eval(expr []token, idx int) (int, int) {
	res := 0
	op := tplus
	for idx < len(expr) {
		t := expr[idx]
		switch t.typ {
		case tnum:
			n := common.MustAtoi(t.value)
			switch op {
			case tplus:
				res += n
				op = ""
			case tmult:
				res *= n
				op = ""
			}
			idx++
		case tplus:
			op = tplus
			idx++
		case tmult:
			op = tmult
			idx++
		case tlpar:
			n, nextIdx := eval(expr, idx+1)
			switch op {
			case tplus:
				res += n
				op = ""
			case tmult:
				res *= n
				op = ""
			}
			idx = nextIdx
		case trpar:
			return res, idx + 1
		}
	}
	return res, idx
}

func evalOrder(expr []token, idx int, level int) (int, int) {
	refinedTokens := make([]token, 0)

	// Resolve parentheses
	for idx < len(expr) {
		tok := expr[idx]
		switch tok.typ {
		case tnum, tplus, tmult:
			refinedTokens = append(refinedTokens, tok)
			idx++
		case tlpar:
			n, nextIdx := evalOrder(expr, idx+1, level+1)
			refinedTokens = append(refinedTokens, token{value: fmt.Sprint(n), typ: tnum})
			idx = nextIdx
		case trpar:
			idx++
		}
		if tok.typ == trpar {
			break
		}
	}

	// Resolve additions
	i := 0
	for i < len(refinedTokens) {
		if refinedTokens[i].typ == tplus {
			tok := token{
				value: fmt.Sprint(common.MustAtoi(refinedTokens[i-1].value) + common.MustAtoi(refinedTokens[i+1].value)),
				typ:   tnum,
			}
			toks := make([]token, len(refinedTokens)-2)
			copy(toks, refinedTokens[:i-1])
			toks[i-1] = tok
			copy(toks[i:], refinedTokens[i+2:])
			refinedTokens = toks
			continue
		}
		i++
	}

	// Multiplications
	res := 1
	for _, tok := range refinedTokens {
		if tok.typ == tnum {
			res *= common.MustAtoi(tok.value)
		}
	}

	return res, idx
}

func main() {
	common.Run(new(Solution))
}
