package napton

import "fmt"

type NumValue float64

func (n NumValue) Print() {
	fmt.Print(n)
}

type StringValue string

func (s StringValue) Print() {
	fmt.Printf("%q", s)
}

type AtomValue struct {
	atom string
	isDouble bool
}

func (a AtomValue) Print() {
	if !a.isDouble {
		// single
		fmt.Printf(":%s", a.atom)
	} else {
		// double
		fmt.Printf(":%s(double)", a.atom)
	}
}

type ListValue []LispValue

func (l ListValue) Print() {
	fmt.Print("(")
	for _, v := range l {
		v.Print()
		fmt.Print(" ")
	}
	fmt.Print(")")
}

type LambdaValue struct {
	ctx CtxStackNode
	expr ASTNode
}

func (lam LambdaValue) Print()  {
	fmt.Print("<Lambda Value>")
}

func (lam LambdaValue) Eval(args []ASTNode, top CtxStackNode) LispValue {
	var eargs ListValue
	for _, v := range args {
		eargs = append(eargs, EvalWithStack(v, top))
	}

	stack := &LambdaStack{previous: lam.ctx, args: eargs}
	return EvalWithStack(lam.expr, stack)
}

type MacroValue struct {
	ctx CtxStackNode
	expr ASTNode
}

func (mac MacroValue) Print() {
	fmt.Print("<Macro Value>")
}

func (mac MacroValue) Eval(args []ASTNode, top CtxStackNode) LispValue {
	var largs ListValue
	for _,arg := range args {
		largs = append(largs, intoListFromAST(arg))
	}

	stack := &MacroStack{previous: mac.ctx, args: largs}
	result := EvalWithStack(mac.expr, stack)
	return EvalWithStack(intoASTFromList(result), top)
}

func intoListFromAST(ast ASTNode) LispValue {
	switch ast := ast.(type) {
	case ASTAtom:
		return AtomValue{
			atom: ast.Value,
			isDouble: true,
		}
	case ASTString:
		return StringValue(ast.Value)
	case ASTIdent:
		return AtomValue{
			atom: ast.Value,
			isDouble: false,
		}
	case ASTNum:
		return NumValue(ast.Value)
	case ASTList:
		var list ListValue
		for _,v := range ast.Value {
			list = append(list, intoListFromAST(v))
		}
		return list
	default:
		panic("Unknown AST")
	}
}

func intoASTFromList(value LispValue) ASTNode {
	switch value := value.(type) {
	case AtomValue:
		if !value.isDouble {
			// single
			return ASTIdent{Value: value.atom}
		} else {
			// double
			return ASTAtom{Value: value.atom}
		}
	case StringValue:
		return ASTString{Value: string(value)}
	case NumValue:
		return ASTNum{Value: float64(value)}
	case ListValue:
		var list ASTList
		for _, v := range value {
			list.Value = append(list.Value, intoASTFromList(v))
		}
		return list
	default:
		panic(fmt.Sprintf("'%v' can't convert to AST", value))
	}
}
type BuiltinValue struct {
	builtin BuiltinFunc
}

func (b BuiltinValue) Print() {
	fmt.Printf("<Builtin (%s)>", b.builtin.Name())
}

func (b BuiltinValue) Eval(args []ASTNode, top CtxStackNode) LispValue {
	return b.builtin.OnEval(args, top)
}