package common

import "strconv"

// AbsInt returns absolute value of the given integer
func AbsInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// ExtGCD implements extended euler algorithm to find roots of a*x+b*y=gcd(a,b).
// It returns a tuple of in form of (gcd(a,b), x, y) for given a, b.
func ExtGCD(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x := ExtGCD(b%a, a)
	return g, x - (b/a)*y, y
}

// Max returns the bigger number of two
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Max returns the smaller number of two
func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// ModInv returns modulo inversion for a modulo m
func ModInv(a, m int64) int64 {
	g, x, _ := ExtGCD(a, m)
	if g != 1 {
		panic("no solution")
	}
	return x % m
}

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
