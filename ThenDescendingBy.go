package slices

import (
	"sort"

	"github.com/thereisnoplanb/compare"
)

func ThenDescendingBy[TObject any, TResult compare.Comparable](valueSelector valueSelector[TObject, TResult]) thenBy[TObject] {
	return func(source []TObject) {
		sort.SliceStable(source, func(i, j int) bool {
			return valueSelector(source[i]) > valueSelector(source[j])
		})
	}
}

func ThenDescendingByComparable[TObject any, TResult compare.IComparable[TResult]](valueSelector valueSelector[TObject, TResult]) thenBy[TObject] {
	return func(source []TObject) {
		sort.SliceStable(source, func(i, j int) bool {
			return valueSelector(source[i]).Compare(valueSelector(source[j])) > 0
		})
	}
}
