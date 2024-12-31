// TODO: エラー表示を見やすくする

package napton

import "fmt"

func checkIdent(r rune) bool {
	return r != ' ' &&
		r != '\t' &&
		r != '\n' &&
		r != '\r' &&
		r != '(' &&
		r != ')' &&
		r != '"' &&
		r != ':'
}

type parser struct {
	code string
	pos  int
}

func (p *parser) peek() (rune, error) {
	if p.pos >= len(p.code) {
		return 0, fmt.Errorf("eof")
	}

	return rune(p.code[p.pos]), nil
}

func (p *parser) move() error {
	if p.pos >= len(p.code) {
		return fmt.Errorf("eof")
	}

	p.pos++
	return nil
}

func (p *parser) isEOF() bool {
	return p.pos >= len(p.code)
}

func (p *parser) getPos() int {
	return p.pos
}

func (p *parser) parseSpace() error {
	for {
		r, err := p.peek()
		if err != nil {
			if p.isEOF() {
				break
			} else {
				return err
			}
		}

		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			err := p.move()
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	return nil
}

func (p *parser) parseNum() (ASTNum, error) {
	if r, err := p.peek(); err != nil || r < '0' || r > '9' {
		return ASTNum{}, fmt.Errorf("expected number")
	}

	value := ASTNum{}

	value.RangeBegin = p.getPos()

	for {
		r, err := p.peek()
		if err != nil {
			if p.isEOF() {
				break
			} else {
				return ASTNum{}, err
			}
		}

		if r >= '0' && r <= '9' {
			value.Value = value.Value*10 + float64(r-'0')
			err := p.move()
			if err != nil {
				return ASTNum{}, err
			}
		} else {
			break
		}
	}

	value.RangeEnd = p.getPos() - 1

	return value, nil
}

func (p *parser) parseIdent() (ASTIdent, error) {
	if r, err := p.peek(); err != nil || !checkIdent(r) {
		return ASTIdent{}, fmt.Errorf("expected ident")
	}

	value := ASTIdent{}
	value.RangeBegin = p.getPos()

	for {
		r, err := p.peek()
		if err != nil {
			if p.isEOF() {
				break
			} else {
				return ASTIdent{}, err
			}
		}

		if checkIdent(r) {
			value.Value += string(r)
			err := p.move()
			if err != nil {
				return ASTIdent{}, err
			}
		} else {
			break
		}
	}

	value.RangeEnd = p.getPos() - 1

	return value, nil
}

func (p *parser) parseString() (ASTString, error) {
	var value ASTString

	value.RangeBegin = p.getPos()

	if r, err := p.peek(); err != nil || r != '"' {
		return ASTString{}, fmt.Errorf("expected '\"'")
	}

	if err := p.move(); err != nil {
		return ASTString{}, err
	}

	for {
		r, err := p.peek()
		if err != nil {
			return ASTString{}, err
		}

		if r == '"' {
			if err := p.move(); err != nil {
				return ASTString{}, err
			}
			break
		}

		if r == '\\' {
			if err := p.move(); err != nil {
				return ASTString{}, err
			}

			r, err = p.peek()
			if err != nil {
				return ASTString{}, err
			}

			switch r {
			case 'n':
				value.Value += "\n"
			case 't':
				value.Value += "\t"
			case '\\':
				value.Value += "\\"
			case '"':
				value.Value += "\""
			default:
				return ASTString{}, fmt.Errorf("unknown escape sequence: \\%c", r)
			}
		} else {
			value.Value += string(r)
		}

		if err := p.move(); err != nil {
			return ASTString{}, err
		}
	}

	value.RangeEnd = p.getPos() - 1

	return value, nil
}

func (p *parser) parseAtom() (ASTAtom, error) {
	value := ASTAtom{}

	r, err := p.peek()
	if err != nil {
		return ASTAtom{}, err
	} else if r != ':' {
		return ASTAtom{}, fmt.Errorf("expected ':' but found '%#v'", r)
	}

	if err := p.move(); err != nil {
		return ASTAtom{}, err
	}

	for {
		r, err := p.peek()
		if err != nil {
			if p.isEOF() {
				break
			} else {
				return ASTAtom{}, err
			}
		}

		if checkIdent(r) {
			value.Value += string(r)
			err := p.move()
			if err != nil {
				return ASTAtom{}, err
			}
		} else {
			break
		}
	}

	return value, nil
}

func (p *parser) parseList() (ASTList, error) {
	list := ASTList{
		Value: make([]ASTNode, 0),
	}

	list.RangeBegin = p.getPos()

	r, err := p.peek()
	if err != nil {
		return ASTList{}, err
	}

	if r != '(' {
		return ASTList{}, fmt.Errorf("expected '('")
	}

	if err := p.move(); err != nil {
		return ASTList{}, err
	}

	for {
		err := p.parseSpace()
		if err != nil {
			return ASTList{}, err
		}

		r, err := p.peek()
		if err != nil {
			if p.isEOF() {
				return ASTList{}, fmt.Errorf("expect ')' but found eof")
			} else {
				return ASTList{}, err
			}
		}

		if r == ')' {
			if err := p.move(); err != nil {
				return ASTList{}, err
			}

			break
		}

		node, err := p.parseAll()
		if err != nil {
			return ASTList{}, err
		}

		list.Value = append(list.Value, node)
	}

	list.RangeEnd = p.getPos() - 1

	return list, nil
}

func (p *parser) parseAll() (ASTNode, error) {
	if err := p.parseSpace(); err != nil {
		return nil, err
	}

	r, err := p.peek()
	if err != nil {
		return nil, err
	}

	if r >= '0' && r <= '9' {
		f, err := p.parseNum()
		if err != nil {
			return nil, err
		}

		return f, nil
	}

	if r == ':' {
		a, err := p.parseAtom()
		if err != nil {
			return nil, err
		}

		return a, nil
	}

	if checkIdent(r) {
		i, err := p.parseIdent()
		if err != nil {
			return nil, err
		}

		return i, nil
	}

	if r == '"' {
		s, err := p.parseString()
		if err != nil {
			return nil, err
		}

		return s, nil
	}

	if r == '(' {
		l, err := p.parseList()
		if err != nil {
			return nil, err
		}

		return l, nil
	}

	return nil, fmt.Errorf("unknown token: %v", r)
}

func Parse(code string) (ASTNode, error) {
	p := parser{code: code, pos: 0}

	return p.parseAll()
}
