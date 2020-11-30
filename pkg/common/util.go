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
