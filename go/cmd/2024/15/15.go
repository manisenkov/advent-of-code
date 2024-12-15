package main

import (
	"maps"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type move int

const (
	up move = iota
	right
	down
	left
)

func (m move) to(pos [2]int) [2]int {
	switch m {
	case up:
		return [2]int{pos[0] - 1, pos[1]}
	case right:
		return [2]int{pos[0], pos[1] + 1}
	case down:
		return [2]int{pos[0] + 1, pos[1]}
	case left:
		return [2]int{pos[0], pos[1] - 1}
	}
	panic("wrong move")
}

func moveBox(boxes collections.Set[[2]int], walls collections.Set[[2]int], pos [2]int, mov move) (collections.Set[[2]int], bool) {
	toPos := mov.to(pos)
	if walls[toPos] {
		return boxes, false
	}
	if boxes[toPos] {
		res, ok := moveBox(boxes, walls, toPos, mov)
		if ok {
			delete(res, pos)
			res[toPos] = true
			return res, true
		}
		return boxes, false
	}
	res := maps.Clone(boxes)
	delete(res, pos)
	res[toPos] = true
	return res, true
}

func moveBigBox(bigBoxes collections.Set[[2]int], bigWalls collections.Set[[2]int], pos [][2]int, mov move) (collections.Set[[2]int], bool) {
	toPos := collections.MapTo(pos, mov.to)
	if collections.Any(toPos, func(p [2]int) bool {
		return bigWalls[p] || bigWalls[right.to(p)]
	}) {
		return bigBoxes, false
	}
	for _, p := range pos {
		delete(bigBoxes, p)
	}
	touchedSet := make(collections.Set[[2]int])
	for _, p := range toPos {
		if bigBoxes[p] {
			touchedSet[p] = true
		}
		if bigBoxes[left.to(p)] {
			touchedSet[left.to(p)] = true
		}
		if bigBoxes[right.to(p)] {
			touchedSet[right.to(p)] = true
		}
	}
	for _, p := range pos {
		bigBoxes[p] = true
	}
	if len(touchedSet) > 0 {
		var touched [][2]int
		for k := range touchedSet {
			touched = append(touched, k)
		}
		res, ok := moveBigBox(bigBoxes, bigWalls, touched, mov)
		if ok {
			for _, p := range pos {
				delete(res, p)
			}
			for _, p := range toPos {
				res[p] = true
			}
			return res, true
		}
		return bigBoxes, false
	}
	res := maps.Clone(bigBoxes)
	for _, p := range pos {
		delete(res, p)
	}
	for _, p := range toPos {
		res[p] = true
	}
	return res, true
}

// Solution contains a solution for day 15
type Solution struct {
	start [2]int
	walls collections.Set[[2]int]
	boxes collections.Set[[2]int]
	moves []move
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.walls = make(collections.Set[[2]int])
	sol.boxes = make(collections.Set[[2]int])
	sol.moves = []move{}
	i := 0
	for ; input[i] != ""; i++ {
		for j, c := range input[i] {
			switch c {
			case '#':
				sol.walls[[2]int{i, j}] = true
			case 'O':
				sol.boxes[[2]int{i, j}] = true
			case '@':
				sol.start = [2]int{i, j}
			}
		}
	}
	i++
	for ; i < len(input); i++ {
		for _, c := range input[i] {
			switch c {
			case '^':
				sol.moves = append(sol.moves, up)
			case '>':
				sol.moves = append(sol.moves, right)
			case 'v':
				sol.moves = append(sol.moves, down)
			case '<':
				sol.moves = append(sol.moves, left)
			}
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	curBoxes := maps.Clone(sol.boxes)
	curPos := sol.start
	for _, mov := range sol.moves {
		toPos := mov.to(curPos)
		if sol.walls[toPos] {
			continue
		} else if curBoxes[toPos] {
			var ok bool
			curBoxes, ok = moveBox(curBoxes, sol.walls, toPos, mov)
			if ok {
				curPos = toPos
			}
		} else {
			curPos = toPos
		}
	}

	res := 0
	for box := range curBoxes {
		res += 100*box[0] + box[1]
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	bigWalls := make(collections.Set[[2]int])
	for w := range sol.walls {
		bigWalls[[2]int{w[0], w[1] * 2}] = true
		bigWalls[[2]int{w[0], w[1]*2 + 1}] = true
	}
	bigBoxes := make(collections.Set[[2]int])
	for b := range sol.boxes {
		bigBoxes[[2]int{b[0], b[1] * 2}] = true
	}
	curPos := [2]int{sol.start[0], sol.start[1] * 2}
	for _, mov := range sol.moves {
		toPos := mov.to(curPos)
		if bigWalls[toPos] {
			continue
		} else if bigBoxes[toPos] || bigBoxes[left.to(toPos)] {
			var boxPos [2]int
			if bigBoxes[toPos] {
				boxPos = toPos
			} else {
				boxPos = left.to(toPos)
			}
			var ok bool
			bigBoxes, ok = moveBigBox(bigBoxes, bigWalls, [][2]int{boxPos}, mov)
			if ok {
				curPos = toPos
			}
		} else {
			curPos = toPos
		}
	}

	res := 0
	for box := range bigBoxes {
		res += 100*box[0] + box[1]
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
