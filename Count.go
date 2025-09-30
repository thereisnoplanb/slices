package slices

// Returns a number that represents how many elements in the specified sequence satisfy a condition.
//
// # Parameters
//
//	source []TObject
//
// A sequence that contains elements to be tested and counted.
//
//	predicate predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	count int
//
// A number that represents how many elements in the sequence satisfy the condition in the predicate function or
// the number of elements in the input sequence if predicate is ommited.
func Count[TSource ~[]TObject, TObject any](source TSource, predicate ...predicate[TObject]) (count int) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for _, item := range source {
			if Predicate(item) {
				count++
			}
		}
		return count
	}
	return len(source)
}
