package slices

import "errors"

func AverageBy[TSlice ~[]T, T any](slice TSlice, valueSelector ValueSelector[T, float64]) (value float64, err error) {
	if valueSelector == nil {
		return 0, errors.New("value selector is nil")
	}
	count := 0
	for _, item := range slice {
		value += valueSelector(item)
		count++
	}
	if count == 0 {
		return 0, errors.New("sequence contains no elements")
	}
	return value / float64(count), nil
}
