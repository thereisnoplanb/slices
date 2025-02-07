package slices

import "github.com/thereisnoplanb/generic"

func Single[TSource ~[]TObject, TObject any](source TSource) (result TObject, err error) {
	if len(source) == 1 {
		return source[0], nil
	}
	if len(source) > 0 {
		return *new(TObject), ErrSourceSequenceHasMoreThanOneElement
	}
	return *new(TObject), ErrSourceSequenceIsEmpty
}

func SingleBy[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject, err error) {
	found := false
	for _, item := range source {
		if predicate(item) {
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
