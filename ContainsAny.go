package slices

import "github.com/thereisnoplanb/compare"

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
func ContainsAny[TSource ~[]TObject, TObject comparable](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if Contains(source, item) {
			return true
		}
	}
	return false
}

func ContainsAnyFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject], values ...TObject) (result bool) {
	for _, item := range values {
		if ContainsFunc(source, item, equality) {
			return true
		}
	}
	return false
}

func ContainsAnyEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if ContainsEquatable(source, item) {
			return true
		}
	}
	return false
}
