package slices

import "github.com/thereisnoplanb/delegate"

// Returns the last element in a sequence that satisfies a specified condition or the fallback value value if there is no such element.
//
// # Parameters
//
//	source []TObject
//
// A slice to return an element from.
//
//	fallback TObject
//
// The value that is returned when source sequence is empty or if no element passes the test specified by predicate.
//
//	predicate Predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result TObject
//
// The last element in the sequence that passes the test in the specified predicate function or
// the last element in the sequence when predicate is ommited or
// the fallback value if there is no such element.
func LastOrFallback[TSource ~[]TObject, TObject any](source TSource, fallback TObject, predicate ...delegate.Predicate[TObject]) (result TObject) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for i := len(source) - 1; i >= 0; i-- {
			if Predicate(source[i]) {
				return source[i]
			}
		}
		return fallback
	}
	if len(source) > 0 {
		return source[len(source)-1]
	}
	return fallback
}
