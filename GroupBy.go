package slices

import "github.com/thereisnoplanb/delegate"

func GroupBy[TSource ~[]TObject, TObject any, TKey comparable](source TSource, keySelector delegate.KeySelector[TObject, TKey]) (result map[TKey][]TObject) {
	result = make(map[TKey][]TObject)
	for _, value := range source {
		key := keySelector(value)
		result[key] = append(result[key], value)
	}
	return result
}
