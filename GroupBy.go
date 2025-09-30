package slices

func GroupBy[TSource ~[]TObject, TObject any, TKey comparable](source TSource, keySelector keySelector[TObject, TKey]) (result map[TKey][]TObject) {
	result = make(map[TKey][]TObject)
	for _, value := range source {
		key := keySelector(value)
		result[key] = append(result[key], value)
	}
	return result
}
