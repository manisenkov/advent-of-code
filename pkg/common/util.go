package common

import "strconv"

// MustAtoi converts string to integer and panics in case of error
func MustAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

// MustParseInt interprets a string s in the given base (0, 2 to 36) and bit
// size (0 to 64) and returns the corresponding value.
// It panics in case of error
func MustParseInt(s string, base int, bitSize int) int64 {
	res, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return res
}
