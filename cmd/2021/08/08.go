package main

import (
	"sort"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

var numbersMap = map[int]int{
	0b1110111: 0,
	0b0010010: 1,
	0b1011101: 2,
	0b1011011: 3,
	0b0111010: 4,
	0b1101011: 5,
	0b1101111: 6,
	0b1010010: 7,
	0b1111111: 8,
	0b1111011: 9,
}

// Solution contains solution for day 8
type Solution struct {
	signals [][]string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	signals := make([][]string, len(input))
	for i, s := range input {
		parts := strings.Split(s, " | ")
		noteSignals := append(strings.Split(parts[0], " "), strings.Split(parts[1], " ")...)
		for j, p := range noteSignals {
			b := []byte(p)
			sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
			noteSignals[j] = string(b)
		}
		signals[i] = noteSignals
	}
	sol.signals = signals
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := 0
	for _, s := range sol.signals {
		for _, d := range s[len(s)-4:] {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				res++
			}
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	perms := generatePermutations([]byte("abcdefg"), 0)
	patterns := make([]map[byte]int, len(perms))
	for i, p := range perms {
		patterns[i] = generatePattern(p)
	}
	res := 0
	for _, sig := range sol.signals {
		for _, pat := range patterns {
			application := make([]int, len(sig))
			success := true
			for i, s := range sig {
				n := applyPattern(pat, s)
				if n == -1 {
					success = false
					break
				}
				application[i] = n
			}
			if success {
				res +=
					application[len(application)-1] +
						application[len(application)-2]*10 +
						application[len(application)-3]*100 +
						application[len(application)-4]*1000
				break
			}
		}
	}
	return res
}

func applyPattern(pattern map[byte]int, input string) int {
	sum := 0
	for _, b := range []byte(input) {
		sum += pattern[b]
	}
	n, ok := numbersMap[sum]
	if ok {
		return n
	}
	return -1
}

func generatePattern(source []byte) map[byte]int {
	return map[byte]int{
		source[0]: 0b1,
		source[1]: 0b10,
		source[2]: 0b100,
		source[3]: 0b1000,
		source[4]: 0b10000,
		source[5]: 0b100000,
		source[6]: 0b1000000,
	}
}

func generatePermutations(source []byte, i int) [][]byte {
	if i > len(source) {
		return [][]byte{[]byte(string(source))}
	}
	res := generatePermutations(source, i+1)
	for j := i + 1; j < len(source); j++ {
		source[i], source[j] = source[j], source[i]
		res = append(res, generatePermutations(source, i+1)...)
		source[i], source[j] = source[j], source[i]
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
