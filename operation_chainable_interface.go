package gubrak

type Operation string

const (
	OperationNone             = ""
	OperationChunk            = "Chunk()"
	OperationCompact          = "Compact()"
	OperationConcatMany       = "ConcatMany()"
	OperationConcat           = "Concat()"
	OperationContains         = "Contains()"
	OperationCountBy          = "CountBy()"
	OperationCount            = "Count()"
	OperationDifferenceMany   = "DifferenceMany()"
	OperationDifference       = "Difference()"
	OperationDrop             = "Drop()"
	OperationDropRight        = "DropRight()"
	OperationEach             = "Each()"
	OperationEachRight        = "EachRight()"
	OperationExclude          = "Exclude()"
	OperationExcludeMany      = "ExcludeMany()"
	OperationExcludeAt        = "ExcludeAt()"
	OperationExcludeAtMany    = "ExcludeAtMany()"
	OperationForEach          = "ForEach()"
	OperationForEachRight     = "ForEachRight()"
	OperationFill             = "Fill()"
	OperationFilter           = "Filter()"
	OperationFind             = "Find()"
	OperationFindIndex        = "FindIndex()"
	OperationFindLast         = "FindLast()"
	OperationFindLastIndex    = "FindLastIndex()"
	OperationFirst            = "First()"
	OperationHead             = "Head()"
	OperationFromPairs        = "FromPairs()"
	OperationGroupBy          = "GroupBy()"
	OperationIndexOf          = "IndexOf()"
	OperationInitial          = "Initial()"
	OperationIntersection     = "Intersection()"
	OperationIntersectionMany = "IntersectionMany()"
	OperationJoin             = "Join()"
	OperationKeyBy            = "KeyBy()"
	OperationLast             = "Last()"
	OperationLastIndexOf      = "LastIndexOf()"
	OperationMap              = "Map()"
	OperationNth              = "Nth()"
	OperationOrderBy          = "OrderBy()"
	OperationPartition        = "Partition()"
	OperationReduce           = "Reduce()"
	OperationReject           = "Reject()"
	OperationReverse          = "Reverse()"
	OperationSample           = "Sample()"
	OperationSampleSize       = "SampleSize()"
	OperationShuffle          = "Shuffle()"
	OperationSize             = "Size()"
	OperationTail             = "Tail()"
	OperationTake             = "Take()"
	OperationTakeRight        = "TakeRight()"
	OperationUniq             = "Uniq()"
	OperationUnionMany        = "UnionMany()"
)

type IChainable interface {
	IChainableOperation

	ResultAndError() (any, error)
	Result() any
	Error() error
	IsError() bool
	LastSuccessOperation() Operation
	LastErrorOperation() Operation
	LastOperation() Operation
}

type IChainableOperation interface {
	Chunk(int) IChainable
	Compact() IChainable
	ConcatMany(...any) IChainable
	Concat(any) IChainable
	CountBy(any) IChainableNumberResult
	Count() IChainableNumberResult
	DifferenceMany(...any) IChainable
	Difference(any) IChainable
	Drop(int) IChainable
	DropRight(int) IChainable
	Each(any) IChainableNoReturnValueResult
	EachRight(any) IChainableNoReturnValueResult
	Exclude(any) IChainable
	ExcludeMany(...any) IChainable
	ExcludeAt(int) IChainable
	ExcludeAtMany(...int) IChainable
	Fill(any, ...int) IChainable
	Filter(any) IChainable
	Find(any, ...int) IChainable
	FindIndex(any, ...int) IChainable
	FindLast(any, ...int) IChainable
	FindLastIndex(any, ...int) IChainable
	First() IChainable
	FromPairs() IChainable
	GroupBy(any) IChainable
	Contains(any, ...int) IChainableBoolResult
	IndexOf(any, ...int) IChainableNumberResult
	Initial() IChainable
	Intersection(any) IChainable
	IntersectionMany(data ...any) IChainable
	Join(string) IChainableStringResult
	KeyBy(any) IChainable
	Last() IChainable
	LastIndexOf(any, ...int) IChainableNumberResult
	Map(any) IChainable
	Nth(int) IChainable
	OrderBy(any, ...bool) IChainable
	Partition(any) IChainableTwoReturnValueResult
	Reduce(any, any) IChainable
	Reject(any) IChainable
	Reverse() IChainable
	Sample() IChainable
	SampleSize(int) IChainable
	Shuffle() IChainable
	Size() IChainable
	Tail() IChainable
	Take(int) IChainable
	TakeRight(int) IChainable
	Uniq() IChainable
	UnionMany(...any) IChainable
}

type Chainable struct {
	data                 any
	lastOperation        Operation
	lastSuccessOperation Operation
	lastErrorOperation   Operation
	lastErrorCaught      error
}

func From(data any) IChainable { _ = "STUB: not implemented"; return *new(IChainable) }

func (g *Chainable) markError(data any, err error) *Chainable {
	_ = "STUB: not implemented"
	return nil
}

func (g *Chainable) markResult(data any) *Chainable { _ = "STUB: not implemented"; return nil }

func (g *Chainable) shouldReturn() bool { _ = "STUB: not implemented"; return false }

func (g *Chainable) ResultAndError() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (g *Chainable) Result() any { _ = "STUB: not implemented"; return *new(any) }

func (g *Chainable) Error() error { _ = "STUB: not implemented"; return nil }

func (g *Chainable) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *Chainable) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *Chainable) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *Chainable) LastOperation() Operation { _ = "STUB: not implemented"; return *new(Operation) }
