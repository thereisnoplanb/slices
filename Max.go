package slices

import "github.com/thereisnoplanb/compare"

func Max[TSource ~[]TObject, TObject compare.Comparable](source TSource) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if value > result {
			result = value
		}
	}
	return result, nil
}

func MaxComparable[TSource ~[]TObject, TObject compare.IComparable[TObject]](source TSource) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if value.Compare(result) > 0 {
			result = value
		}
	}
	return result, nil
}

func MaxFunc[TSource ~[]TObject, TObject any](source TSource, compare compare.Comparison[TObject]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if compare(value, result) > 0 {
			result = value
		}
	}
	return result, nil
}
