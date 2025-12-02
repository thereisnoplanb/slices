package slices

import "github.com/thereisnoplanb/compare"

// Determines whether a sequence contains any specified element by using the default equality comparer.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a values.
//
//	values TObject
//
// The values to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains any element that has value ​​equal to any of the values ​​in the specified values; false otherwise.
func ContainsAny[TSource ~[]TObject, TObject comparable](source TSource, values ...TObject) (result bool) {
	for i := range values {
		if Contains(source, values[i]) {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains any specified element by using a specified equality comparer function.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a values.
//
//	values TObject
//
// The values to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains any element that has value ​​equal to any of the values ​​in the specified values; false otherwise.
func ContainsAnyFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject], values ...TObject) (result bool) {
	for i := range values {
		if ContainsFunc(source, values[i], equality) {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains any specified element by using the IEquatable interface.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a values.
//
//	values TObject
//
// The values to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains any element that has value ​​equal to any of the values ​​in the specified values; false otherwise.
func ContainsAnyEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, values ...TObject) (result bool) {
	for i := range values {
		if ContainsEquatable(source, values[i]) {
			return true
		}
	}
	return false
}
