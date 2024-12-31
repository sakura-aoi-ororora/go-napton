package napton

// -- make Runtime state --
type RuntimeBuilder struct {
	baseStack CtxStackNode
	builtins  []BuiltinFunc
}

func NewRuntimeBuilder() *RuntimeBuilder {
	return &RuntimeBuilder{
		baseStack: nil,
		builtins:  make([]BuiltinFunc, 0),
	}
}

func (rb *RuntimeBuilder) BaseStack(stack CtxStackNode) *RuntimeBuilder {
	rb.baseStack = stack
	return rb
}

func (rb *RuntimeBuilder) Builtin(builtin BuiltinFunc) *RuntimeBuilder {
	rb.builtins = append(rb.builtins, builtin)
	return rb
}

func (rb *RuntimeBuilder) Make() *Runtime {
	builtinMap := make(map[string]BuiltinFunc)
	for _, v := range rb.builtins {
		builtinMap[v.Name()] = v
	}

	gh := &GlobalHandleStack{previous: rb.baseStack, idents: builtinMap}
	return &Runtime{
		runtimeStack: gh,
	}
}

type Runtime struct {
	runtimeStack CtxStackNode
}

func (r *Runtime) Run(code string) (LispValue, error) {
	node, err := Parse(code)
	if err != nil {
		return nil, err
	}

	return EvalWithStack(node, r.runtimeStack)
}
