package collections

// Map applies the given function to every item in a slice and return slice of results
func Map[T, R any](input []T, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, x := range input {
		result[i] = mapper(x)
	}
	return result
}

// Return true if all items in a slice satisfy the given predicate
func All[T any](input []T, predicate func(T) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}
	return true
}
