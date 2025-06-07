package slices

import (
	"sort"

	"github.com/thereisnoplanb/compare"
	"github.com/thereisnoplanb/delegate"
)

func OrderDescending[TSource ~[]TObject, TObject compare.Comparable](source TSource) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return result[i] > result[j]
	})
	return result
}

func OrderDescendingComparable[TSource ~[]TObject, TObject compare.IComparable[TObject]](source TSource) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Compare(result[j]) > 0
	})
	return result
}

func OrderDescendingComparator[TSource ~[]TObject, TObject any](source TSource, compare compare.Comparison[TObject]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return compare(result[i], result[j]) > 0
	})
	return result
}

func OrderDescendingBy[TSource ~[]TObject, TObject any, TResult compare.Comparable](source TSource, valueSelector delegate.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return valueSelector(result[i]) > valueSelector(result[j])
	})
	return result
}

func OrdeDescendingByComparable[TSource ~[]TObject, TObject any, TResult compare.IComparable[TResult]](source TSource, valueSelector delegate.ValueSelector[TObject, TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return valueSelector(result[i]).Compare(valueSelector(result[j])) > 0
	})
	return result
}

func OrderDescendingByComparator[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector delegate.ValueSelector[TObject, TResult], compare compare.Comparison[TResult]) (result TSource) {
	result = make(TSource, len(source))
	copy(result, source)
	sort.Slice(result, func(i, j int) bool {
		return compare(valueSelector(result[i]), valueSelector(result[j])) > 0
	})
	return result
}
