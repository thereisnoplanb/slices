package slices

import "cmp"

func Order[TSlice ~[]T, T cmp.Ordered](slice TSlice) (result TSlice) {
	result = make(TSlice, len(slice))
	copy(result, slice)
	for i := 0; i < len(result); i++ {
		for j := i; j < len(result); j++ {
			if result[i] < result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

func OrderBy[TSlice ~[]T, T any, TResult cmp.Ordered](slice TSlice, valueSelector ValueSelector[T, TResult]) (result TSlice) {
	result = make(TSlice, len(slice))
	copy(result, slice)
	if valueSelector != nil {
		for i := 0; i < len(result); i++ {
			for j := i; j < len(result); j++ {
				if valueSelector(result[i]) < valueSelector(result[j]) {
					result[i], result[j] = result[j], result[i]
				}
			}
		}
	}
	return result
}
