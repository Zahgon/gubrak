package gubrak

type IChainableTwoReturnValueResult interface {
	ResultAndError() (any, any, error)
	ResultTruthy() any
	ResultFalsey() any
	Error() error
	IsError() bool
}

type resultTwoReturnValue struct {
	chainable *Chainable
	IChainableTwoReturnValueResult
}

type resultPartition = resultTwoReturnValue

func (g *resultTwoReturnValue) ResultAndError() (any, any, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(any), nil
}

func (g *resultTwoReturnValue) ResultTruthy() any { _ = "STUB: not implemented"; return *new(any) }

func (g *resultTwoReturnValue) ResultFalsey() any { _ = "STUB: not implemented"; return *new(any) }

func (g *resultTwoReturnValue) Error() error { _ = "STUB: not implemented"; return nil }

func (g *resultTwoReturnValue) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *resultTwoReturnValue) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultTwoReturnValue) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultTwoReturnValue) LastOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}
