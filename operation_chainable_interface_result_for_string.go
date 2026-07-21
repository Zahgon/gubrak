package gubrak

type IChainableStringResult interface {
	ResultAndError() (string, error)
	Result() string
	Error() error
	IsError() bool
}

type resultString struct {
	chainable *Chainable
	IChainableStringResult
}

type resultJoin = resultString

func (g *resultString) ResultAndError() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (g *resultString) Result() string { _ = "STUB: not implemented"; return "" }

func (g *resultString) Error() error { _ = "STUB: not implemented"; return nil }

func (g *resultString) IsError() bool { _ = "STUB: not implemented"; return false }

func (g *resultString) LastSuccessOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultString) LastErrorOperation() Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (g *resultString) LastOperation() Operation { _ = "STUB: not implemented"; return *new(Operation) }
