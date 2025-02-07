package slices

import "github.com/thereisnoplanb/generic"

func Select[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result []TResult) {
	result = make([]TResult, len(source))
	for i := range source {
		result[i] = valueSelector(source[i])
	}
	return result
}
