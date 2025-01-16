package slices

func SingleOrDefault[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T) {
	if predicate != nil {
		found := false
		for _, item := range slice {
			if predicate(item) {
				if found {
					return *new(T)
				}
				result = item
				found = true
			}
		}
		if found {
			return result
		}
	}
	return *new(T)
}
