package slices

import "github.com/thereisnoplanb/generic"

func LastOrDefault[TSource ~[]TObject, TObject any](source TSource) (result TObject) {
	if len(source) > 0 {
		return source[len(source)-1]
	}
	return *new(TObject)
}

func LastByOrDefault[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject) {
	for i := len(source) - 1; i >= 0; i-- {
		if predicate(source[i]) {
			return source[i]
		}
	}
	return *new(TObject)
}
