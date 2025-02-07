package slices

import "github.com/thereisnoplanb/generic"

// Computes the sum of a sequence of T values.
//
// Parameters
//
//	source []TObject - A slice of values to calculate the sum of.
//
// Return Value
//
//	result T - The sum of the values in the sequence.
func Sum[TSource ~[]TObject, TObject generic.Number | generic.String](source TSource) (result TObject) {
	for _, item := range source {
		result += item
	}
	return result
}

// Computes the sum of a sequence of TResult values that are obtained by invoking a transform function on each element T of the input sequence.
//
// Parameters
//
//	source []TObject - A sequence of values to calculate the sum of.
//	valueSelector ValueSelector[T, TResult] - A transform function to apply to each element.
//
// Return Value
//
//	result TResult - The sum of the values in the sequence.
func SumBy[TSource ~[]TObject, TObject any, TResult generic.Number | generic.String](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TResult) {
	for _, item := range source {
		result += valueSelector(item)
	}
	return result
}
