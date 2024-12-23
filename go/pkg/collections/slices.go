package collections

import (
	"errors"
	"iter"
)

var ErrEmptySlice = errors.New("slice is empty")

// All returns true if all items in a slice satisfy the given predicate
func All[T any, S ~[]T](input S, predicate func(T) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Any returns true if any of the given items satisfies the given predicate
func Any[T any, S ~[]T](input S, predicate func(T) bool) bool {
	for _, item := range input {
		if predicate(item) {
			return true
		}
	}
	return false
}

// AllCombinations returns all possible combinations of elements in the given input slice
func AllCombinations[T any, S ~[]T](input S) [][]T {
	if len(input) == 0 {
		return [][]T{
			{},
		}
	}
	if len(input) == 1 {
		return [][]T{
			{},
			{input[0]},
		}
	}
	res := [][]T{}
	for _, c := range AllCombinations(input[1:]) {
		res = append(res, c)
		res = append(res, append(append([]T{}, input[0]), c...))
	}
	return res
}

// IsEqualSlices returns true if both given slices have same lengths and contains same elements
func IsEqualSlices[T comparable, S ~[]T](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, x := range s1 {
		y := s2[i]
		if x != y {
			return false
		}
	}
	return true
}

// Filter applies predicate to each element of the given slice and return slice with the elements
// where predicates returned true
func Filter[T any, S ~[]T](input S, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, x := range input {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// IterToSlice collects all of the items of the given iterator sequence to the slice
func IterToSlice[T any](input iter.Seq[T]) []T {
	res := []T{}
	for x := range input {
		res = append(res, x)
	}
	return res
}

// MapTo applies the given function to every item in a slice and return slice of results
func MapTo[T, R any, S ~[]T](input S, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, x := range input {
		result[i] = mapper(x)
	}
	return result
}

// Reduce takes a slice of values of type T and reduces it to a single value by applying
// the given reducer function to each value in sequence to combine it with the current result.
// It starts from the first element of the slice. It panics if the given slice is empty.
func Reduce[T any, S ~[]T](input S, reducer func(T, T) T) T {
	if len(input) == 0 {
		panic(ErrEmptySlice)
	}
	return ReduceWithInit(input[1:], input[0], reducer)
}

// ReduceWithInit takes a slice of values of type T and reduces it to a single value by applying
// the given reducer function to each value in sequence to combine it with the current result.
// It starts from the given init value.
func ReduceWithInit[T any, S ~[]T](input S, init T, reducer func(T, T) T) T {
	res := init
	for _, x := range input {
		res = reducer(res, x)
	}
	return res
}

// Unique returns slice contains unique elements from the given slice
func Unique[T comparable, S ~[]T](input S) S {
	return SetToSlice[T, S](SetFromSlice(input))
}
