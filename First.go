package slices

import "github.com/thereisnoplanb/delegate"

// Returns the first element in a sequence that satisfies a specified condition.
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
// the first element in the sequence when predicate is ommited.
//
//	err error
//
// ErrNoElementSatisfiesTheConditionInPredicate - when the sequence contains no element that satisfies the condition in predicate.
// ErrSourceSequenceIsEmpty - when the sequence contains no elements.
func First[TSource ~[]TObject, TObject any](source TSource, predicate ...delegate.Predicate[TObject]) (result TObject, err error) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			Predicate := predicate[0]
			for i, item := range source {
				if Predicate(item) {
					return source[i], nil
				}
			}
			return *(new(TObject)), ErrNoElementSatisfiesTheConditionInPredicate
		}
		return source[0], nil
	}
	return *new(TObject), ErrSourceSequenceIsEmpty
}
