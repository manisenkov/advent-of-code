package main

import (
	"strconv"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type item struct {
	value int
	next  *item
}

// Solution contains solution for day 23
type Solution struct {
	initOrder []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.initOrder = make([]int, 9)
	for i, r := range input[0] {
		sol.initOrder[i] = int(r - '0')
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	lastItem := &item{
		value: sol.initOrder[len(sol.initOrder)-1],
	}
	itemMap := map[int]*item{
		lastItem.value: lastItem,
	}

	cur := lastItem
	for i := len(sol.initOrder) - 2; i >= 0; i-- {
		cur = &item{
			value: sol.initOrder[i],
			next:  cur,
		}
		itemMap[sol.initOrder[i]] = cur
	}
	lastItem.next = cur

	for i := 0; i < 100; i++ {
		turn(cur, itemMap)
		cur = cur.next
	}

	res := ""
	cur = itemMap[1].next
	for cur.value != 1 {
		res += strconv.Itoa(cur.value)
		cur = cur.next
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	lastItem := &item{
		value: 1000000,
	}
	itemMap := map[int]*item{
		lastItem.value: lastItem,
	}

	cur := lastItem
	for i := 999999; i >= 10; i-- {
		cur = &item{
			value: i,
			next:  cur,
		}
		itemMap[i] = cur
	}
	for i := len(sol.initOrder) - 1; i >= 0; i-- {
		cur = &item{
			value: sol.initOrder[i],
			next:  cur,
		}
		itemMap[sol.initOrder[i]] = cur
	}
	lastItem.next = cur

	for i := 0; i < 10000000; i++ {
		turn(cur, itemMap)
		cur = cur.next
	}

	return itemMap[1].next.value * itemMap[1].next.next.value
}

func turn(cur *item, itemMap map[int]*item) {
	nextThree := []*item{
		cur.next,
		cur.next.next,
		cur.next.next.next,
	}
	curNext := nextThree[2].next
	var destValue int
	if cur.value == 1 {
		destValue = len(itemMap)
	} else {
		destValue = cur.value - 1
	}
	for nextThree[0].value == destValue || nextThree[1].value == destValue || nextThree[2].value == destValue {
		destValue--
		if destValue == 0 {
			destValue = len(itemMap)
		}
	}
	dest := itemMap[destValue]
	destNext := dest.next

	cur.next = curNext
	dest.next = nextThree[0]
	nextThree[2].next = destNext
}

func main() {
	common.Run(new(Solution))
}
