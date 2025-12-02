package slices

import "github.com/thereisnoplanb/compare"

func Distinct[TSource ~[]TObject, TObject comparable](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !Contains(result, source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}

func DistinctFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsFunc(result, source[i], equality) {
			result = append(result, source[i])
		}
	}
	return result
}

func DistinctEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsEquatable(result, source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}

func DistinctBy[TSource ~[]TObject, TObject any, TResult comparable](source TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsBy(result, source[i], valueSelector) {
			result = append(result, source[i])
		}
	}
	return result
}

func DistinctByEquatable[TSource ~[]TObject, TObject any, TResult compare.IEquatable[TResult]](source TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsByEquatable(result, source[i], valueSelector) {
			result = append(result, source[i])
		}
	}
	return result
}

func DistinctByFunc[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector valueSelector[TObject, TResult], equality compare.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !ContainsByFunc(result, source[i], valueSelector, equality) {
			result = append(result, source[i])
		}
	}
	return result
}
