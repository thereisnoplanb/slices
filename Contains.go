package slices

import "github.com/thereisnoplanb/compare"

// Determines whether a sequence contains a specified element by using the default equality comparer.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a value.
//
//	value TObject
//
// The value to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains an element that has the specified value; otherwise, false.
func Contains[TSource ~[]TObject, TObject comparable](source TSource, value TObject) (result bool) {
	for i := range source {
		if source[i] == value {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains a specified element by using a specified equality comparer function.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a value.
//
//	value TObject
//
// The value to locate in the sequence.
//
//	equality Equality[TObject]
//
// A function that compares two objects of type TObject for equality.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains an element that has the specified value; otherwise, false.
func ContainsFunc[TSource ~[]TObject, TObject any](source TSource, value TObject, equality compare.Equality[TObject]) (result bool) {
	for i := range source {
		if equality(source[i], value) {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains a specified element by using the IEquatable interface.
//
// # Parameters
//
//	source []TObject
//
// A sequence in which to locate a value.
//
//	value TObject - The value to locate in the sequence.
//
// # Returns
//
//	result bool
//
// True if the source sequence contains an element that has the specified value; otherwise, false.
func ContainsEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, value TObject) (result bool) {
	for i := range source {
		if value.Equal(source[i]) {
			return true
		}
	}
	return false
}

func ContainsBy[TSource ~[]TObject, TObject any, TResult comparable](source TSource, value TObject, valueSelector valueSelector[TObject, TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for i := range source {
		if selectedValue == valueSelector(source[i]) {
			return true
		}
	}
	return false
}

func ContainsByEquatable[TSource ~[]TObject, TObject any, TResult compare.IEquatable[TResult]](source TSource, value TObject, valueSelector valueSelector[TObject, TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for i := range source {
		if selectedValue.Equal(valueSelector(source[i])) {
			return true
		}
	}
	return false
}

func ContainsByFunc[TSource ~[]TObject, TObject any, TResult any](source TSource, value TObject, valueSelector valueSelector[TObject, TResult], equality compare.Equality[TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for i := range source {
		if equality(selectedValue, valueSelector(source[i])) {
			return true
		}
	}
	return false
}
