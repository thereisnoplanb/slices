package slices

func Contains[TSlice ~[]T, T comparable](slice TSlice, value T) (result bool) {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func ContainsBy[TSlice ~[]T, T any, TResult comparable](slice TSlice, valueSelector ValueSelector[T, TResult], value T) (result bool) {
	if valueSelector != nil {
		selectedValue := valueSelector(value)
		for _, item := range slice {
			if valueSelector(item) == selectedValue {
				return true
			}
		}
	}
	return false
}
