package slices

import (
	"github.com/thereisnoplanb/compare"
)

func Min[TSource ~[]TObject, TObject compare.Comparable](source TSource) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		if source[i] < result {
			result = source[i]
		}
	}
	return result, nil
}

func MinComparable[TSource ~[]TObject, TObject compare.IComparable[TObject]](source TSource) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		if source[i].Compare(result) < 0 {
			result = source[i]
		}
	}
	return result, nil
}

func MinFunc[TSource ~[]TObject, TObject any](source TSource, compare compare.Comparison[TObject]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		if compare(source[i], result) < 0 {
			result = source[i]
		}
	}
	return result, nil
}
