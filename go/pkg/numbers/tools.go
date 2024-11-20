package numbers

import (
	"strconv"
	"unsafe"
)

// Abs returns the absolute value of the given number
func Abs[N Number](a N) N {
	if a < 0 {
		return -a
	}
	return a
}

// ExtGCD finds roots of `a * x + b * y = gcd(a, b)` using the extended Euler algorithm
// for given a, b. It returns the result in as a tuple (gcd(a, b), x, y)
func ExtGCD[N AnyInt](a, b N) (N, N, N) {
	if a == 0 {
		return b, 0, 1
	}
	g, y, x := ExtGCD(b%a, a)
	return g, x - (b/a)*y, y
}

// Max returns the largest number among the given arguments
func Max[N Number](args ...N) N {
	res := args[0]
	for _, n := range args[1:] {
		if n > res {
			res = n
		}
	}
	return res
}

// Min returns the smallest number among the given arguments
func Min[N Number](args ...N) N {
	res := args[0]
	for _, n := range args[1:] {
		if n < res {
			res = n
		}
	}
	return res
}

// ModInv returns the result of the inversion of the given argument a and the modulo m
func ModInv[N AnyInt](a, m N) N {
	g, x, _ := ExtGCD(a, m)
	if g != 1 {
		panic("no solution")
	}
	return x % m
}

// MustAtoi converts the given string to an integer. It panics in case of an error
func MustAtoi[N AnyInt](s string) N {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return N(res)
}

// MustParseInt interprets a string s in the given base (0, 2 to 36) and returns
// the corresponding value. It panics in case of an error
func MustParseInt[N AnyInt](s string, base int) N {
	bitSize := int(unsafe.Sizeof(N(0)) * 8)
	res, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return N(res)
}
