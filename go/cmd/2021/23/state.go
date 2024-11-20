package main

import (
	"github.com/manisenkov/advent-of-code/pkg/numbers"
)

var deriveCache map[int]map[state][]move = map[int]map[state][]move{
	2: {},
	4: {},
}

var (
	targetRooms = map[byte]int{'A': 0, 'B': 1, 'C': 2, 'D': 3}
	energy      = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
)

type move struct {
	state  state
	energy int
}

type state struct {
	hallway [11]byte
	rooms   [4][4]byte
	depth   int
}

var finalStateP1 = state{
	rooms: [4][4]byte{{'A', 'A'}, {'B', 'B'}, {'C', 'C'}, {'D', 'D'}},
	depth: 2,
}

var finalStateP2 = state{
	rooms: [4][4]byte{
		{'A', 'A', 'A', 'A'},
		{'B', 'B', 'B', 'B'},
		{'C', 'C', 'C', 'C'},
		{'D', 'D', 'D', 'D'},
	},
	depth: 4,
}

func (st state) String() string {
	res := "#############\n"

	sHallway := "#"
	for _, b := range st.hallway {
		if b == 0 {
			sHallway += "."
		} else {
			sHallway += string([]byte{b})
		}
	}
	sHallway += "#\n"
	res += sHallway

	sLines := [4]string{"###", "  #", "  #", "  #"}
	for i := range st.rooms {
		for j := 0; j < 4; j++ {
			if st.rooms[i][j] == 0 {
				sLines[j] += ".#"
			} else {
				sLines[j] += string([]byte{st.rooms[i][j]}) + "#"
			}
		}
	}
	for i := 0; i < st.depth; i++ {
		if i == 0 {
			res += sLines[i] + "##\n"
		} else {
			res += sLines[i] + "\n"
		}
	}

	res += "  #########  \n"
	return res
}

func (st state) isHallwayPathClear(startIndex, endIndex int) bool {
	if endIndex > startIndex {
		for i := startIndex + 1; i <= endIndex; i++ {
			if st.hallway[i] != 0 {
				return false
			}
		}
	} else {
		for i := startIndex - 1; i >= endIndex; i-- {
			if st.hallway[i] != 0 {
				return false
			}
		}
	}
	return true
}

func (st state) isRoomReachable(hallwayIndex int) int {
	h := st.hallway[hallwayIndex]
	roomIndex := targetRooms[h]
	if !st.isHallwayPathClear(hallwayIndex, 2+2*roomIndex) {
		return -1
	}
	for i := st.depth - 1; i >= 0; i-- {
		if st.rooms[roomIndex][i] == 0 {
			return i
		}
		if st.rooms[roomIndex][i] != h {
			break
		}
	}
	return -1
}

func (st state) deriveNext() []move {
	if r, ok := deriveCache[st.depth][st]; ok {
		return r
	}

	res := []move{}

	// Is it possible to move straight to the target room?
	for roomIndex, room := range st.rooms {
		for roomDepth, h := range room {
			if h == 0 {
				continue
			}
			targetRoomIndex := targetRooms[h]
			if targetRoomIndex == roomIndex || !st.isHallwayPathClear(2+2*targetRoomIndex, 2+2*roomIndex) {
				break
			}
			var targetRoomDepth = -1
			for i := st.depth - 1; i >= 0; i-- {
				if st.rooms[targetRoomIndex][i] == 0 {
					targetRoomDepth = i
					break
				}
				if st.rooms[targetRoomIndex][i] != h {
					break
				}
			}
			if targetRoomDepth == -1 {
				break
			}
			nextState := st
			nextState.rooms[roomIndex][roomDepth] = 0
			nextState.rooms[targetRoomIndex][targetRoomDepth] = h
			energySpent := dist(2+2*roomIndex, targetRoomIndex, targetRoomDepth+roomDepth+2) * energy[h]
			res = append(res, move{state: nextState, energy: energySpent})
			break
		}
	}

	// If no direct room-to-room moves found
	if len(res) == 0 {
		// Is it possible to move from the hallway to one of the rooms?
		for i, h := range st.hallway {
			if h == 0 {
				continue
			}
			roomIndex := targetRooms[h]
			roomDepth := st.isRoomReachable(i)
			if roomDepth == -1 {
				continue
			}
			nextState := st
			nextState.hallway[i] = 0
			nextState.rooms[roomIndex][roomDepth] = h
			energySpent := dist(i, roomIndex, roomDepth+1) * energy[h]
			res = append(res, move{state: nextState, energy: energySpent})
		}

		// Is it possible to move from a room to the hallway?
		for roomIndex, room := range st.rooms {
			for roomDepth, h := range room {
				if h == 0 {
					continue
				}
				if roomIndex == targetRooms[h] {
					var needMove bool
					for i := roomDepth + 1; i < st.depth; i++ {
						if room[i] != h {
							needMove = true
						}
					}
					if !needMove {
						break
					}
				}
				for i := range st.hallway {
					if i == 2 || i == 4 || i == 6 || i == 8 || !st.isHallwayPathClear(i, 2+2*roomIndex) {
						continue
					}
					nextState := st
					nextState.hallway[i] = h
					nextState.rooms[roomIndex][roomDepth] = 0
					energySpent := dist(i, roomIndex, roomDepth+1) * energy[h]
					res = append(res, move{state: nextState, energy: energySpent})
				}
				break
			}
		}
	}

	deriveCache[st.depth][st] = res

	return res
}

func dist(hallwayIndex, roomIndex, depth int) int {
	return numbers.Abs(hallwayIndex-(2+2*roomIndex)) + depth
}
