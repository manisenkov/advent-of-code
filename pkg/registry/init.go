package registry

import (
	day2018_01 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day01"
	day2018_02 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day02"
	day2018_03 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day03"

	day2020_01 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day01"
	day2020_02 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day02"
	day2020_03 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day03"
	day2020_04 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day04"
	day2020_05 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day05"
	day2020_06 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day06"
)

func init() {
	Register(2018, 1, new(day2018_01.Solution))
	Register(2018, 2, new(day2018_02.Solution))
	Register(2018, 3, new(day2018_03.Solution))

	Register(2020, 1, new(day2020_01.Solution))
	Register(2020, 2, new(day2020_02.Solution))
	Register(2020, 3, new(day2020_03.Solution))
	Register(2020, 4, new(day2020_04.Solution))
	Register(2020, 5, new(day2020_05.Solution))
	Register(2020, 6, new(day2020_06.Solution))
}
