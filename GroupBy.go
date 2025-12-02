package slices

func GroupBy[TSource ~[]TObject, TObject any, TKey comparable](source TSource, keySelector keySelector[TObject, TKey]) (result map[TKey][]TObject) {
	result = make(map[TKey][]TObject)
	for i := range source {
		key := keySelector(source[i])
		result[key] = append(result[key], source[i])
	}
	return result
}
