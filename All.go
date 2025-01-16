package slices

func All[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result bool) {
	if predicate != nil {
		for _, item := range slice {
			if !predicate(item) {
				return false
			}
		}
	}
	return true
}
