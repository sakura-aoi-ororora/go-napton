package napton

import "fmt"

func EvalWithStack(ast ASTNode, top CtxStackNode) (LispValue, error) {
	switch ast := ast.(type) {
	case ASTNum:
		return NumValue(ast.Value), nil
	case ASTString:
		return StringValue(ast.Value), nil
	case ASTAtom:
		return AtomValue{
			atom:     ast.Value,
			isDouble: false,
		}, nil
	case ASTIdent:
		cstack := top
		for {
			if gh, ok := cstack.(*GlobalHandleStack); ok {
				if val, ok := gh.idents[ast.Value]; ok {
					return BuiltinValue{builtin: val}, nil
				}
			}

			if ih, ok := cstack.(*IHandleStack); ok {
				stack := &IdentStack{previous: ih.GetPrevious(), ident: AtomValue{atom: ast.Value, isDouble: false}}
				return EvalWithStack(ih.handler, stack)
			}

			if cstack.GetPrevious() == nil {
				return ListValue(nil), fmt.Errorf("'%s' is not found", ast.Value)
			}

			cstack = cstack.GetPrevious()
		}
	case ASTList:
		head, err := EvalWithStack(ast.Value[0], top)
		if err != nil {
			return ListValue(nil), err
		}

		if evaled, ok := head.(EvalableValue); ok {
			return evaled.Eval(ast.Value[1:len(ast.Value)], top)
		} else {
			panic(fmt.Sprintf("'%s' can't eval", head))
		}
	default:
		panic("eval not implmented")
	}
}
