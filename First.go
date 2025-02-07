package slices

import "github.com/thereisnoplanb/generic"

func First[TSource ~[]TObject, TObject any](source TSource) (result TObject, err error) {
	if len(source) == 0 {
		return *new(TObject), ErrSourceSequenceIsEmpty
	}
	return source[0], nil
}

func FirstBy[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject, err error) {
	for _, item := range source {
		if predicate(item) {
			return item, nil
		}
	}
	return *new(TObject), ErrNoElementSatisfiesTheConditionInPredicate
}
