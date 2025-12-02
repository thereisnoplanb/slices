package slices

import "github.com/thereisnoplanb/compare"

// Determines whether a sequence contains all specified elements by using the default equality comparer.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate values.
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
	for i := range values {
		if !Contains(source, values[i]) {
			return false
		}
	}
	return true
}

// Determines whether a sequence contains all specified elements by using the IEquatable interface.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate values.
//
//	values TObject - The values to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains all elements that have values ​​equal to all of the values ​​in the specified values; false otherwise.
func ContainsAllEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, values ...TObject) (result bool) {
	for i := range values {
		if !ContainsEquatable(source, values[i]) {
			return false
		}
	}
	return true
}

// Determines whether a sequence contains all specified elements by using a specified equality comparer function.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate values.
//
//	values TObject
//
// The values to locate in the sequence.
//
//	equality Equality[TObject]
//
// A function that compares two objects of type TObject for equality.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains all elements that have values ​​equal to all of the values ​​in the specified values; false otherwise.
func ContainsAllFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject], values ...TObject) (result bool) {
	for i := range values {
		if !ContainsFunc(source, values[i], equality) {
			return false
		}
	}
	return true
}
