package gubrak

type IChainableNumberResult interface {
	ResultAndError() (int, error)
	Result() int
	Error() error
	IsError() bool
	LastSuccessOperation() Operation
	LastErrorOperation() Operation
	LastOperation() Operation
}

type resultNumber struct {
	chainable *Chainable
	IChainableNumberResult
}

type resultCount = resultNumber
type resultLastIndexOf = resultNumber
type resultIndexOf = resultNumber

func (g *resultNumber) ResultAndError() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (g *resultNumber) Result() int { _ = "STUB: not implemented"; return 0 }

func (g *resultNumber) Error() error { _ = "STUB: not implemented"; return nil }

func (g *resultNumber) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *resultNumber) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultNumber) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultNumber) LastOperation() Operation { _ = "STUB: not implemented"; return *new(Operation) }
