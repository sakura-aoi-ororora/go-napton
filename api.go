package napton

type BuiltinFunc interface {
	Name() string
	OnEval(args []ASTNode, eval func(ASTNode) LispValue) LispValue
}

type LispValue interface {
	Print()
}

type EvalableValue interface {
	Eval(args []ASTNode, top CtxStackNode) LispValue
}

type CtxStackNode interface {
	SetPrevious(previous CtxStackNode)
	GetPrevious() CtxStackNode
	GetValue() LispValue
}