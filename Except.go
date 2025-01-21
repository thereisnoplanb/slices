package slices

func Except[TSlice ~[]T, T comparable](slice TSlice, except TSlice) (result TSlice) {
	result = make(TSlice, 0, len(slice))
	for _, item := range slice {
		if !Contains(except, item) {
			result = append(result, item)
		}
	}
	return result
}

func ExceptBy[TSlice ~[]T, T any, TResult comparable](slice TSlice, except TSlice, valueSelector ValueSelector[T, TResult]) (result TSlice) {
	result = make(TSlice, 0, len(slice))
	for _, item := range slice {
		if !ContainsBy(except, valueSelector, item) {
			result = append(result, item)
		}
	}
	return result
}
