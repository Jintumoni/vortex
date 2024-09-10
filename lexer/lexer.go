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
	CurrentToken *Token
	Input        []byte
}

func NewLexer(r io.Reader) *Lexer {
  data, _ := io.ReadAll(r)
  return &Lexer{
    Index: 0,
    CurrentToken: nil,
    Input: data,
  }
}

func (l *Lexer) Peek() byte {
	if l.Index+1 < len(l.Input) {
		return l.Input[l.Index+1]
	}

	return 0
}

func (l *Lexer) advance() {
	l.Index++
}

func (l *Lexer) ignoreSpace() {
	for l.Index < len(l.Input) && unicode.IsSpace(rune(l.Input[l.Index])) {
		l.advance()
	}
}

func (l *Lexer) getNumberToken() *Token {
  if !unicode.IsNumber(rune(l.Input[l.Index])) {
    return &Token{TokenInvalid, "INVALID"}
  }

	buffer := bytes.Buffer{}
	for l.Index < len(l.Input) && unicode.IsNumber(rune(l.Input[l.Index])) {
		buffer.WriteByte(l.Input[l.Index])
		l.advance()
	}

	return &Token{TokenIntegerConstant, buffer.String()}
}

func (l *Lexer) getIDToken() *Token {
  if !unicode.IsLetter(rune(l.Input[l.Index])) {
    return &Token{TokenInvalid, "INVALID"}
  }

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
    return &Token{tokenType, buffer.String()}
  }

	return &Token{TokenIdentifier, buffer.String()}
}

func (l *Lexer) getStringToken() *Token {
	buffer := bytes.Buffer{}

  if l.Input[l.Index] != '"' {
    return &Token{TokenInvalid, "INVALID"}
  }
	l.advance()

	for l.Index < len(l.Input) && l.Input[l.Index] != '"' {
		buffer.WriteByte(l.Input[l.Index])
		l.advance()
	}

  if l.Index >= len(l.Input) || l.Input[l.Index] != '"' {
    return &Token{TokenInvalid, "INVALID"}
  }
	l.advance()

	return &Token{TokenStringConstant, buffer.String()}
}

func (l *Lexer) GetNextToken() *Token {
	l.ignoreSpace()

  if l.Index >= len(l.Input) {
    return &Token{TokenEOF, "EOF"}
  }

	if unicode.IsLetter(rune(l.Input[l.Index])) {
		return l.getIDToken()
	} else if unicode.IsNumber(rune(l.Input[l.Index])) {
		return l.getNumberToken()
	}

	switch l.Input[l.Index] {
	case '+':
		l.advance()
		return &Token{TokenPlus, "+"}
	case '-':
		l.advance()
		return &Token{TokenMinus, "-"}
	case '*':
		l.advance()
		return &Token{TokenMultiply, "*"}
	case '/':
		l.advance()
		return &Token{TokenDivide, "/"}
	case ',':
		l.advance()
		return &Token{TokenComma, ","}
	case '{':
		l.advance()
		return &Token{TokenLCB, "{"}
	case '}':
		l.advance()
		return &Token{TokenRCB, "}"}
	case '(':
		l.advance()
		return &Token{TokenLRB, "("}
	case ')':
		l.advance()
		return &Token{TokenRRB, ")"}
	case '[':
		l.advance()
		return &Token{TokenLSB, "["}
	case ']':
		l.advance()
		return &Token{TokenRSB, "]"}
	case '.':
		if l.Peek() == '.' {
      l.advance()
      l.advance()
			return &Token{TokenRange, ".."}
		} else {
      l.advance()
			return &Token{TokenDot, "."}
		}
	case '<':
		if l.Peek() == '=' {
      l.advance()
      l.advance()
			return &Token{TokenLessThanEqual, "<="}
		} else {
      l.advance()
			return &Token{TokenLessThan, "<"}
		}
	case '>':
		if l.Peek() == '=' {
      l.advance()
      l.advance()
			return &Token{TokenGreaterThanEqual, ">="}
		} else {
      l.advance()
			return &Token{TokenGreaterThan, ">"}
		}
	case '=':
    l.advance()
		return &Token{TokenEqual, "="}
	case '"':
		return l.getStringToken()
	default:
    l.advance()
		return &Token{TokenInvalid, "INVALID"}
	}
}
