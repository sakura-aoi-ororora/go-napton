package napton

import "fmt"

func EvalWithStack(ast ASTNode, top CtxStackNode) LispValue {
	switch ast := ast.(type) {
	case ASTNum:
		return NumValue(ast.Value)
	case ASTString:
		return StringValue(ast.Value)
	case ASTAtom:
		return AtomValue{
			atom: ast.Value,
			isDouble: false,
		}
	case ASTIdent:
		cstack := top
		for {
			if cstack.GetPrevious() == nil {
				panic(fmt.Sprintf("'%s' is not found", ast.Value))
			}

			if ih, ok := cstack.(*IdentHandleStack); ok {
				stack := &IdentStack{previous: cstack, ident: AtomValue{atom: ast.Value, isDouble: false}}
				return EvalWithStack(ih.handler, stack)
			}

			cstack = cstack.GetPrevious()
		}
	case ASTList:
		head := EvalWithStack(ast.Value[0], top)
		if evaled, ok := head.(EvalableValue); ok {
			return evaled.Eval(ast.Value[1:len(ast.Value)], top)
		} else {
			panic(fmt.Sprintf("'%s' can't eval", head))
		}
	default:
		panic("eval not implmented")
	}
}
