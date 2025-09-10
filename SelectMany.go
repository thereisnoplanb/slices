package slices

import "github.com/thereisnoplanb/delegate"

// Projects each element of a sequence and flattens the resulting sequences into one sequence.
//
// # Parameters
//
//	source []TObject
//
// A sequence of values to invoke a one-to-many transform function on.
//
//	valueSelector ValueSelector[TObject, []TResult]
//
// A one-to-many transform function to apply to each element.
//
// # Returns
//
//	result []TResult
//
// A sequence whose elements are the result of invoking the one-to-many transform function on each element of the input sequence.
func SelectMany[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector delegate.ValueSelector[TObject, []TResult]) (result []TResult) {
	result = make([]TResult, 0)
	for i := range source {
		result = append(result, valueSelector(source[i])...)
	}
	return result
}
