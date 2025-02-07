package slices

import "github.com/thereisnoplanb/generic"

func GroupBy[TSource ~[]TObject, TObject any, TKey generic.Equatable](source TSource, keySelector generic.KeySelector[TObject, TKey]) (result map[TKey][]TObject) {
	result = make(map[TKey][]TObject)
	for _, value := range source {
		key := keySelector(value)
		result[key] = append(result[key], value)
	}
	return result
}
