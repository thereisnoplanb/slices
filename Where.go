package slices

import "github.com/thereisnoplanb/generic"

func Where[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
