package main

import (
	"regexp"
	"sort"
	"time"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

const timeLayout = "2006-01-02 15:04"

var logEntryRegex = regexp.MustCompile(`^\[(\d{4}\-\d{2}\-\d{2} \d{2}:\d{2})\] (Guard #(\d+) )?(begins shift|falls asleep|wakes up)$`)

type entry struct {
	timestamp time.Time
	event     string
	guardID   string
}

// Solution contains solution for day 4
type Solution struct {
	// Guard ID -> minute -> sleep duration
	guardSleepMinutes map[string]map[int]int

	// Guard ID -> total sleep duration
	guardTotalSleepDurations map[string]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	// Read input
	entries := make([]entry, len(input))
	for i, inp := range input {
		m := logEntryRegex.FindAllStringSubmatch(inp, -1)
		timestamp, _ := time.Parse(timeLayout, m[0][1])
		entries[i] = entry{
			timestamp: timestamp,
			event:     m[0][4],
			guardID:   m[0][3],
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		e1 := entries[i]
		e2 := entries[j]
		return e1.timestamp.Before(e2.timestamp)
	})

	// Calculate minutes
	sol.guardSleepMinutes = map[string]map[int]int{}
	sol.guardTotalSleepDurations = map[string]int{}
	currentGuardID := ""
	var beginSleepTime time.Time
	for _, entry := range entries {
		switch entry.event {
		case "begins shift":
			currentGuardID = entry.guardID
		case "falls asleep":
			beginSleepTime = entry.timestamp
		case "wakes up":
			for cur := beginSleepTime; cur.Before(entry.timestamp); cur = cur.Add(time.Minute) {
				sol.guardTotalSleepDurations[currentGuardID]++
				minutes, ok := sol.guardSleepMinutes[currentGuardID]
				if !ok {
					minutes = map[int]int{}
				}
				minutes[cur.Minute()]++
				sol.guardSleepMinutes[currentGuardID] = minutes
			}
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	maxTotalSleepDuration := 0
	sleepiestGuardID := ""
	for guardID, dur := range sol.guardTotalSleepDurations {
		if dur > maxTotalSleepDuration {
			maxTotalSleepDuration = dur
			sleepiestGuardID = guardID
		}
	}

	maxSleepMinuteDuration := 0
	maxSleepMinute := -1
	for minute, dur := range sol.guardSleepMinutes[sleepiestGuardID] {
		if dur > maxSleepMinuteDuration {
			maxSleepMinuteDuration = dur
			maxSleepMinute = minute
		}
	}

	return common.MustAtoi(sleepiestGuardID) * maxSleepMinute
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	maxTotalSleepMinuteDuration := 0
	maxSleepMinute := -1
	sleepiestGuardID := ""
	for guardID, sleepMinutes := range sol.guardSleepMinutes {
		for minute, dur := range sleepMinutes {
			if dur > maxTotalSleepMinuteDuration {
				maxTotalSleepMinuteDuration = dur
				sleepiestGuardID = guardID
				maxSleepMinute = minute
			}
		}
	}
	return common.MustAtoi(sleepiestGuardID) * maxSleepMinute
}

func main() {
	common.Run(new(Solution))
}
