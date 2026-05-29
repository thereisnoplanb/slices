package slices

// Casts the elements of the source []TSource to the specified type TResult.
//
// # Parameters
//
//	source []TSource
//
// A slice that contains the elements to be cast to type TResult.
//
// # Returns
//
//	result []TResult
//
// A slice that contains each element of the source sequence cast to the specified type.
//
//	err error
//
// ErrInvalidCast when any element in the sequence cannot be casted to type TResult.
func Cast[TSource any, TResult any](source []TSource) (result []TResult, err error) {
	result = make([]TResult, len(source))
	for i := range source {
		value, ok := any(source[i]).(TResult)
		if !ok {
			return nil, ErrInvalidCast
		}
		result[i] = value
	}
	return result, nil
}
