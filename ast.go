package napton

type ASTNode interface {
	GetRange() (int, int)
}

type ASTNum struct {
	Value      float64
	RangeBegin int
	RangeEnd   int
}

func (n ASTNum) GetRange() (int, int) {
	return n.RangeBegin, n.RangeEnd
}

type ASTString struct {
	Value      string
	RangeBegin int
	RangeEnd   int
}

func (s ASTString) GetRange() (int, int) {
	return s.RangeBegin, s.RangeEnd
}

type ASTIdent struct {
	Value      string
	RangeBegin int
	RangeEnd   int
}

func (i ASTIdent) GetRange() (int, int) {
	return i.RangeBegin, i.RangeEnd
}

type ASTAtom struct {
	Value      string
	RangeBegin int
	RangeEnd   int
}

func (i ASTAtom) GetRange() (int, int) {
	return i.RangeBegin, i.RangeEnd
}

type ASTList struct {
	Value      []ASTNode
	RangeBegin int
	RangeEnd   int
}

func (l ASTList) GetRange() (int, int) {
	return l.RangeBegin, l.RangeEnd
}
