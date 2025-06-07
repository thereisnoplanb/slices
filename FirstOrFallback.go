package slices

import "github.com/thereisnoplanb/delegate"

// Returns the first element in a sequence that satisfies a specified condition or the fallback value value if there is no such element.
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
// The first element in the sequence that passes the test in the specified predicate function or
// the first element in the sequence when predicate is ommited or
// the fallback value if there is no such element.
func FirstOrFallback[TSource ~[]TObject, TObject any](source TSource, fallback TObject, predicate ...delegate.Predicate[TObject]) (result TObject) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			for i, item := range source {
				Predicate := predicate[0]
				if Predicate(item) {
					return source[i]
				}
			}
			return fallback
		}
		return source[0]
	}
	return fallback
}
