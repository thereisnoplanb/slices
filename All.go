package slices

// Determines whether all elements of a sequence satisfy a condition.
//
// # Parameters
//
//	source []TObject
//
// A slice that contains the elements to apply the predicate to.
//
// 	predicate predicate[TObject]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result bool
//
// True if every element of the source sequence passes the test in the specified predicate, or if the sequence is empty or nil; otherwise, false.
func All[TSource ~[]TObject, TObject any](source TSource, predicate predicate[TObject]) (result bool) {
	for _, item := range source {
		if !predicate(item) {
			return false
		}
	}
	return true
}
