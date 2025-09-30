package slices

import "github.com/thereisnoplanb/compare"

func Distinct[TSource ~[]TObject, TObject comparable](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctFunc[TSource ~[]TObject, TObject any](source TSource, equality compare.Equality[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsFunc(result, item, equality) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](source TSource) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsEquatable(result, item) {
			result = append(result, item)
		}
	}
	return result
}

// func DistinctEqualityComparer[TSource ~[]TObject, TObject any](source TSource, equality generic.Equality[TObject]) (result TSource) {
// 	result = make(TSource, 0, len(source))
// 	for _, item := range source {
// 		if !ContainsEquality(result, item, equality) {
// 			result = append(result, item)
// 		}
// 	}
// 	return result
// }

func DistinctBy[TSource ~[]TObject, TObject any, TResult comparable](source TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsBy(result, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctByEquatable[TSource ~[]TObject, TObject any, TResult compare.IEquatable[TResult]](source TSource, valueSelector valueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByEquatable(result, item, valueSelector) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctByFunc[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector valueSelector[TObject, TResult], equality compare.Equality[TResult]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if !ContainsByFunc(result, item, valueSelector, equality) {
			result = append(result, item)
		}
	}
	return result
}
