package slices

func Distinct[TSlice ~[]T, T comparable](slice TSlice) (result TSlice) {
	result = make(TSlice, 0, len(slice))
	for _, item := range slice {
		if !Contains(result, item) {
			result = append(result, item)
		}
	}
	return result
}

func DistinctBy[TSlice ~[]T, T any, TResult comparable](slice TSlice, valueSelector ValueSelector[T, TResult]) (result TSlice) {
	result = make(TSlice, 0, len(slice))
	for _, item := range slice {
		if !ContainsBy(result, valueSelector, item) {
			result = append(result, item)
		}
	}
	return result
}
