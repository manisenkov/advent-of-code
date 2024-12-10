package main

import (
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func checkSum(disk []int) int {
	res := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		res += i * disk[i]
	}
	return res
}

func getFreeBlockPtr(disk []int, size int) int {
	ptr := 0
	for ptr < len(disk) {
		if disk[ptr] == -1 {
			// check if the current free block is big enough
			blockSize := 0
			for i := 0; i < size && (ptr+i) < len(disk); i++ {
				if disk[ptr+i] == -1 {
					blockSize++
				} else {
					break
				}
			}
			if blockSize == size {
				break
			} else {
				ptr += blockSize
			}
		} else {
			ptr++
		}
	}
	return ptr
}

// Solution contains a solution for day 9
type Solution struct {
	fileBlocks []int
	freeBlocks []int
	disk       []int
	diskMap    []int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.fileBlocks = []int{}
	sol.freeBlocks = []int{}
	diskSize := 0
	for i := 0; i < len(input[0]); i += 2 {
		sol.fileBlocks = append(sol.fileBlocks, int(input[0][i]-'0'))
		diskSize += int(input[0][i] - '0')
		if i+1 < len(input[0]) {
			sol.freeBlocks = append(sol.freeBlocks, int(input[0][i+1]-'0'))
			diskSize += int(input[0][i+1] - '0')
		}
	}

	sol.disk = make([]int, diskSize)
	sol.diskMap = make([]int, len(sol.fileBlocks))
	ptr := 0
	for fileID, fileSize := range sol.fileBlocks {
		sol.diskMap[fileID] = ptr
		for i := 0; i < fileSize; i++ {
			sol.disk[ptr] = fileID
			ptr++
		}
		if fileID < len(sol.freeBlocks) {
			freeSize := sol.freeBlocks[fileID]
			for i := 0; i < freeSize; i++ {
				sol.disk[ptr] = -1
				ptr++
			}
		}
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	disk := make([]int, len(sol.disk))
	copy(disk, sol.disk)
	leftPtr := 0
	rightPtr := len(disk) - 1
	for leftPtr < rightPtr {
		if disk[rightPtr] == -1 {
			rightPtr--
			continue
		}
		if disk[leftPtr] != -1 {
			leftPtr++
			continue
		}
		fileID := disk[rightPtr]
		disk[leftPtr] = fileID
		disk[rightPtr] = -1
		leftPtr++
		rightPtr--
	}
	return checkSum(disk)
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	disk := make([]int, len(sol.disk))
	copy(disk, sol.disk)
	for fileID := len(sol.fileBlocks) - 1; fileID >= 0; fileID-- {
		toPtr := getFreeBlockPtr(disk, sol.fileBlocks[fileID])
		if toPtr >= sol.diskMap[fileID] {
			continue
		}
		for i := 0; i < sol.fileBlocks[fileID]; i++ {
			disk[toPtr+i] = fileID
			disk[sol.diskMap[fileID]+i] = -1
		}
	}
	return checkSum(disk)
}

func main() {
	problem.Solve(new(Solution))
}
