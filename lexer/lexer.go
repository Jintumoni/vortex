package lexer

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type LexerInterface interface {
	GetNextToken() *Token
	PeekNextToken() *Token
	GetSourceContext() string
	Commit()
	Rollback()
}

type Lexer struct {
	PeekIndex    int
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
		PeekIndex:    0,
		CurrentToken: nil,
		Input:        data,
	}
}

func (l *Lexer) Peek() byte {
	if l.PeekIndex+1 < len(l.Input) {
		return l.Input[l.PeekIndex+1]
	}

	return 0
}

func (l *Lexer) Commit() {
	l.Index = l.PeekIndex
}

func (l *Lexer) Rollback() {
	l.PeekIndex = l.Index
}

func (l *Lexer) advance() {
	l.Col++
	if l.Input[l.PeekIndex] == '\n' {
		l.Row++
		l.Col = 0
	}
	l.PeekIndex++
}

// Use this method only for single line tokens
// Use the Token{} struct directly for multiline tokens
func (l *Lexer) addSLToken(token TokenType, value string) *Token {
	return &Token{
		Type:  token,
		Value: value,
		Row:   l.Row,
		Col:   l.Col - len(value),
		Span:  len(value),
	}
}

func (l *Lexer) ignoreSpace() {
	for l.PeekIndex < len(l.Input) && unicode.IsSpace(rune(l.Input[l.PeekIndex])) {
		l.advance()
	}
}

func (l *Lexer) getNumberToken() *Token {
	if !unicode.IsNumber(rune(l.Input[l.PeekIndex])) {
		return l.addSLToken(TokenInvalid, "INVALID")
	}

	row, col := l.Row, l.Col
	buffer := bytes.Buffer{}
	for l.PeekIndex < len(l.Input) && unicode.IsNumber(rune(l.Input[l.PeekIndex])) {
		buffer.WriteByte(l.Input[l.PeekIndex])
		l.advance()
	}

	return &Token{TokenIntegerConstant, buffer.String(), row, col, buffer.Len()}
}

func (l *Lexer) getIDToken() *Token {
	if !unicode.IsLetter(rune(l.Input[l.PeekIndex])) {
		return &Token{TokenInvalid, "INVALID", l.Row, l.Col, 1}
	}

	row, col := l.Row, l.Col
	buffer := bytes.Buffer{}
	for l.PeekIndex < len(l.Input) {
		if unicode.IsLetter(rune(l.Input[l.PeekIndex])) || unicode.IsNumber(rune(l.Input[l.PeekIndex])) {
			buffer.WriteByte(l.Input[l.PeekIndex])
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

	if l.Input[l.PeekIndex] != '"' {
		return l.addSLToken(TokenInvalid, "INVALID")
	}
	l.advance()

	row, col := l.Row, l.Col
	for l.PeekIndex < len(l.Input) && l.Input[l.PeekIndex] != '"' {
		buffer.WriteByte(l.Input[l.PeekIndex])
		l.advance()
	}

	if l.PeekIndex >= len(l.Input) || l.Input[l.PeekIndex] != '"' {
		return l.addSLToken(TokenInvalid, "INVALID")
	}
	l.advance()

	return &Token{TokenStringConstant, buffer.String(), row, col, buffer.Len()}
}

func (l *Lexer) GetNextToken() *Token {
	l.ignoreSpace()

	if l.PeekIndex >= len(l.Input) {
		return l.addSLToken(TokenEOF, "EOF")
	}

	if unicode.IsLetter(rune(l.Input[l.PeekIndex])) {
		return l.getIDToken()
	} else if unicode.IsNumber(rune(l.Input[l.PeekIndex])) {
		return l.getNumberToken()
	}

	switch l.Input[l.PeekIndex] {
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
	case '|':
		l.advance()
		return l.addSLToken(TokenOr, "|")
	case '&':
		l.advance()
		return l.addSLToken(TokenAnd, "&")
	case '#':
		l.advance()
		return l.addSLToken(TokenHash, "#")
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

func (l *Lexer) PeekNextToken() *Token {
	token := l.GetNextToken()
	l.Rollback()
	return token
}

func (l *Lexer) GetSourceContext() string {
	lines := strings.Split(string(l.Input), "\n")
	source := new(bytes.Buffer)

	i := max(0, l.Row-2)
	for i <= l.Row {
		source.WriteString(fmt.Sprintf("%d\t|\t", i+1))
		source.WriteString(lines[i])
		source.WriteString("\n")
		i++
	}

	return source.String()
}
