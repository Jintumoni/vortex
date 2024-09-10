package lexer

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestAdvance(t *testing.T) {
	mockData := "Hello"
	l := NewLexer(strings.NewReader(mockData))
	if l.Index != 0 {
		assert.Equal(t, 0, l.Index)
	}
	l.advance()
	if l.Index != 1 {
		assert.Equal(t, 1, l.Index)
	}
}

func TestPeek(t *testing.T) {
	mockData := "Hello"
	l := NewLexer(strings.NewReader(mockData))
	if l.Peek() != 'e' {
		assert.Equal(t, 'e', l.Peek())
	}
	for range mockData {
		l.advance()
	}
	if l.Peek() != 0 {
		assert.Equal(t, 0, l.Peek())
	}
}

func TestIgnoreSpace(t *testing.T) {
	mockData := "Hello  \t  \n\n     World"
	l := NewLexer(strings.NewReader(mockData))
	l.ignoreSpace()
	if l.Index != 0 {
		assert.Equal(t, 0, l.Index)
	}

	l.Index = strings.Index(mockData, " ")
	l.ignoreSpace()

	if l.Index != strings.Index(mockData, "World") {
		assert.Equal(t, 17, l.Index)
	}
}

func TestGetNumberToken(t *testing.T) {
	mockData := "123456 Hello"
	l := NewLexer(strings.NewReader(mockData))

	token := l.getNumberToken()
	if token.Value != "123456" || token.Type != TokenIntegerConstant {
		assert.Equal(t, "123456", token.Value)
	}
}

func TestIDToken(t *testing.T) {
	mockData := "Name123 var 123456"
	l := NewLexer(strings.NewReader(mockData))

	token := l.getIDToken()
	if token.Type != TokenIdentifier || token.Value != "Name123" {
		assert.Equal(t, "Name123", token.Value)
	}

	l.ignoreSpace()
	token = l.getIDToken()
	if token.Type != TokenIdentifier || token.Value != "var" {
		assert.Equal(t, "var", token.Value)
	}

	l.ignoreSpace()
	token = l.getIDToken()
	if token.Type == TokenIdentifier {
		assert.Equal(t, TokenIntegerConstant, token.Type)
	}
}

func TestGetStringTokenShouldReturnStringToken(t *testing.T) {
	mockData := `"this is a test 123 +=-/ >= <= ,.:? .. \n"`
	l := NewLexer(strings.NewReader(mockData))

	expected := `this is a test 123 +=-/ >= <= ,.:? .. \n`

	token := l.getStringToken()
	if token.Type != TokenStringConstant || token.Value != expected {
		assert.Equal(t, expected, token.Value)
	}
}

func TestGetStringTokenShouldReturnInvalidToken(t *testing.T) {
	mockData := `"Hello this is a wrong literal`
	l := NewLexer(strings.NewReader(mockData))
	token := l.getStringToken()
	if token.Type != TokenInvalid {
		assert.Equal(t, TokenInvalid, token.Type)
	}
}

func TestGetNextToken(t *testing.T) {
	mockData := `
  Person {
    FriendsWith Person {
      LivesIn Country{.name="India"}
      and .salary >= 100
    }
  }
  `
	l := NewLexer(strings.NewReader(mockData))

	expected := []struct {
		Type  TokenType
		Value string
	}{
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "LivesIn"},
		{TokenIdentifier, "Country"},
		{TokenLCB, "{"},
		{TokenDot, "."},
		{TokenIdentifier, "name"},
		{TokenEqual, "="},
		{TokenStringConstant, "India"},
		{TokenRCB, "}"},
		{TokenAnd, "and"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenGreaterThanEqual, ">="},
		{TokenIntegerConstant, "100"},
		{TokenRCB, "}"},
		{TokenRCB, "}"},
    {TokenEOF, "EOF"},
	}

	for _, e := range expected {
		token := l.GetNextToken()
		assert.Equal(t, e.Type, token.Type)
		assert.Equal(t, e.Value, token.Value)
	}
}

