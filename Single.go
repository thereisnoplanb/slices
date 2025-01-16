package slices

import "errors"

func Single[TSlice ~[]T, T any](slice TSlice, predicate Predicate[T]) (result T, err error) {
	if predicate != nil {
		found := false
		for _, item := range slice {
			if predicate(item) {
				if found {
					return *new(T), errors.New("more than one element satisfies predicate")
				}
				result = item
				found = true
			}
		}
		if found {
			return result, nil
		}
	}
	return *new(T), errors.New("no element satisfies predicate")
}
