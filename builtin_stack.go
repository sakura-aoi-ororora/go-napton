package napton

// TODO: 定数値のBuiltinができるようにする
type GlobalHandleStack struct {
	previous CtxStackNode
	idents map[string]BuiltinFunc
}

func (gh *GlobalHandleStack) SetPrevious(previous CtxStackNode) {
	gh.previous = previous
}

func (gh *GlobalHandleStack) GetPrevious() CtxStackNode  {
	return gh.previous
}

func (gh *GlobalHandleStack) GetValue() LispValue {
	return nil
}

type IdentHandleStack struct {
	previous CtxStackNode
	handler  ASTNode
}

func (ih *IdentHandleStack) SetPrevious(previous CtxStackNode) {
	ih.previous = previous
}

func (ih *IdentHandleStack) GetPrevious() CtxStackNode {
	return ih.previous
}

func (ih *IdentHandleStack) GetValue() LispValue {
	return nil
}

type IdentStack struct {
	previous CtxStackNode
	ident    AtomValue
}

func (is *IdentStack) SetPrevious(previous CtxStackNode) {
	is.previous = previous
}

func (is *IdentStack) GetPrevious() CtxStackNode {
	return is.previous
}

func (is *IdentStack) GetValue() LispValue {
	return is.ident
}

type LambdaStack struct {
	previous CtxStackNode
	args ListValue
}

func (l *LambdaStack) SetPrevious(previous CtxStackNode) {
	l.previous = previous
}

func (l *LambdaStack) GetPrevious() CtxStackNode {
	return l.previous
}

func (l *LambdaStack) GetValue() LispValue {
	return nil
}

type MacroStack struct {
	previous CtxStackNode
	args ListValue
}

func (m *MacroStack) SetPrevious(previous CtxStackNode) {
	m.previous = previous
}

func (m *MacroStack) GetPrevious() CtxStackNode {
	return m.previous
}

func (m *MacroStack) GetValue() LispValue {
	return nil
}
