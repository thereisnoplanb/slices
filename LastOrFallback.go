package slices

import "github.com/thereisnoplanb/generic"

func LastOrFallback[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject], fallback TObject) (result TObject) {
	if len(source) > 0 {
		return source[len(source)-1]
	}
	return fallback
}

func LastByOrFallback[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject], fallback TObject) (result TObject) {
	for i := len(source) - 1; i >= 0; i-- {
		if predicate(source[i]) {
			return source[i]
		}
	}
	return fallback
}
