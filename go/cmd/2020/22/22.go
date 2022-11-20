package main

import (
	"strconv"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 22
type Solution struct {
	startDecks [2][]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.startDecks = [2][]int{
		make([]int, 0),
		make([]int, 0),
	}
	line := 1
	for input[line] != "" {
		sol.startDecks[0] = append(sol.startDecks[0], common.MustAtoi(input[line]))
		line++
	}
	line += 2
	for line < len(input) && input[line] != "" {
		sol.startDecks[1] = append(sol.startDecks[1], common.MustAtoi(input[line]))
		line++
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	ptrs := [2]int{len(sol.startDecks[0]) - 1, len(sol.startDecks[1]) - 1}
	decks := play(sol.startDecks, ptrs, false, 1)
	winner := whoWin(decks)
	res := 0
	for i := 0; i < len(decks[winner]); i++ {
		res += decks[winner][i] * (len(decks[winner]) - i)
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	ptrs := [2]int{len(sol.startDecks[0]) - 1, len(sol.startDecks[1]) - 1}
	decks := play(sol.startDecks, ptrs, true, 1)
	winner := whoWin(decks)
	res := 0
	for i := 0; i < len(decks[winner]); i++ {
		res += decks[winner][i] * (len(decks[winner]) - i)
	}
	return res
}

func play(startDecks [2][]int, ptrs [2]int, isRecursive bool, depth int) [2][]int {
	gameMem := map[string]bool{}
	decks := [2][]int{
		make([]int, len(startDecks[0])+len(startDecks[1])),
		make([]int, len(startDecks[0])+len(startDecks[1])),
	}
	copy(decks[0], startDecks[0])
	copy(decks[1], startDecks[1])

	for decks[0][0] != 0 && decks[1][0] != 0 {
		key := gameKey(decks, ptrs)
		if gameMem[key] {
			// Game was played before, player 1 win
			return [2][]int{{1}, {0}}
		}
		gameMem[key] = true

		card1 := decks[0][0]
		card2 := decks[1][0]
		var isPlayer1Win bool
		copy(decks[0], decks[0][1:])
		copy(decks[1], decks[1][1:])
		if isRecursive && ptrs[0] >= card1 && ptrs[1] >= card2 {
			recDecks := [2][]int{
				make([]int, card1),
				make([]int, card2),
			}
			copy(recDecks[0], decks[0][:card1])
			copy(recDecks[1], decks[1][:card2])
			recDecks = play(recDecks, [2]int{card1 - 1, card2 - 1}, true, depth+1)
			isPlayer1Win = whoWin(recDecks) == 0
		} else {
			isPlayer1Win = card1 > card2
		}
		if isPlayer1Win {
			// Player 1 win
			decks[0][ptrs[0]] = card1
			decks[0][ptrs[0]+1] = card2
			ptrs[0]++
			ptrs[1]--
		} else {
			// Player 2 win
			decks[1][ptrs[1]] = card2
			decks[1][ptrs[1]+1] = card1
			ptrs[0]--
			ptrs[1]++
		}
	}
	return decks
}

func whoWin(decks [2][]int) int {
	if decks[0][0] != 0 {
		return 0
	}
	return 1
}

func gameKey(decks [2][]int, ptrs [2]int) string {
	k1Parts := make([]string, ptrs[0]+1)
	k2Parts := make([]string, ptrs[1]+1)
	for i, n1 := range decks[0][:ptrs[0]+1] {
		k1Parts[i] = strconv.Itoa(n1)
	}
	for i, n2 := range decks[1][:ptrs[1]+1] {
		k2Parts[i] = strconv.Itoa(n2)
	}
	return strings.Join(k1Parts, ",") + ":" + strings.Join(k2Parts, ",")
}

func main() {
	common.Run(new(Solution))
}
