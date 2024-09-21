package lexer

import (
	"bytes"
	"io"
	"unicode"
)

type LexerInterface interface {
	GetNextToken() *Token
}

type Lexer struct {
	Index        int
	Row          int
	Col          int
	CurrentToken *Token
	Input        []byte
}

func NewLexer(r io.Reader) *Lexer {
	data, _ := io.ReadAll(r)
	return &Lexer{
		Index:        0,
		CurrentToken: nil,
		Input:        data,
	}
}

func (l *Lexer) Peek() byte {
	if l.Index+1 < len(l.Input) {
		return l.Input[l.Index+1]
	}

	return 0
}

func (l *Lexer) advance() {
	l.Col++
	if l.Input[l.Index] == '\n' {
		l.Row++
		l.Col = 0
	}
	l.Index++
}

// Use this method only for single line tokens
// Use the Token{} struct directly for multiline tokens
func (l *Lexer) addSLToken(token TokenType, value string) *Token {
  return &Token{
    Type: token,
    Value: value,
    Row: l.Row,
    Col: l.Col - len(value),
    Span: len(value),
  }
}

func (l *Lexer) ignoreSpace() {
	for l.Index < len(l.Input) && unicode.IsSpace(rune(l.Input[l.Index])) {
		l.advance()
	}
}

func (l *Lexer) getNumberToken() *Token {
	if !unicode.IsNumber(rune(l.Input[l.Index])) {
		return l.addSLToken(TokenInvalid, "INVALID")
	}

	row, col := l.Row, l.Col
	buffer := bytes.Buffer{}
	for l.Index < len(l.Input) && unicode.IsNumber(rune(l.Input[l.Index])) {
		buffer.WriteByte(l.Input[l.Index])
		l.advance()
	}

	return &Token{TokenIntegerConstant, buffer.String(), row, col, buffer.Len()}
}

func (l *Lexer) getIDToken() *Token {
	if !unicode.IsLetter(rune(l.Input[l.Index])) {
		return &Token{TokenInvalid, "INVALID", l.Row, l.Col, 1}
	}

	row, col := l.Row, l.Col
	buffer := bytes.Buffer{}
	for l.Index < len(l.Input) {
		if unicode.IsLetter(rune(l.Input[l.Index])) || unicode.IsNumber(rune(l.Input[l.Index])) {
			buffer.WriteByte(l.Input[l.Index])
		} else {
			break
		}
		l.advance()
	}

	tokenType, ok := ReservedKeywords[buffer.String()]
	if ok {
		return &Token{tokenType, buffer.String(), row, col, buffer.Len()}
	}

	return &Token{TokenIdentifier, buffer.String(), row, col, buffer.Len()}
}

func (l *Lexer) getStringToken() *Token {
	buffer := bytes.Buffer{}

	if l.Input[l.Index] != '"' {
		return l.addSLToken(TokenInvalid, "INVALID")
	}
	l.advance()

	row, col := l.Row, l.Col
	for l.Index < len(l.Input) && l.Input[l.Index] != '"' {
		buffer.WriteByte(l.Input[l.Index])
		l.advance()
	}

	if l.Index >= len(l.Input) || l.Input[l.Index] != '"' {
		return l.addSLToken(TokenInvalid, "INVALID")
	}
	l.advance()

	return &Token{TokenStringConstant, buffer.String(), row, col, buffer.Len()}
}

func (l *Lexer) GetNextToken() *Token {
	l.ignoreSpace()

	if l.Index >= len(l.Input) {
		return l.addSLToken(TokenEOF, "EOF")
	}

	if unicode.IsLetter(rune(l.Input[l.Index])) {
		return l.getIDToken()
	} else if unicode.IsNumber(rune(l.Input[l.Index])) {
		return l.getNumberToken()
	}

	switch l.Input[l.Index] {
	case '+':
		l.advance()
		return l.addSLToken(TokenPlus, "+")
	case '-':
		l.advance()
		return l.addSLToken(TokenMinus, "-")
	case '*':
		l.advance()
		return l.addSLToken(TokenMultiply, "*")
	case '/':
		l.advance()
		return l.addSLToken(TokenDivide, "/")
	case ',':
		l.advance()
		return l.addSLToken(TokenComma, ",")
	case '{':
		l.advance()
		return l.addSLToken(TokenLCB, "{")
	case '}':
		l.advance()
		return l.addSLToken(TokenRCB, "}")
	case '(':
		l.advance()
		return l.addSLToken(TokenLRB, "(")
	case ')':
		l.advance()
		return l.addSLToken(TokenRRB, ")")
	case '[':
		l.advance()
		return l.addSLToken(TokenLSB, "[")
	case ']':
		l.advance()
		return l.addSLToken(TokenRSB, "]")
	case '.':
		if l.Peek() == '.' {
			l.advance()
			l.advance()
			return l.addSLToken(TokenRange, "..")
		} else {
			l.advance()
			return l.addSLToken(TokenDot, ".")
		}
	case '<':
		if l.Peek() == '=' {
			l.advance()
			l.advance()
			return l.addSLToken(TokenLessThanEqual, "<=")
		} else {
			l.advance()
			return l.addSLToken(TokenLessThan, "<")
		}
	case '>':
		if l.Peek() == '=' {
			l.advance()
			l.advance()
			return l.addSLToken(TokenGreaterThanEqual, ">=")
		} else {
			l.advance()
			return l.addSLToken(TokenGreaterThan, ">")
		}
	case '=':
		l.advance()
		return l.addSLToken(TokenEqual, "=")
	case '"':
		return l.getStringToken()
	default:
		l.advance()
		return l.addSLToken(TokenInvalid, "INVALID")
	}
}
