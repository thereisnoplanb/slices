package slices

import "github.com/thereisnoplanb/generic"

func Distinct[TSource ~[]TObject, TObject generic.Equatable](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctEquatable[TSource ~[]TObject, TObject generic.IEquatable[TObject]](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquatable(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctEqualityComparer[TSource ~[]TObject, TObject any](source TSource, equality generic.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquality(result, item, equality) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctBy[TSource ~[]TObject, TObject any, TResult generic.Equatable](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsBy(result, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctByEquatable[TSource ~[]TObject, TObject any, TResult generic.IEquatable[TResult]](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquatable(result, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctByEqualityComparer[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult], equality generic.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquality(result, item, valueSelector, equality) {
			result = append(result, item)
		}
	}
	return result
}
