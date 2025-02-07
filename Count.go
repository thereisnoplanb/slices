package slices

import "github.com/thereisnoplanb/generic"

func Count[TSource ~[]TObject, TObject any](source TSource) (count int) {
	return len(source)
}

func CountBy[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (count int) {
	for _, item := range source {
		if predicate(item) {
			count++
		}
	}
	return count
}
