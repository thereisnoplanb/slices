package slices

import "github.com/thereisnoplanb/generic"

func MinMax[TSource ~[]TObject, TObject generic.Comparable](source TSource) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	max = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if value < min {
			min = value
		} else if value > max {
			max = value
		}
	}
	return min, max, nil
}

func MinMaxComparable[TSource ~[]TObject, TObject generic.IComparable[TObject]](source TSource) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	max = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if value.Compare(min) < 0 {
			min = value
		} else if value.Compare(max) > 0 {
			max = value
		}
	}
	return min, max, nil
}

func MinMaxComparator[TSource ~[]TObject, TObject any](source TSource, compare generic.Comparison[TObject]) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	max = source[0]
	for i := 1; i < len(source); i++ {
		value := source[i]
		if compare(value, min) < 0 {
			min = value
		} else if compare(value, max) > 0 {
			max = value
		}
	}
	return min, max, nil
}

func MinMaxBy[TSource ~[]TObject, TObject any, TResult generic.Comparable](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	minSelectedValue := valueSelector(min)
	max = source[0]
	maxSelectedValue := valueSelector(max)
	for i := 1; i < len(source); i++ {
		value := source[i]
		selectedValue := valueSelector(value)
		if selectedValue < minSelectedValue {
			minSelectedValue = selectedValue
			min = value
		} else if selectedValue > maxSelectedValue {
			maxSelectedValue = selectedValue
			max = value
		}
	}
	return min, max, nil
}

func MinMaxByComparable[TSource ~[]TObject, TObject any, TResult generic.IComparable[TResult]](source TSource, valueSelector generic.ValueSelector[TObject, TResult]) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	minSelectedValue := valueSelector(min)
	max = source[0]
	maxSelectedValue := valueSelector(max)
	for i := 1; i < len(source); i++ {
		value := source[i]
		selectedValue := valueSelector(value)
		if selectedValue.Compare(minSelectedValue) < 0 {
			minSelectedValue = selectedValue
			min = value
		} else if selectedValue.Compare(maxSelectedValue) > 0 {
			maxSelectedValue = selectedValue
			max = value
		}
	}
	return min, max, nil
}

func MinMaxByComparator[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector generic.ValueSelector[TObject, TResult], compare generic.Comparison[TResult]) (min, max TObject, err error) {
	if len(source) == 0 {
		return min, max, ErrSourceSequenceIsEmpty
	}
	min = source[0]
	minSelectedValue := valueSelector(min)
	max = source[0]
	maxSelectedValue := valueSelector(max)
	for i := 1; i < len(source); i++ {
		value := source[i]
		selectedValue := valueSelector(value)
		if compare(selectedValue, minSelectedValue) < 0 {
			minSelectedValue = selectedValue
			min = value
		} else if compare(selectedValue, maxSelectedValue) > 0 {
			maxSelectedValue = selectedValue
			max = value
		}
	}
	return min, max, nil
}
