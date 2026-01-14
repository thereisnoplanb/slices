package slices

// Converts the elements of a source sequence to the specified type.
//
// # Parameters
//
//	source []TObject
//
// The source sequence that contains the elements TObject to be converted to type TResult.
//
// 	convert func(object TObject) (TResult, error)
//
// A function that converts element of type TObject to another element of type TResult, and returns error while converting.
//
// # Returns
//
//	result []TResult
//
// A sequence that contains each element of the source sequence converted to the specified type.
//
//	err error
//
// An error if convertion fails for at lest one object.
func Convert[TSource ~[]TObject, TObject any, TResult any](source TSource, convert func(object TObject) (TResult, error)) (result []TResult, err error) {
	result = make([]TResult, len(source))
	for i := range source {
		result[i], err = convert(source[i])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
