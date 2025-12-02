package slices

// Filters a sequence of values based on a predicate.
//
// # Parameters
//
//	source []TObject
//
// A slice to filter.
//
//	predicate predicate[TObject]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result []TObject
//
// A slice that contains elements from the input sequence that satisfy the condition.
func Where[TSource ~[]TObject, TObject any](source TSource, predicate predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if predicate(source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}
