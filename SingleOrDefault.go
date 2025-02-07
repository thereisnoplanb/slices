package slices

import "github.com/thereisnoplanb/generic"

func SingleOrDefault[TSource ~[]TObject, TObject any](source TSource) (result TObject) {
	if len(source) == 1 {
		return source[0]
	}
	return *new(TObject)
}

func SingleByOrDefault[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject) {
	found := false
	for _, item := range source {
		if predicate(item) {
			if found {
				return *new(TObject)
			}
			result = item
			found = true
		}
	}
	return result
}
