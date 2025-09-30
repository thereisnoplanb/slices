package slices

// Returns the only element of a sequence that satisfies a specified condition or the default value if there is no such element.
//
// # Parameters
//
//	source []TObject
//
// A slice to return a single element from.
//
//	predicate predicate[TObject] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result TObject
//
// The single element in the sequence that passes the test in the specified predicate function or
// the single element in the sequence when predicate is ommited or
// the default value if there is no such element.
func SingleOrDefault[TSource ~[]TObject, TObject any](source TSource, predicate ...predicate[TObject]) (result TObject) {
	if len(source) > 0 {
		if len(predicate) > 0 {
			found := false
			Predicate := predicate[0]
			for _, item := range source {
				if Predicate(item) {
					if found {
						return *new(TObject)
					}
					result = item
					found = true
				}
			}
			if found {
				return result
			}
			return *new(TObject)
		}
		if len(source) == 1 {
			return source[0]
		}
		return *new(TObject)
	}
	return *new(TObject)
}
