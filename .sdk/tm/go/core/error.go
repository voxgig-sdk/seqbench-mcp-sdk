package core

type SeqbenchMcpError struct {
	IsSeqbenchMcpError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewSeqbenchMcpError(code string, msg string, ctx *Context) *SeqbenchMcpError {
	return &SeqbenchMcpError{
		IsSeqbenchMcpError: true,
		Sdk:              "SeqbenchMcp",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *SeqbenchMcpError) Error() string {
	return e.Msg
}
