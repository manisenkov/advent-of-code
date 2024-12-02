package numbers

import (
	"errors"
	"math/big"
)

var (
	ErrNotAnInteger = errors.New("not an integer")
	ErrInvalidType  = errors.New("invalid type")
)

func RatAbs(a *big.Rat) *big.Rat {
	return new(big.Rat).Abs(a)
}

func RatAdd(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Add(a, b)
}

func RatCopy(a *big.Rat) *big.Rat {
	return new(big.Rat).Set(a)
}

func RatEqual(a, b *big.Rat) bool {
	return a.Cmp(b) == 0
}

func RatFromInt[N AnyInt](n N) *big.Rat {
	return big.NewRat(int64(n), 1)
}

func RatIsZero(a *big.Rat) bool {
	return RatEqual(a, RatZero())
}

func RatMul(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Mul(a, b)
}

func RatNeg(a *big.Rat) *big.Rat {
	return new(big.Rat).Neg(a)
}

func RatOne() *big.Rat {
	return big.NewRat(1, 1)
}

func RatSub(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Sub(a, b)
}

func RatQuo(a, b *big.Rat) *big.Rat {
	return new(big.Rat).Quo(a, b)
}

func RatZero() *big.Rat {
	return new(big.Rat)
}

func RatToInt[N AnyInt](a *big.Rat) N {
	if !a.IsInt() {
		panic(ErrNotAnInteger)
	}
	return N(a.Num().Int64())
}

func NumberToRat[N Number](n N) *big.Rat {
	var d any = n
	switch t := d.(type) {
	case float32:
		return (&big.Rat{}).SetFloat64(float64(t))
	case float64:
		return (&big.Rat{}).SetFloat64(t)
	case int:
		return big.NewRat(int64(t), 1)
	case int8:
		return big.NewRat(int64(t), 1)
	case int16:
		return big.NewRat(int64(t), 1)
	case int32:
		return big.NewRat(int64(t), 1)
	case int64:
		return big.NewRat(int64(t), 1)
	case uint:
		return big.NewRat(int64(t), 1)
	case uint8:
		return big.NewRat(int64(t), 1)
	case uint16:
		return big.NewRat(int64(t), 1)
	case uint32:
		return big.NewRat(int64(t), 1)
	case uint64:
		return big.NewRat(int64(t), 1)
	}
	panic(ErrInvalidType)
}

func RatToNumber[N Number](a *big.Rat) N {
	var t any = N(0)

	// Floats
	switch t.(type) {
	case float32:
		f, _ := a.Float32()
		return N(f)
	case float64:
		f, _ := a.Float64()
		return N(f)
	}

	// Integers
	if !a.IsInt() {
		panic(ErrNotAnInteger)
	}
	switch t.(type) {
	case int:
		return N(a.Num().Int64())
	case int8:
		return N(a.Num().Int64())
	case int16:
		return N(a.Num().Int64())
	case int32:
		return N(a.Num().Int64())
	case int64:
		return N(a.Num().Int64())
	case uint:
		return N(a.Num().Int64())
	case uint8:
		return N(a.Num().Int64())
	case uint16:
		return N(a.Num().Int64())
	case uint32:
		return N(a.Num().Int64())
	case uint64:
		return N(a.Num().Int64())
	}
	panic(ErrInvalidType)
}
