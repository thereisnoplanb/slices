package slices

import (
	"github.com/thereisnoplanb/generic"
)

func Max[TSource ~[]TObject, TObject generic.Comparable](source TSource) (result TObject, err error) {
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

func MaxComparable[TSource ~[]TObject, TObject generic.IComparable[TObject]](source TSource) (result TObject, err error) {
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

func MaxComparator[TSource ~[]TObject, TObject any](source TSource, compare generic.Comparison[TObject]) (result TObject, err error) {
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

func MaxBy[TSource ~[]TObject, TObject any, TResult generic.Comparable](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedValue := valueSelector(result)
	for i := 1; i < len(source); i++ {
		value := source[i]
		if valueSelector(value) > selectedValue {
			result = value
		}
	}
	return result, nil
}

func MaxByComparable[TSource ~[]TObject, TObject any, TResult generic.IComparable[TResult]](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedResult := valueSelector(result)
	for i := 1; i < len(source); i++ {
		value := source[i]
		selectedValue := valueSelector(value)
		if selectedValue.Compare(selectedResult) > 0 {
			result = value
			selectedResult = selectedValue
		}
	}
	return result, nil
}

func MaxByComparator[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult], compare generic.Comparison[TResult]) (result TObject, err error) {
	if len(source) == 0 {
		return result, ErrSourceSequenceIsEmpty
	}
	result = source[0]
	selectedResult := valueSelector(result)
	for i := 1; i < len(source); i++ {
		value := source[i]
		selectedValue := valueSelector(value)
		if compare(selectedValue, selectedResult) > 0 {
			result = source[i]
			selectedResult = selectedValue
		}
	}
	return result, nil
}
