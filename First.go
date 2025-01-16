package slices

import "errors"

func First[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T, err error) {
	if predicate != nil {
		for _, item := range slice {
			if predicate(item) {
				return item, nil
			}
		}
	}
	return *new(T), errors.New("no element satisfies predicate")
}
