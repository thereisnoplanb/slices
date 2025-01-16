package slices

func FirstOrDefault[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T) {
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				return item
			}
		}
	}
	return *new(T)
}
