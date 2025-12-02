package slices

// Converts the elements of a source sequence to the specified type.
//
// # Parameters
//
//	source []TObject
//
// The source sequence that contains the elements TObject to be converted to type TResult.
//
// 	convert convert[TObject, TResult]
//
// A function that converts element of type TObject to another element of type TResult.
//
// # Returns
//
//	result []TResult
//
// A sequence that contains each element of the source sequence converted to the specified type.
func Convert[TSource ~[]TObject, TObject any, TResult any](source TSource, convert convert[TObject, TResult]) (result []TResult) {
	result = make([]TResult, len(source))
	for i := range source {
		result[i] = convert(source[i])
	}
	return result
}
