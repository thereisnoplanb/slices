package slices

// Returns the last element in a sequence that satisfies a specified condition.
//
// # Parameters
//
//	source []TObject
//
// A slice to return an element from.
//
//	predicate predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result TObject
//
// The last element in the sequence that passes the test in the specified predicate function or
// the last element in the sequence when predicate is ommited.
//
//	err error
//
// ErrNoElementSatisfiesTheConditionInPredicate - when the sequence contains no element that satisfies the condition in predicate.
// ErrSourceSequenceIsEmpty - when the sequence contains no elements.
func Last[TSource ~[]TObject, TObject any](source TSource, predicate ...predicate[TObject]) (result TObject, err error) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for i := len(source) - 1; i >= 0; i-- {
			if Predicate(source[i]) {
				return source[i], nil
			}
		}
		return *new(TObject), ErrNoElementSatisfiesTheConditionInPredicate
	}
	if len(source) > 0 {
		return source[len(source)-1], nil
	}
	return *new(TObject), ErrSourceSequenceIsEmpty
}