func TestGetNextTokenWithRangeExpression(t *testing.T) {
	mockData := `
  Person {
    FriendsWith Person {
      LivesIn[2..] Country{.name="India"}
      and .salary >= 100
    }
  }
  `
	l := NewLexer(strings.NewReader(mockData))

	expected := []struct {
		Type  TokenType
		Value string
	}{
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "LivesIn"},
		{TokenLSB, "["},
		{TokenIntegerConstant, "2"},
		{TokenRange, ".."},
		{TokenRSB, "]"},
		{TokenIdentifier, "Country"},
		{TokenLCB, "{"},
		{TokenDot, "."},
		{TokenIdentifier, "name"},
		{TokenEqual, "="},
		{TokenStringConstant, "India"},
		{TokenRCB, "}"},
		{TokenAnd, "and"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenGreaterThanEqual, ">="},
		{TokenIntegerConstant, "100"},
		{TokenRCB, "}"},
		{TokenRCB, "}"},
    {TokenEOF, "EOF"},
	}

	for _, e := range expected {
		token := l.GetNextToken()
		assert.Equal(t, e.Type, token.Type)
		assert.Equal(t, e.Value, token.Value)
	}
}

func TestGetNextTokenWithAlias(t *testing.T) {
	mockData := `
  Person as A {
    FriendsWith Person {
      LivesIn[2..] Country{.name="India"}
      and .salary >= A.salary
    }
  }
  `
	l := NewLexer(strings.NewReader(mockData))

	expected := []struct {
		Type  TokenType
		Value string
	}{
		{TokenIdentifier, "Person"},
		{TokenAlias, "as"},
		{TokenIdentifier, "A"},
		{TokenLCB, "{"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "LivesIn"},
		{TokenLSB, "["},
		{TokenIntegerConstant, "2"},
		{TokenRange, ".."},
		{TokenRSB, "]"},
		{TokenIdentifier, "Country"},
		{TokenLCB, "{"},
		{TokenDot, "."},
		{TokenIdentifier, "name"},
		{TokenEqual, "="},
		{TokenStringConstant, "India"},
		{TokenRCB, "}"},
		{TokenAnd, "and"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenGreaterThanEqual, ">="},
		{TokenIdentifier, "A"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenRCB, "}"},
		{TokenRCB, "}"},
    {TokenEOF, "EOF"},
	}

	for _, e := range expected {
		token := l.GetNextToken()
		assert.Equal(t, e.Type, token.Type)
		assert.Equal(t, e.Value, token.Value)
	}
}

func TestGetNextTokenWithFunctions(t *testing.T) {
	mockData := `
  Person as A {
    Sum(FriendsWith Person {
      LivesIn[2..] Country{.name="India"}
      and .salary >= A.salary
    }, .salary)
  }
  `
	l := NewLexer(strings.NewReader(mockData))

	expected := []struct {
		Type  TokenType
		Value string
	}{
		{TokenIdentifier, "Person"},
		{TokenAlias, "as"},
		{TokenIdentifier, "A"},
		{TokenLCB, "{"},
    {TokenFunction, "Sum"},
    {TokenLRB, "("},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenIdentifier, "LivesIn"},
		{TokenLSB, "["},
		{TokenIntegerConstant, "2"},
		{TokenRange, ".."},
		{TokenRSB, "]"},
		{TokenIdentifier, "Country"},
		{TokenLCB, "{"},
		{TokenDot, "."},
		{TokenIdentifier, "name"},
		{TokenEqual, "="},
		{TokenStringConstant, "India"},
		{TokenRCB, "}"},
		{TokenAnd, "and"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenGreaterThanEqual, ">="},
		{TokenIdentifier, "A"},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
		{TokenRCB, "}"},
    {TokenComma, ","},
		{TokenDot, "."},
		{TokenIdentifier, "salary"},
    {TokenRRB, ")"},
		{TokenRCB, "}"},
    {TokenEOF, "EOF"},
	}

	for _, e := range expected {
		token := l.GetNextToken()
		assert.Equal(t, e.Type, token.Type)
		assert.Equal(t, e.Value, token.Value)
	}
}

func  TestUnknownToken(t *testing.T) {
  mockData := ";~"
	l := NewLexer(strings.NewReader(mockData))
  token := l.GetNextToken()
  assert.Equal(t, token.Type, TokenInvalid)
}
