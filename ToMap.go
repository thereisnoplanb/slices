package slices

// Creates a map from a slice according to specified key selector and element selector functions.
//
// # Parameters
//
//	source []TObject
//
// The source sequence to create a map.
//
// 	keySelector keySelector[TObject, TKey]
//
// A function to extract a key from each element.
//
// 	valueSelector valueSelector[TObject, TResult]
//
// A transform function to produce a result element value from each element.
//
// # Returns
//
//	result map[TKey]TResult
//
// A map[TKey]TResult that contains values of type TResult selected from the input sequence according to specified key.
func ToMap[TSource ~[]TObject, TObject any, TKey comparable, TResult any](source TSource, keySelector keySelector[TObject, TKey], valueSelector valueSelector[TObject, TResult]) (result map[TKey]TResult) {
	result = make(map[TKey]TResult)
	for i := range source {
		result[keySelector(source[i])] = valueSelector(source[i])
	}
	return result
}
