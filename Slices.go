package slices

type valueSelector[TObject any, TResult any] = func(object TObject) TResult
type accumulate[TAccumulator any, TObject any] = func(accumulator TAccumulator, object TObject) TAccumulator
type predicate[TObject any] = func(object TObject) bool
type convert[TObject any, TResult any] = func(object TObject) TResult
type keySelector[TObject any, TKey comparable] = func(object TObject) TKey
