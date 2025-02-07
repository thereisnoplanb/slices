package slices

import "github.com/thereisnoplanb/generic"

func SingleOrFallback[TSource ~[]TObject, TObject any](source TSource, fallback TObject) (result TObject) {
	if len(source) == 1 {
		return source[0]
	}
	return fallback
}

func SingleByOrFallback[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject], fallback TObject) (result TObject) {
	found := false
	for _, item := range source {
		if predicate(item) {
			if found {
				return fallback
			}
			result = item
			found = true
		}
	}
	if found {
		return result
	}
	return fallback
}
