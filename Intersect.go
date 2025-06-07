package slices

import "github.com/thereisnoplanb/compare"

// Produces the set intersection of two sequences by using the default equality comparer to compare values.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements that also appear in the second sequence will be returned.
//
// 	second []TObject
//
// A sequence whose distinct elements that also appear in the first sequence will be returned.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that form the set intersection of two sequences.
func Intersect[TSource ~[]TObject, TObject comparable](first, second TSource) (result TSource) {
	first = Distinct(first)
	second = Distinct(second)
	result = make(TSource, 0, len(first))
	for _, item := range first {
		if Contains(second, item) {
			result = append(result, item)
			continue
		}
	}
	return result
}

// Produces the set intersection of two sequences by using the default equality comparer to compare values.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements that also appear in the second sequence will be returned.
//
// 	second []TObject
//
// A sequence whose distinct elements that also appear in the first sequence will be returned.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that form the set intersection of two sequences.
func IntersectEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](first, second TSource) (result TSource) {
	first = DistinctEquatable(first)
	second = DistinctEquatable(second)
	result = make(TSource, 0, len(first))
	for _, item := range first {
		if ContainsEquatable(second, item) {
			result = append(result, item)
			continue
		}
	}
	return result
}

// Produces the set intersection of two sequences by using a specified equality comparer to compare values.
//
// # Parameters
//
// 	first []TObject
//
// A sequence whose distinct elements that also appear in the second sequence will be returned.
//
// 	second []TObject
//
// A sequence whose distinct elements that also appear in the first sequence will be returned.
//
//	equality Equality[TObject]
//
// The Equality[TObject] function to compare values for equality.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that form the set intersection of two sequences.
func IntersectFunc[TSource ~[]TObject, TObject any](first, second TSource, equality compare.Equality[TObject]) (result TSource) {
	first = DistinctFunc(first, equality)
	second = DistinctFunc(second, equality)
	result = make(TSource, 0, len(first))
	for _, item := range first {
		if ContainsFunc(second, item, equality) {
			result = append(result, item)
			continue
		}
	}
	return result
}
