package napton

import (
	"testing"
)

func Test_parser_peek(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    rune
		wantErr bool
	}{
		{
			name: "Peek Test 1",
			fields: fields{
				code: "123456",
				pos:  0,
			},
			want:    '1',
			wantErr: false,
		},
		{
			name: "Peek Test Error",
			fields: fields{
				code: "123",
				pos:  3,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			got, err := p.peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parser.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_parser_isEOF(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "isEOF Test 1",
			fields: fields{
				code: "123456",
				pos:  6,
			},
			want: true,
		},
		{
			name: "isEOF Test 2",
			fields: fields{
				code: "123456",
				pos:  5,
			},
			want: false,
		},
		{
			name: "isEOF Test 3",
			fields: fields{
				code: "",
				pos:  0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			if got := p.isEOF(); got != tt.want {
				t.Errorf("parser.isEOF() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_parser_parseNum(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    ASTNum
		wantErr bool
	}{
		{
			name: "ParseNum Success",
			fields: fields{
				code: "123456",
				pos:  0,
			},
			want: ASTNum{
				Value:      123456,
				RangeBegin: 0,
				RangeEnd:   5,
			},
			wantErr: false,
		},
		{
			name: "ParseNum with Trailing Characters",
			fields: fields{
				code: "789xyz",
				pos:  0,
			},
			want: ASTNum{
				Value:      789,
				RangeBegin: 0,
				RangeEnd:   2,
			},
			wantErr: false,
		},
		{
			name: "ParseNum at Position 3",
			fields: fields{
				code: "12",
				pos:  3,
			},
			want:    ASTNum{},
			wantErr: true,
		},
		{
			name: "ParseNum Empty Code",
			fields: fields{
				code: "",
				pos:  0,
			},
			want:    ASTNum{},
			wantErr: true,
		},
		{
			name: "ParseNum Single Digit",
			fields: fields{
				code: "5",
				pos:  0,
			},
			want: ASTNum{
				Value:      5,
				RangeBegin: 0,
				RangeEnd:   0,
			},
			wantErr: false,
		},
		{
			name: "ParseNum Leading Zeros",
			fields: fields{
				code: "000123",
				pos:  0,
			},
			want: ASTNum{
				Value:      123,
				RangeBegin: 0,
				RangeEnd:   5,
			},
			wantErr: false,
		},
		{
			name: "ParseNum Non-Digit at Start",
			fields: fields{
				code: "a123",
				pos:  0,
			},
			want:    ASTNum{},
			wantErr: true,
		},
		{
			name: "ParseNum Multiple Numbers",
			fields: fields{
				code: "42 and 7",
				pos:  0,
			},
			want: ASTNum{
				Value:      42,
				RangeBegin: 0,
				RangeEnd:   1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			got, err := p.parseNum()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parseNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Value != tt.want.Value {
					t.Errorf("parser.parseNum() Value = %v, want %v", got.Value, tt.want.Value)
				}
				if got.RangeBegin != tt.want.RangeBegin {
					t.Errorf("parser.parseNum() RangeBegin = %v, want %v", got.RangeBegin, tt.want.RangeBegin)
				}
				if got.RangeEnd != tt.want.RangeEnd {
					t.Errorf("parser.parseNum() RangeEnd = %v, want %v", got.RangeEnd, tt.want.RangeEnd)
				}
			}
		})
	}
}
func Test_parser_parseIdent(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    ASTIdent
		wantErr bool
	}{
		{
			name: "ParseIdent Success",
			fields: fields{
				code: "identifier",
				pos:  0,
			},
			want: ASTIdent{
				Value:      "identifier",
				RangeBegin: 0,
				RangeEnd:   9,
			},
			wantErr: false,
		},
		{
			name: "ParseIdent with Numbers",
			fields: fields{
				code: "var123",
				pos:  0,
			},
			want: ASTIdent{
				Value:      "var123",
				RangeBegin: 0,
				RangeEnd:   5,
			},
			wantErr: false,
		},
		{
			name: "ParseIdent Empty Code",
			fields: fields{
				code: "",
				pos:  0,
			},
			want:    ASTIdent{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			got, err := p.parseIdent()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parseIdent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Value != tt.want.Value {
					t.Errorf("parser.parseIdent() Value = %v, want %v", got.Value, tt.want.Value)
				}
				if got.RangeBegin != tt.want.RangeBegin {
					t.Errorf("parser.parseIdent() RangeBegin = %v, want %v", got.RangeBegin, tt.want.RangeBegin)
				}
				if got.RangeEnd != tt.want.RangeEnd {
					t.Errorf("parser.parseIdent() RangeEnd = %v, want %v", got.RangeEnd, tt.want.RangeEnd)
				}
			}
		})
	}
}

func Test_parser_parseString(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    ASTString
		wantErr bool
	}{
		{
			name: "ParseString Success",
			fields: fields{
				code: `"hello world"`,
				pos:  0,
			},
			want: ASTString{
				Value:      "hello world",
				RangeBegin: 0,
				RangeEnd:   12,
			},
			wantErr: false,
		},
		{
			name: "ParseString with Escapes",
			fields: fields{
				code: `"line\nnewline"`,
				pos:  0,
			},
			want: ASTString{
				Value:      "line\nnewline",
				RangeBegin: 0,
				RangeEnd:   14,
			},
			wantErr: false,
		},
		{
			name: "ParseString Missing Closing Quote",
			fields: fields{
				code: `"unterminated`,
				pos:  0,
			},
			want:    ASTString{},
			wantErr: true,
		},
		{
			name: "ParseString Empty String",
			fields: fields{
				code: `""`,
				pos:  0,
			},
			want: ASTString{
				Value:      "",
				RangeBegin: 0,
				RangeEnd:   1,
			},
			wantErr: false,
		},
		{
			name: "ParseString with Unknown Escape",
			fields: fields{
				code: `"bad\escape"`,
				pos:  0,
			},
			want:    ASTString{},
			wantErr: true,
		},
		{
			name: "ParseString Single Character",
			fields: fields{
				code: `"a"`,
				pos:  0,
			},
			want: ASTString{
				Value:      "a",
				RangeBegin: 0,
				RangeEnd:   2,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			got, err := p.parseString()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.Value != tt.want.Value {
					t.Errorf("parser.parseString() Value = %v, want %v", got.Value, tt.want.Value)
				}
				if got.RangeBegin != tt.want.RangeBegin {
					t.Errorf("parser.parseString() RangeBegin = %v, want %v", got.RangeBegin, tt.want.RangeBegin)
				}
				if got.RangeEnd != tt.want.RangeEnd {
					t.Errorf("parser.parseString() RangeEnd = %v, want %v", got.RangeEnd, tt.want.RangeEnd)
				}
			}
		})
	}
}

