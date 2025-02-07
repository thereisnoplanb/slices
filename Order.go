package slices

import (
	"sort"

	"github.com/thereisnoplanb/generic"
)

func Order[TSource ~[]TObject, TObject generic.Comparable](source TSource) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func OrderComparable[TSource ~[]TObject, TObject generic.IComparable[TObject]](source TSource) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Compare(result[j]) < 0
	})
	return result
}

func OrderComparator[TSource ~[]TObject, TObject any](source TSource, compare generic.Comparison[TObject]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j]) < 0
	})
	return result
}

func OrderBy[TSource ~[]TObject, TObject any, TResult generic.Comparable](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return valueSelector(result[i]) < valueSelector(result[j])
	})
	return result
}

func OrderByComparable[TSource ~[]TObject, TObject any, TResult generic.IComparable[TResult]](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return valueSelector(result[i]).Compare(valueSelector(result[j])) < 0
	})
	return result
}

func OrderByComparator[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult], compare generic.Comparison[TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return compare(valueSelector(result[i]), valueSelector(result[j])) < 0
	})
	return result
}
