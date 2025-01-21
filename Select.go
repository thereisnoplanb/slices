package slices

func Select[TSlice ~[]T, T any, TResult any](slice TSlice, valueSelector ValueSelector[T, TResult]) (result []TResult) {
	result = make([]TResult, len(slice))
	if valueSelector != nil {
		for i := range slice {
			result[i] = valueSelector(slice[i])
		}
	}
	return result
}
