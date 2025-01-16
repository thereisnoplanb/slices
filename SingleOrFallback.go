package slices

func SingleOrFallback[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T], fallback T) (result T) {
	if predicate != nil {
		found := false
		for _, item := range slice {
			if predicate(item) {
				if found {
					return fallback
				}
				result = item
				found = true
			}
		}
		if found {
			return result
		}
	}
	return fallback
}
