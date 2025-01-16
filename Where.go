package slices

type Predicate[T any] func(value T) bool

func Where[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result TSlice) {
	result = make(TSlice, 0, len(slice))
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				result = append(result, item)
			}
		}
	}
	return result
}
