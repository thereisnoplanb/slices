package slices

import "github.com/thereisnoplanb/generic"

// Return Value the first element of a sequence, or a specified default value if the sequence contains no elements.
//
// Parameters
//	source []TObject - Slice to return an element from.
//	fallback TObject - The default value to return if the sequence is empty.
//
// Return Value
//	result TObject - fallback if source is empty; otherwise, the first element in source.
func FirstOrFallback[TSource ~[]TObject, TObject any](source TSource, fallback TObject) (result TObject) {
	if len(source) > 0 {
		return source[0]
	}
	return fallback
}

// Return Value the first element of the sequence that satisfies a condition, or a specified default value if no such element is found.
//
// Parameters
//	source TObject - Slice to return an element from.
//	predicate Predicate[TObject] - A function to test each element for a condition.
//	fallback TObject - The default value to return if the sequence is empty.
//
// Return Value
//	result T - fallback if source is empty or if no element passes the test specified by predicate; otherwise, the first element in source that passes the test specified by predicate.
func FirstByOrFallback[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject], fallback TObject) (result TObject) {
	for _, item := range source {
		if predicate(item) {
			return item
		}
	}
	return fallback
}
