package slices

import "github.com/thereisnoplanb/delegate"

// Filters a sequence of values based on a predicate.
//
// # Parameters
//
//	source []TObject
//
// A slice to filter.
//
//	predicate Predicate[TObject]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result []TObject
//
// A slice that contains elements from the input sequence that satisfy the condition.
func Where[TSource ~[]TObject, TObject any](source TSource, predicate delegate.Predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
