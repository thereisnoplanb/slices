package slices

func LastOrDefault[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T) {
	if predicate != nil {
		for i := len(slice) - 1; i >= 0; i-- {
			if predicate(slice[i]) {
				return slice[i]
			}
		}
	}
	return *new(T)
}
