package slices

func FirstOrFallback[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T], fallback T) (result T) {
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				return item
			}
		}
	}
	return fallback
}
