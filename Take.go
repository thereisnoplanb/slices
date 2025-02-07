package slices

import "github.com/thereisnoplanb/generic"

func Take[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	count = min(len(source), count)
	result = make(TSource, count)
	for i := 0; i < count; i++ {
		result[i] = source[i]
	}
	return result
}

func TakeLast[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	length := len(source) - count
	if length <= 0 {
		return make(TSource, 0)
	}
	result = make(TSource, len(source))
	for i := 0; i < length; i++ {
		result[i] = source[i]
	}
	return result
}

func TakeWhile[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
