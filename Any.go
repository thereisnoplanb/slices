package slices

import "github.com/thereisnoplanb/generic"

// Determines whether any element of a sequence satisfies a condition.
//
// Parameters
//
//	source []TObject - A slice whose elements to apply the predicate to.
//	predicate Predicate[TObject] - A function to test each element for a condition.
//
// Return Value
//
//	result bool - true if the source sequence is not empty and at least one of its elements passes the test in the specified predicate; otherwise, false.
func Any[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result bool) {
	for _, item := range source {
		if predicate(item) {
			return true
		}
	}
	return false
}
