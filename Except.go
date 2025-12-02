package slices

import "github.com/thereisnoplanb/compare"

func Except[TSource ~[]TObject, TObject comparable](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !Contains(except, source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}

func ExceptEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsEquatable(except, source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}

func ExceptFunc[TSource ~[]TObject, TObject any](source TSource, except TSource, equality compare.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsFunc(except, source[i], equality) {
			result = append(result, source[i])
		}
	}
	return result
}

func ExceptBy[TSource ~[]TObject, TObject any, TResult comparable](source TSource, except TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsBy(except, source[i], valueSelector) {
			result = append(result, source[i])
		}
	}
	return result
}

func ExceptByEquatable[TSource ~[]TObject, TObject any, TResult compare.IEquatable[TResult]](source TSource, except TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsByEquatable(except, source[i], valueSelector) {
			result = append(result, source[i])
		}
	}
	return result
}

func ExceptByFunc[TSource ~[]TObject, TObject any, TResult any](source TSource, except TSource, valueSelctor valueSelector[TObject, TResult], equality compare.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsByFunc(except, source[i], valueSelctor, equality) {
			result = append(result, source[i])
		}
	}
	return result
}
