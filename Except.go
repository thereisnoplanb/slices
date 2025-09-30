package slices

import "github.com/thereisnoplanb/compare"

func Except[TSource ~[]TObject, TObject comparable](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !Contains(except, item) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquatable(except, item) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptFunc[TSource ~[]TObject, TObject any](source TSource, except TSource, equality compare.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsFunc(except, item, equality) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptBy[TSource ~[]TObject, TObject any, TResult comparable](source TSource, except TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsBy(except, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptByEquatable[TSource ~[]TObject, TObject any, TResult compare.IEquatable[TResult]](source TSource, except TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquatable(except, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptByFunc[TSource ~[]TObject, TObject any, TResult any](source TSource, except TSource, valueSelctor valueSelector[TObject, TResult], equality compare.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByFunc(except, item, valueSelctor, equality) {
			result = append(result, item)
		}
	}
	return result
}
