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
	length := len(source)
	if count >= length {
		return make(TSource, 0)
	}
	result = make(TSource, length-count)
	for i := 0; count < length; count++ {
		result[i] = source[count]
		i++
	}
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
	length := len(source) - count
	if length <= 0 {
		return make(TSource, 0)
	}
	result = make(TSource, len(source))
	for i := 0; i < length; i++ {
		result[i] = source[i]
	}
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
	for _, item := range source {
		if !predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
