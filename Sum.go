package slices

import "github.com/thereisnoplanb/compare"

// Computes the sum of the values in a sequence.
//
// # Parameters
//
//	source []TObject
//
// A slice of values to calculate the sum of.
//
// # Returns
//
//	result TObject
//
// The sum of the values in the sequence.
func Sum[TSource ~[]TObject, TObject compare.Number | compare.String](source TSource) (result TObject) {
	for _, item := range source {
		result += item
	}
	return result
}

// Computes the sum of the transformed values in a sequence.
//
// # Parameters
//
//	source []TObject
//
// A slice of values to calculate the sum of transformed values.
//
//	valueSelector valueSelector[TObject,TResult]
//
// A transform function to apply to each element.
//
// # Returns
//
//	result TObject
//
// The sum of the transformed values in the sequence.
func SumFunc[TSource ~[]TObject, TObject any, TResult compare.Number | compare.String](source TSource, valueSelector valueSelector[TObject, TResult]) (result TResult) {
	for _, item := range source {
		result += valueSelector(item)
	}
	return result
}
