package slices

func Intersect[TSlice ~[]T, T comparable](slice1 TSlice, slice2 TSlice) (result TSlice) {
	result = make(TSlice, 0, len(slice1))
	for _, item := range slice1 {
		if Contains(slice2, item) {
			result = append(result, item)
		}
	}
	return result
}

func IntersectBy[TSlice ~[]T, T any, TResult comparable](slice1 TSlice, slice2 TSlice, valueSelector ValueSelector[T, TResult]) (result TSlice) {
	result = make(TSlice, 0, len(slice1))
	for _, item := range slice1 {
		if ContainsBy(slice2, valueSelector, item) {
			result = append(result, item)
		}
	}
	return result
}
