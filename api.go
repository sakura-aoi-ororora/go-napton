package napton

type BuiltinFunc interface {
	Name() string
	OnEval(args []ASTNode, top CtxStackNode) (LispValue, error)
}

type LispValue interface {
	Print()
}

type EvalableValue interface {
	Eval(args []ASTNode, top CtxStackNode) (LispValue, error)
}

type CtxStackNode interface {
	SetPrevious(previous CtxStackNode)
	GetPrevious() CtxStackNode
	GetValue() LispValue
}
