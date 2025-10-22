package slices

// Returns a specified number of contiguous elements from the start of a sequence.
//
// # Parameters
//
// 	source []TObject
//
// A sequence to return elements from.
//
//	count int
//
// The number of elements to return.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the specified number of elements from the start of the input sequence.
func Take[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	count = min(len(source), max(count, 0))
	result = make(TSource, count)
	_ = copy(result, source)
	return result
}

// Returns a new sequence that contains the last [count] elements from a source sequence.
//
// # Parameters
//
// 	source []TObject
//
// A sequence to return elements from.
//
//	count int
//
// The number of elements to take from the end of the input sequence.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that occur before the last [count] elements in the input sequence.
func TakeLast[TSource ~[]TObject, TObject any](source TSource, count int) (result TSource) {
	count = min(len(source), max(count, 0))
	result = make(TSource, count)
	offset := len(source) - count
	copy(result, source[offset:])
	return result
}

// Returns elements from a sequence as long as a specified condition is true.
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
// A sequence that contains the elements from the input sequence that occur before the element at which the test no longer passes.
func TakeWhile[TSource ~[]TObject, TObject any](source TSource, predicate predicate[TObject]) (result TSource) {
	result = make(TSource, 0, len(source))
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
