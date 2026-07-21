package gubrak

import (
	"reflect"
)

func (g *Chainable) Chunk(size int) IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Compact() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Concat(sliceToConcat any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) ConcatMany(slicesToConcat ...any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _concat(err *error, data any, slicesToConcat ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Contains(search any, args ...int) IChainableBoolResult {
	_ = "STUB: not implemented"
	return *new(IChainableBoolResult)
}

func _containsSlice(err *error, dataValue reflect.Value, dataValueLen int, search any, startIndex int) bool {
	_ = "STUB: not implemented"
	return false
}

func _containsCollection(err *error, dataValue reflect.Value, search any, startIndex int) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Chainable) Count() IChainableNumberResult {
	_ = "STUB: not implemented"
	return *new(IChainableNumberResult)
}

func (g *Chainable) CountBy(iteratee any) IChainableNumberResult {
	_ = "STUB: not implemented"
	return *new(IChainableNumberResult)
}

func _count(err *error, data, predicate any) int { _ = "STUB: not implemented"; return 0 }

func _countSlice(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any) int {
	_ = "STUB: not implemented"
	return 0
}

func _countCollection(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any) int {
	_ = "STUB: not implemented"
	return 0
}

func (g *Chainable) Difference(dataToCompare any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) DifferenceMany(datasToCompare ...any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _difference(err *error, data any, dataToCompare ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Drop(size int) IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) DropRight(size int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Each(iteratee any) IChainableNoReturnValueResult {
	_ = "STUB: not implemented"
	return *new(IChainableNoReturnValueResult)
}

func (g *Chainable) EachRight(iteratee any) IChainableNoReturnValueResult {
	_ = "STUB: not implemented"
	return *new(IChainableNoReturnValueResult)
}

func _each(err *error, data, iteratee any, isForward bool) { _ = "STUB: not implemented"; return }

func _eachSlice(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any, isLoopIncremental bool) {
	_ = "STUB: not implemented"
	return
}

func _eachCollection(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any, isLoopIncremental bool) {
	_ = "STUB: not implemented"
	return
}

func (g *Chainable) Exclude(itemToExclude any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) ExcludeMany(itemsToExclude ...any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _exclude(err *error, data any, items ...any) any { _ = "STUB: not implemented"; return *new(any) }

func (g *Chainable) ExcludeAt(indexOfItemToExclude int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) ExcludeAtMany(indexesOfItemToExclude ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _excludeAt(err *error, data any, indexes ...int) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Fill(value any, args ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Filter(predicate any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _filterSlice(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func _filterCollection(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Find(predicate any, args ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) FindIndex(predicate any, args ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) FindLast(predicate any, args ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) FindLastIndex(predicate any, args ...int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) First() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) FromPairs() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) GroupBy(predicate any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) IndexOf(search any, args ...int) IChainableNumberResult {
	_ = "STUB: not implemented"
	return *new(IChainableNumberResult)
}

func (g *Chainable) Initial() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Intersection(dataIntersect any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) IntersectionMany(dataToIntersects ...any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _intersection(err *error, data any, dataIntersects ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Join(separator string) IChainableStringResult {
	_ = "STUB: not implemented"
	return *new(IChainableStringResult)
}

func (g *Chainable) KeyBy(predicate any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Last() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) LastIndexOf(search any, args ...int) IChainableNumberResult {
	_ = "STUB: not implemented"
	return *new(IChainableNumberResult)
}

func (g *Chainable) Map(callback any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Nth(index int) IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) OrderBy(predicate any, args ...bool) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _orderBy(err *error, data, callback any, args ...bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Partition(callback any) IChainableTwoReturnValueResult {
	_ = "STUB: not implemented"
	return *new(IChainableTwoReturnValueResult)
}

func (g *Chainable) Reduce(iteratee, initial any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _reduceCollection(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback, initial any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func _reduceSlice(err *error, dataValue reflect.Value, dataValueType reflect.Type, dataValueKind reflect.Kind, dataValueLen int, callback, initial any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (g *Chainable) Reject(predicate any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Reverse() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Sample() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) SampleSize(take int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Shuffle() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Size() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Tail() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) Take(size int) IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) TakeRight(size int) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func (g *Chainable) Uniq() IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) UnionMany(sliceToUnion ...any) IChainable {
	_ = "STUB: not implemented"
	return *new(IChainable)
}

func _union(err *error, data any, slices ...any) any { _ = "STUB: not implemented"; return *new(any) }
