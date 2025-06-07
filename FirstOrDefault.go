package slices

import "github.com/thereisnoplanb/delegate"

// Returns the first element in a sequence that satisfies a specified condition or the default value if there is no such element.
//
// # Parameters
//
//	source []TObject
//
// A slice to return an element from.
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
// the default value if there is no such element.
func FirstOrDefault[TSource ~[]TObject, TObject any](source TSource, predicate ...delegate.Predicate[TObject]) (result TObject) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			Predicate := predicate[0]
			for i, item := range source {
				if Predicate(item) {
					return source[i]
				}
			}
			return *(new(TObject))
		}
		return source[0]
	}
	return *(new(TObject))
}
