package slices

func Any[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result bool) {
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				return true
			}
		}
	}
	return false
}
