package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

const (
	floor    = '.'
	occupied = '#'
	empty    = 'L'
)

type seatMap struct {
	nCols, nRows int
	seats        []rune
}

func (m seatMap) turnNear() seatMap {
	res := seatMap{
		nCols: m.nCols,
		nRows: m.nRows,
		seats: make([]rune, len(m.seats)),
	}
	for row := 0; row < m.nRows; row++ {
		for col := 0; col < m.nCols; col++ {
			seat := m.get(row, col)
			if seat == floor {
				res.seats[row*m.nCols+col] = floor
				continue
			}

			isOccupied := seat == occupied
			countAdj := 0

			if row > 0 && col > 0 && m.get(row-1, col-1) == occupied {
				countAdj++
			}
			if row > 0 && m.get(row-1, col) == occupied {
				countAdj++
			}
			if row > 0 && col < m.nCols-1 && m.get(row-1, col+1) == occupied {
				countAdj++
			}

			if col > 0 && m.get(row, col-1) == occupied {
				countAdj++
			}
			if col < m.nCols-1 && m.get(row, col+1) == occupied {
				countAdj++
			}

			if row < m.nRows-1 && col > 0 && m.get(row+1, col-1) == occupied {
				countAdj++
			}
			if row < m.nRows-1 && m.get(row+1, col) == occupied {
				countAdj++
			}
			if row < m.nRows-1 && col < m.nCols-1 && m.get(row+1, col+1) == occupied {
				countAdj++
			}

			if !isOccupied && countAdj == 0 {
				res.seats[row*m.nCols+col] = occupied
			} else if isOccupied && countAdj >= 4 {
				res.seats[row*m.nCols+col] = empty
			} else {
				res.seats[row*m.nCols+col] = m.seats[row*m.nCols+col]
			}
		}
	}
	return res
}

func (m seatMap) turnBroad() seatMap {
	res := seatMap{
		nCols: m.nCols,
		nRows: m.nRows,
		seats: make([]rune, len(m.seats)),
	}
	for row := 0; row < m.nRows; row++ {
		for col := 0; col < m.nCols; col++ {
			seat := m.get(row, col)
			if seat == floor {
				res.seats[row*m.nCols+col] = floor
				continue
			}

			isOccupied := seat == occupied
			countAdj := 0

			if m.checkDir(row, col, -1, -1) {
				countAdj++
			}
			if m.checkDir(row, col, -1, 0) {
				countAdj++
			}
			if m.checkDir(row, col, -1, 1) {
				countAdj++
			}
			if m.checkDir(row, col, 0, -1) {
				countAdj++
			}
			if m.checkDir(row, col, 0, 1) {
				countAdj++
			}
			if m.checkDir(row, col, 1, -1) {
				countAdj++
			}
			if m.checkDir(row, col, 1, 0) {
				countAdj++
			}
			if m.checkDir(row, col, 1, 1) {
				countAdj++
			}

			if !isOccupied && countAdj == 0 {
				res.seats[row*m.nCols+col] = occupied
			} else if isOccupied && countAdj >= 5 {
				res.seats[row*m.nCols+col] = empty
			} else {
				res.seats[row*m.nCols+col] = m.seats[row*m.nCols+col]
			}
		}
	}
	return res
}

func (m seatMap) checkDir(row, col, dRow, dCol int) bool {
	row += dRow
	col += dCol
	for row >= 0 && col >= 0 && row < m.nRows && col < m.nCols {
		if m.get(row, col) == occupied {
			return true
		}
		if m.get(row, col) == empty {
			return false
		}
		row += dRow
		col += dCol
	}
	return false
}

func (m seatMap) get(row, col int) rune {
	return m.seats[row*m.nCols+col]
}

func (m seatMap) equal(other seatMap) bool {
	for row := 0; row < m.nRows; row++ {
		for col := 0; col < m.nCols; col++ {
			if m.get(row, col) != other.get(row, col) {
				return false
			}
		}
	}
	return true
}

func (m seatMap) howManyOccupied() int {
	res := 0
	for row := 0; row < m.nRows; row++ {
		for col := 0; col < m.nCols; col++ {
			if m.get(row, col) == occupied {
				res++
			}
		}
	}
	return res
}

func (m seatMap) String() string {
	res := ""
	for row := 0; row < m.nRows; row++ {
		for col := 0; col < m.nCols; col++ {
			res += string(m.seats[row*m.nCols+col])
		}
		res += "\n"
	}
	return res
}

// Solution contains solution for day 11
type Solution struct {
	plan seatMap
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	nCols := len(input[0])
	nRows := len(input)
	sol.plan = seatMap{
		nCols: nCols,
		nRows: nRows,
		seats: make([]rune, len(input)*len(input[0])),
	}
	for row, inp := range input {
		for col, c := range inp {
			sol.plan.seats[row*nCols+col] = c
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	prev := sol.plan
	next := sol.plan.turnNear()
	for !next.equal(prev) {
		prev = next
		next = next.turnNear()
	}
	return next.howManyOccupied()
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	prev := sol.plan
	next := sol.plan.turnBroad()
	for !next.equal(prev) {
		prev = next
		next = next.turnBroad()
	}
	return next.howManyOccupied()
}

func main() {
	common.Run(new(Solution))
}
