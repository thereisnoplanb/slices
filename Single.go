package slices

import "github.com/thereisnoplanb/delegate"

// Returns the only element of a sequence that satisfies a specified condition.
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
// the single element in the sequence when predicate is ommited.
//
//	err error
//
// ErrNoElementSatisfiesTheConditionInPredicate - when the sequence contains no element that satisfies the condition in predicate.
// ErrMoreThanOneElementSatisfiesTheConditionInPredicate - when the sequence contains more than one element that satisfies the condition in predicate.
// ErrSourceSequenceHasMoreThanOneElement - when the sequence contains more than one element and predicate function is ommited.
// ErrSourceSequenceIsEmpty - when the sequence contains no elements.
func Single[TSource ~[]TObject, TObject any](source TSource, predicate ...delegate.Predicate[TObject]) (result TObject, err error) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			found := false
			Predicate := predicate[0]
			for _, item := range source {
				if Predicate(item) {
					if found {
						return *new(TObject), ErrMoreThanOneElementSatisfiesTheConditionInPredicate
					}
					result = item
					found = true
				}
			}
			if found {
				return result, nil
			}
			return *new(TObject), ErrNoElementSatisfiesTheConditionInPredicate
		}
		if len(source) == 1 {
			return source[0], nil
		}
		return *new(TObject), ErrSourceSequenceHasMoreThanOneElement
	}
	return *new(TObject), ErrSourceSequenceIsEmpty
}
