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
// 	next ...[]TObject [OPTIONAL]
//
// A sequence of additional sequences whose distinct elements that also appear in the first and second sequences will be returned.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that form the set intersection of two (or more) sequences.
func Intersect[TSource ~[]TObject, TObject comparable](first, second TSource, next ...TSource) (result TSource) {
	first = Distinct(first)
	second = Distinct(second)
	for i := range next {
		next[i] = Distinct(next[i])
	}
	result = make(TSource, 0, len(first))
	for i := range first {
		if Contains(second, first[i]) && All(next, func(item TSource) bool { return Contains(item, first[i]) }) {
			result = append(result, first[i])
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
// 	next ...[]TObject [OPTIONAL]
//
// A sequence of additional sequences whose distinct elements that also appear in the first and second sequences will be returned.
//
// # Returns
//
//	result []TObject
//
// A sequence that contains the elements that form the set intersection of two (or more) sequences.
func IntersectEquatable[TSource ~[]TObject, TObject compare.IEquatable[TObject]](first, second TSource, next ...TSource) (result TSource) {
	first = DistinctEquatable(first)
	second = DistinctEquatable(second)
	for i := range next {
		next[i] = DistinctEquatable(next[i])
	}
	result = make(TSource, 0, len(first))
	for i := range first {
		if ContainsEquatable(second, first[i]) && All(next, func(item TSource) bool { return ContainsEquatable(item, first[i]) }) {
			result = append(result, first[i])
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
	for i := range first {
		if ContainsFunc(second, first[i], equality) {
			result = append(result, first[i])
			continue
		}
	}
	return result
}
