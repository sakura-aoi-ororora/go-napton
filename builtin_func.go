package napton

import "fmt"

const iHandleFuncName = "b/ihandle"

type IHandleFunc struct{}

func (ihf IHandleFunc) Name() string {
	return iHandleFuncName
}

func (ihf IHandleFunc) OnEval(args []ASTNode, top CtxStackNode) (LispValue, error) {
	if len(args) < 2 {
		return ListValue(nil), fmt.Errorf("need 2 args")
	}

	stack := &IHandleStack{previous: top, handler: args[0]}
	return EvalWithStack(args[1], stack)
}

const dismissFuncName = "b/ihandle/dismiss"

type DismissFunc struct{}

func (df DismissFunc) Name() string {
	return dismissFuncName
}

func (df DismissFunc) OnEval(args []ASTNode, top CtxStackNode) (LispValue, error) {
	ident := top.(*IdentStack).ident
	cstack := top
	for {
		switch s := cstack.(type) {
		case *IHandleStack:
			stack := &IdentStack{previous: s.GetPrevious(), ident: ident}
			return EvalWithStack(s.handler, stack)
		case *GlobalHandleStack:
			if fn, ok := s.idents[ident.atom]; ok {
				return BuiltinValue{builtin: fn}, nil
			}
		}
		if cstack.GetPrevious() == nil {
			return ListValue{}, fmt.Errorf("'%s' is not found", ident.atom)
		}
		cstack = cstack.GetPrevious()
	}
}
