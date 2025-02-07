package slices

import "github.com/thereisnoplanb/generic"

// Return Value the first element of a sequence, or a default value if the sequence contains no elements.
//
// Parameters
//	source TObject - Slice to return an element from.
//
// Return Value
//	result TObject - default value if source is empty; otherwise, the first element in source.
func FirstOrDefault[TSource ~[]TObject, TObject any](source TSource) (result TObject) {
	if len(source) > 0 {
		return source[0]
	}
	return *new(TObject)
}

// Return Value the first element of the sequence that satisfies a condition or a default value if no such element is found.
//
// Parameters
//	source []TObject - Slice to return an element from.
//	predicate Predicate[TObject] - A function to test each element for a condition.
//
// Return Value
//	result TObject - default value if source is empty or if no element passes the test specified by predicate; otherwise, the first element in source that passes the test specified by predicate.
func FirstByOrDefault[TSource ~[]TObject, TObject any](source TSource, predicate generic.Predicate[TObject]) (result TObject) {
	for _, item := range source {
		if predicate(item) {
			return item
		}
	}
	return *new(TObject)
}
