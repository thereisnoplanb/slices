package slices

import "github.com/thereisnoplanb/delegate"

// Returns the pointer to the only element of a sequence that satisfies a specified condition or nil value if there is no such element.
//
// # Parameters
//
//	source []TObject
//
// A slice to return a single element from.
//
//	predicate Predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result TObject
//
// The single element in the sequence that passes the test in the specified predicate function or
// the single element in the sequence when predicate is ommited or
// the nil value if there is no such element.
func SingleOrNil[TSource ~[]TObject, TObject any](source TSource, predicate ...delegate.Predicate[TObject]) (result *TObject) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			found := false
			Predicate := predicate[0]
			for i, item := range source {
				if Predicate(item) {
					if found {
						return nil
					}
					result = &source[i]
					found = true
				}
			}
			if found {
				return result
			}
			return nil
		}
		if len(source) == 1 {
			return &source[0]
		}
		return nil
	}
	return nil
}
