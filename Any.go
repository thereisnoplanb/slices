package slices

// Determines whether any element of a sequence satisfies a condition.
//
// # Parameters
//
//	source []TObject
//
// A slice whose elements to apply the predicate to.
//
// 	predicate predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result bool
//
// True if the source sequence is not empty or nil and at least one of its elements passes the test in the specified predicate; otherwise, false.
//
// # Remarks
//
// If predicate parameter is ommited function returns true if the source sequence contains any elements; otherwise, false.
func Any[TSource ~[]TObject, TObject any](source TSource, predicate ...predicate[TObject]) (result bool) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for _, item := range source {
			if Predicate(item) {
				return true
			}
		}
		return false
	}
	return len(source) > 0
}
