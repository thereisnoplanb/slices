package slices

// Splits the elements of a sequence into chunks of size at most size.
//
// Parameters
//
//	source []TObject - A collection whose elements to chunk.
//	size int - The maximum size of each chunk.
//
// Return Value
//
//	result []TObject - A collection that contains the elements the input sequence split into chunks of size size.
//	err error - If size is below 1 then ErrSizeIsBelowOne is returned; otherwise, nil.
//
// Each chunk except the last one will be of size size. The last chunk will contain the remaining elements and may be of a smaller size.
func Chunk[TSource ~[]TObject, TObject any](source TSource, size int) (result []TSource, err error) {
	if size < 1 {
		return nil, ErrSizeIsBelowOne
	}
	result = make([]TSource, 0)
	for size < len(source) {
		source, result = source[size:], append(result, source[0:size:size])
	}
	if len(source) > 0 {
		result = append(result, source)
	}
	return result, nil
}
