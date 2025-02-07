package slices

import "github.com/thereisnoplanb/generic"

// Determines whether a sequence contains any specified element by using the default equality comparer.
//
// Parameters
//
//	source []TObject - A sequence in which to locate a values.
//	values TObject - The values to locate in the sequence.
//
// Return Value
//
//	result bool - true if the source sequence contains any element that has a value equal to any of the specified values; false otherwise.
func ContainsAny[TSource ~[]TObject, TObject generic.Equatable](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if Contains(source, item) {
			return true
		}
	}
	return false
}
