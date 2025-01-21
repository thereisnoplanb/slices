package slices

func Count[TSlice ~[]T, T any](slice TSlice) (count int) {
	return len(slice)
}

func CountBy[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (count int) {
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				count++
			}
		}
	}
	return count
}
