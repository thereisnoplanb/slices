package slices

import "errors"

func Last[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T, err error) {
	if predicate != nil {
		for i := len(slice) - 1; i >= 0; i-- {
			if predicate(slice[i]) {
				return slice[i], nil
			}
		}
	}
	return *new(T), errors.New("no element satisfies predicate")
}
