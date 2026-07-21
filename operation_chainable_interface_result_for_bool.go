package gubrak

type IChainableBoolResult interface {
	ResultAndError() (bool, error)
	Result() bool
	Error() error
	IsError() bool
	LastSuccessOperation() Operation
	LastErrorOperation() Operation
	LastOperation() Operation
}

type resultBool struct {
	chainable *Chainable
	IChainableBoolResult
}

type resultContains = resultBool

func (g *resultBool) ResultAndError() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (g *resultBool) Result() bool { _ = "STUB: not implemented"; return false }

func (g *resultBool) Error() error { _ = "STUB: not implemented"; return nil }

func (g *resultBool) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *resultBool) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultBool) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultBool) LastOperation() Operation { _ = "STUB: not implemented"; return *new(Operation) }
