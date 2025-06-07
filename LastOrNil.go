package slices

import "github.com/thereisnoplanb/delegate"

// Returns pointer to the last element in a sequence that satisfies a specified condition or nil value if there is no such element.
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
//	result *TObject
//
// The pointer to the last element in the sequence that passes the test in the specified predicate function or
// the pointer to the laast element in the sequence when predicate is ommited or
// the nil value if there is no such element.
func LastOrNil[TSource ~[]TObject, TObject any](source TSource, predicate ...delegate.Predicate[TObject]) (result *TObject) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for i := len(source) - 1; i >= 0; i-- {
			if Predicate(source[i]) {
				return &source[i]
			}
		}
		return nil
	}
	if len(source) > 0 {
		return &source[len(source)-1]
	}
	return nil
}