func Test_parser_parseList(t *testing.T) {
	type fields struct {
		code string
		pos  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    ASTList
		wantErr bool
	}{
		{
			name: "ParseList Empty",
			fields: fields{
				code: "()",
				pos:  0,
			},
			want: ASTList{
				Value:      []ASTNode{},
				RangeBegin: 0,
				RangeEnd:   1,
			},
			wantErr: false,
		},
		{
			name: "ParseList with Numbers and Identifiers",
			fields: fields{
				code: "(123 var \"string\")",
				pos:  0,
			},
			want: ASTList{
				Value: []ASTNode{
					ASTNum{Value: 123, RangeBegin: 1, RangeEnd: 3},
					ASTIdent{Value: "var", RangeBegin: 5, RangeEnd: 7},
					ASTString{Value: "string", RangeBegin: 9, RangeEnd: 16},
				},
				RangeBegin: 0,
				RangeEnd:   17,
			},
			wantErr: false,
		},
		{
			name: "ParseList Nested Lists",
			fields: fields{
				code: "(1 (2 3) (4 (5)))",
				pos:  0,
			},
			want: ASTList{
				Value: []ASTNode{
					ASTNum{Value: 1, RangeBegin: 1, RangeEnd: 1},
					ASTList{
						Value: []ASTNode{
							ASTNum{Value: 2, RangeBegin: 4, RangeEnd: 4},
							ASTNum{Value: 3, RangeBegin: 6, RangeEnd: 6},
						},
						RangeBegin: 3,
						RangeEnd:   7,
					},
					ASTList{
						Value: []ASTNode{
							ASTNum{Value: 4, RangeBegin: 10, RangeEnd: 10},
							ASTList{
								Value: []ASTNode{
									ASTNum{Value: 5, RangeBegin: 13, RangeEnd: 13},
								},
								RangeBegin: 12,
								RangeEnd:   14,
							},
						},
						RangeBegin: 9,
						RangeEnd:   15,
					},
				},
				RangeBegin: 0,
				RangeEnd:   16,
			},
			wantErr: false,
		},
		{
			name: "ParseList Missing Closing Parenthesis",
			fields: fields{
				code: "(1 2 3",
				pos:  0,
			},
			want:    ASTList{},
			wantErr: true,
		},
		{
			name: "ParseList Single Element",
			fields: fields{
				code: "(single)",
				pos:  0,
			},
			want: ASTList{
				Value: []ASTNode{
					ASTIdent{Value: "single", RangeBegin: 1, RangeEnd: 6},
				},
				RangeBegin: 0,
				RangeEnd:   7,
			},
			wantErr: false,
		},
		{
			name: "ParseList With Strings and Escapes",
			fields: fields{
				code: `("hello" "world\n")`,
				pos:  0,
			},
			want: ASTList{
				Value: []ASTNode{
					ASTString{Value: "hello", RangeBegin: 1, RangeEnd: 7},
					ASTString{Value: "world\n", RangeBegin: 9, RangeEnd: 17},
				},
				RangeBegin: 0,
				RangeEnd:   18,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				code: tt.fields.code,
				pos:  tt.fields.pos,
			}
			got, err := p.parseList()
			if (err != nil) != tt.wantErr {
				t.Errorf("parser.parseList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.RangeBegin != tt.want.RangeBegin {
					t.Errorf("parser.parseList() RangeBegin = %v, want %v", got.RangeBegin, tt.want.RangeBegin)
				}
				if got.RangeEnd != tt.want.RangeEnd {
					t.Errorf("parser.parseList() RangeEnd = %v, want %v", got.RangeEnd, tt.want.RangeEnd)
				}

				if len(got.Value) != len(tt.want.Value) {
					t.Errorf("parser.parseList() Value length = %v, want %v", len(got.Value), len(tt.want.Value))
				} else {
					for i := range got.Value {
						switch gotNode := got.Value[i].(type) {
						case ASTNum:
							wantNum, ok := tt.want.Value[i].(ASTNum)
							if !ok {
								t.Errorf("parser.parseList() node type mismatch at index %d", i)
								continue
							}
							if gotNode != wantNum {
								t.Errorf("parser.parseList() ASTNum at index %d = %v, want %v", i, gotNode, wantNum)
							}
						case ASTIdent:
							wantIdent, ok := tt.want.Value[i].(ASTIdent)
							if !ok {
								t.Errorf("parser.parseList() node type mismatch at index %d", i)
								continue
							}
							if gotNode != wantIdent {
								t.Errorf("parser.parseList() ASTIdent at index %d = %v, want %v", i, gotNode, wantIdent)
							}
						case ASTString:
							wantString, ok := tt.want.Value[i].(ASTString)
							if !ok {
								t.Errorf("parser.parseList() node type mismatch at index %d", i)
								continue
							}
							if gotNode != wantString {
								t.Errorf("parser.parseList() ASTString at index %d = %v, want %v", i, gotNode, wantString)
							}
						case ASTList:
							wantList, ok := tt.want.Value[i].(ASTList)
							if !ok {
								t.Errorf("parser.parseList() node type mismatch at index %d", i)
								continue
							}
							if gotNode.RangeBegin != wantList.RangeBegin || gotNode.RangeEnd != wantList.RangeEnd {
								t.Errorf("parser.parseList() ASTList range at index %d = (%v, %v), want (%v, %v)", i, gotNode.RangeBegin, gotNode.RangeEnd, wantList.RangeBegin, wantList.RangeEnd)
							}
							// Further nested checks can be added if necessary
						default:
							t.Errorf("parser.parseList() unknown node type at index %d", i)
						}
					}
				}
			}
		})
	}
}
