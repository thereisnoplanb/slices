package slices

import (
	"github.com/thereisnoplanb/generic"
)

func Min[TSource ~[]TObject, TObject generic.Comparable](source TSource) (result TObject, err error) {
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

func MinComparable[TSource ~[]TObject, TObject generic.IComparable[TObject]](source TSource) (result TObject, err error) {
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

func MinComparator[TSource ~[]TObject, TObject any](source TSource, compare generic.Comparison[TObject]) (result TObject, err error) {
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

func MinBy[TSource ~[]TObject, TObject any, TResult generic.Comparable](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedValue := valueSelector(result)
	for i := 1; i < len(source); i++ {
		if valueSelector(source[i]) < selectedValue {
			result = source[i]
		}
	}
	return result, nil
}

func MinByComparable[TSource ~[]TObject, TObject any, TResult generic.IComparable[TResult]](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedValue := valueSelector(result)
	for i := 1; i < len(source); i++ {
		if valueSelector(source[i]).Compare(selectedValue) < 0 {
			result = source[i]
		}
	}
	return result, nil
}

func MinByComparator[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult], compare generic.Comparison[TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedValue := valueSelector(result)
	for i := 1; i < len(source); i++ {
		if compare(valueSelector(source[i]), selectedValue) < 0 {
			result = source[i]
		}
	}
	return result, nil
}
