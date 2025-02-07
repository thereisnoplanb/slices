package slices

import "github.com/thereisnoplanb/generic"

func Except[TSource ~[]TObject, TObject generic.Equatable](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !Contains(except, item) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptEquatable[TSource ~[]TObject, TObject generic.IEquatable[TObject]](source TSource, except TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquatable(except, item) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptEqualityComparer[TSource ~[]TObject, TObject any](source TSource, except TSource, equality generic.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquality(except, item, equality) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptBy[TSource ~[]TObject, TObject any, TResult generic.Equatable](source TSource, except TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsBy(except, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptByEquatable[TSource ~[]TObject, TObject any, TResult generic.IEquatable[TResult]](source TSource, except TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquatable(except, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptByEqualityComparer[TSource ~[]TObject, TObject any, TResult any](source TSource, except TSource, valueSelctor generic.ValueSelector[TObject, TResult], equality generic.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquality(except, item, valueSelctor, equality) {
			result = append(result, item)
		}
	}
	return result
}
