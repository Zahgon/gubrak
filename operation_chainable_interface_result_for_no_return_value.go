package gubrak

type IChainableNoReturnValueResult interface {
	Error() error
	IsError() bool
}

type resultNoReturnValue struct {
	chainable *Chainable
	IChainableNoReturnValueResult
}

type resultEach = resultNoReturnValue

func (g *resultNoReturnValue) Error() error { _ = "STUB: not implemented"; return nil }

func (g *resultNoReturnValue) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *resultNoReturnValue) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultNoReturnValue) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultNoReturnValue) LastOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}
