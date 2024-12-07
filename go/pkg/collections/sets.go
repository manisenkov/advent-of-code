package collections

// Set is an unordered collection of unique elements
type Set[T comparable] map[T]bool

// SetFromSlice creates new set from the give input slice
func SetFromSlice[T comparable, S ~[]T](input S) Set[T] {
	res := make(Set[T])
	for _, el := range input {
		res[el] = true
	}
	return res
}

// AppendSliceToSet adds element of the given input slice to the target set
func AppendSliceToSet[T comparable, S ~[]T](input S, target Set[T]) Set[T] {
	for _, el := range input {
		target[el] = true
	}
	return target
}

// MergeSets add values of the given source set into a target set and return target set as a result
func MergeSets[T comparable](source, target Set[T]) Set[T] {
	for el := range source {
		target[el] = true
	}
	return target
}
