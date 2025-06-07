package slices

import "github.com/thereisnoplanb/compare"

// Determines whether a sequence contains all specified elements by using the default equality comparer.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a values.
//
// 	values TObject
//
// The values to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains all elements that have values ​​equal to all of the values ​​in the specified values; false otherwise.
func ContainsAll[TSource ~[]TObject, TObject comparable](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if !Contains(source, item) {
			return false
		}
	}
	return true
}

func ContainsAllEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, values ...TObject) (result bool) {
	for _, item := range values {
		if !ContainsEquatable(source, item) {
			return false
		}
	}
	return true
}

func ContainsAllFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject], values ...TObject) (result bool) {
	for _, item := range values {
		if !ContainsFunc(source, item, equality) {
			return false
		}
	}
	return true
}
