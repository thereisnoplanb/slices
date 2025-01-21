package slices

func LastOrFallback[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T], fallback T) (result T) {
	if predicate != nil {
		for i := len(slice) - 1; i >= 0; i-- {
			if predicate(slice[i]) {
				return slice[i]
			}
		}
	}
	return fallback
}
