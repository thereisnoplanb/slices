package slices

// Bypasses a specified number of elements in a sequence and then returns the remaining elements.
//
// # Parameters
//
// 	source []TObject
//
// A sequence to return elements from.
//
//	count int
//
// The number of elements to skip before returning the remaining elements.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that occur after the first [count] elements in the input sequence.
func Skip[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	count = min(len(source), max(count, 0))
	result = make(TSource, len(source)-count)
	copy(result, source[count:])
	return result
}

// Returns a new sequence that contains the elements from source with the last count elements of the source sequence omitted.
//
// # Parameters
//
// 	source []TObject
//
// A sequence to return elements from.
//
//	count int
//
// The number of elements to omit from the end of the input sequence.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that occur before the last [count] elements in the input sequence.
func SkipLast[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	count = min(len(source), max(count, 0))
	result = make(TSource, len(source)-count)
	copy(result, source)
	return result
}

// Bypasses elements in a sequence as long as a specified condition is true and then returns the remaining elements.
//
// # Parameters
//
// 	source []TObject
//
// A sequence to return elements from.
//
//	predicate predicate[TObject]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements from the source sequence starting at the first element in the linear series
// that does not pass the test specified by predicate.
func SkipWhile[TSource ~[]TObject, TObject any](source TSource, predicate predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for i := range source {
		if !predicate(source[i]) {
			result = append(result, source[i])
		}
	}
	return result
}
