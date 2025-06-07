package slices

import "errors"

var ErrSourceSequenceIsEmpty = errors.New("the source sequence is empty")
var ErrNoElementSatisfiesTheConditionInPredicate = errors.New("no element satisfies the condition in predicate")
var ErrMoreThanOneElementSatisfiesTheConditionInPredicate = errors.New("more than one element satisfies the condition in predicate")
var ErrSourceSequenceHasMoreThanOneElement = errors.New("the source sequence has more than one element")
var ErrSizeIsBelowOne = errors.New("size is below 1")
var ErrInvalidCast = errors.New("an element in the sequence cannot be cast to requested type")
