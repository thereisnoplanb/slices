package slices

// Projects each element of a sequence into a new form.
//
// # Parameters
//
//	source []TObject
//
// A sequence of values to invoke a transform function on.
//
//	valueSelector valueSelector[TObject,TResult]
//
// A transform function to apply to each element.
//
// # Returns
//
//	result []TResult
//
// A sequence whose elements are the result of invoking the transform function on each element of source sequence.
func Select[TSource ~[]TObject, TObject any, TResult any](source TSource, valueSelector valueSelector[TObject, TResult]) (result []TResult) {
	result = make([]TResult, len(source))
	for i := range source {
		result[i] = valueSelector(source[i])
	}
	return result
}
