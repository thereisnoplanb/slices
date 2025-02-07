package slices

import "github.com/thereisnoplanb/generic"

// Determines whether a sequence contains a specified element by using the default equality comparer.
//
// Parameters
//
//	source []TObject - A sequence in which to locate a value.
//	value TObject - The value to locate in the sequence.
//
// Return Value
//
//	result bool - true if the source sequence contains an element that has the specified value; otherwise, false.
func Contains[TSource ~[]TObject, TObject generic.Equatable](source TSource, value TObject) (result bool) {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains a specified element by using the IEquatable interface.
//
// Parameters
//
//	source []TObject - A sequence in which to locate a value.
//	value T - The value to locate in the sequence.
//
// Return Value
//
//	result bool - true if the source sequence contains an element that has the specified value; otherwise, false.
func ContainsEquatable[TSource ~[]TObject, TObject generic.IEquatable[TObject]](source TSource, value TObject) (result bool) {
	for _, item := range source {
		if value.Equal(item) {
			return true
		}
	}
	return false
}

// Determines whether a sequence contains a specified element by using a specified equality comparer function.
//
// Parameters
//
//	source []TObject - A sequence in which to locate a value.
//	value T - The value to locate in the sequence.
//	equality Equality[T] - A function that compares two objects of type T.
//
// Return Value
//
//	result bool - true if the source sequence contains an element that has the specified value; otherwise, false.
func ContainsEquality[TSource ~[]TObject, TObject any](source TSource, value TObject, equality generic.Equality[TObject]) (result bool) {
	for _, item := range source {
		if equality(value, item) {
			return true
		}
	}
	return false
}

func ContainsBy[TSource ~[]TObject, TObject any, TResult generic.Equatable](source TSource, value TObject, valueSelector generic.ValueSelector[TObject, TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for _, item := range source {
		if selectedValue == valueSelector(item) {
			return true
		}
	}
	return false
}

func ContainsByEquatable[TSource ~[]TObject, TObject any, TResult generic.IEquatable[TResult]](source TSource, value TObject, valueSelector generic.ValueSelector[TObject, TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for _, item := range source {
		if selectedValue.Equal(valueSelector(item)) {
			return true
		}
	}
	return false
}

func ContainsByEquality[TSource ~[]TObject, TObject any, TResult any](source TSource, value TObject, valueSelector generic.ValueSelector[TObject, TResult], equality generic.Equality[TResult]) (result bool) {
	selectedValue := valueSelector(value)
	for _, item := range source {
		if equality(selectedValue, valueSelector(item)) {
			return true
		}
	}
	return false
}
