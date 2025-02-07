package slices

import "github.com/thereisnoplanb/generic"

func Last[TSource ~[]TObject, TObject any](source TSource) (result TObject, err error) {
	if len(source) > 0 {
		return source[len(source)-1], nil
	}
	return *new(TObject), ErrSourceSequenceIsEmpty
}

func LastBy[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject, err error) {
	for i := len(source) - 1; i >= 0; i-- {
		if predicate(source[i]) {
			return source[i], nil
		}
	}
	return *new(TObject), ErrNoElementSatisfiesTheConditionInPredicate
}
