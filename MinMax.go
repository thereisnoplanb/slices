package slices

import "github.com/thereisnoplanb/compare"

func MinMax[TSource ~[]TObject, TObject compare.Comparable](source TSource) (min, max TObject, err error) {
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

func MinMaxComparable[TSource ~[]TObject, TObject compare.IComparable[TObject]](source TSource) (min, max TObject, err error) {
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

func MinMaxFunc[TSource ~[]TObject, TObject any](source TSource, compare compare.Comparison[TObject]) (min, max TObject, err error) {
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
