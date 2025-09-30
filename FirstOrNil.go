package slices

// Returns pointer to the first element in a sequence that satisfies a specified condition or nil value if there is no such element.
//
// # Parameters
//
//	source []TObject
//
// A slice to return an element from.
//
//	predicate predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result *TObject
//
// The pointer to the first element in the sequence that passes the test in the specified predicate function or
// the pointer to the first element in the sequence when predicate is ommited or
// the nil value if there is no such element.
func FirstOrNil[TSource ~[]TObject, TObject any](source TSource, predicate ...predicate[TObject]) (result *TObject) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			for i, item := range source {
				Predicate := predicate[0]
				if Predicate(item) {
					return &source[i]
				}
			}
			return nil
		}
		return &source[0]
	}
	return nil
}
