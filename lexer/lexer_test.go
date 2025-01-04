package lexer

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdvance(t *testing.T) {
	mockData := "Hello"
	l := NewLexer(strings.NewReader(mockData))
	if l.Index != 0 {
		assert.Equal(t, 0, l.PeekIndex)
	}
	l.advance()
	if l.Index != 1 {
		assert.Equal(t, 1, l.PeekIndex)
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
	if l.PeekIndex != 0 {
		assert.Equal(t, 0, l.PeekIndex)
	}

	l.PeekIndex = strings.Index(mockData, " ")
	l.ignoreSpace()

	if l.PeekIndex != strings.Index(mockData, "World") {
		assert.Equal(t, 17, l.PeekIndex)
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
    #FriendsWith Person {
      #LivesIn Country{.name="India"}
      & .salary >= 100
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
		{TokenHash, "#"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenHash, "#"},
		{TokenIdentifier, "LivesIn"},
		{TokenIdentifier, "Country"},
		{TokenLCB, "{"},
		{TokenDot, "."},
		{TokenIdentifier, "name"},
		{TokenEqual, "="},
		{TokenStringConstant, "India"},
		{TokenRCB, "}"},
		{TokenAnd, "&"},
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
    #FriendsWith Person {
      #LivesIn [2..] Country{.name="India"}
      & .salary >= 100
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
		{TokenHash, "#"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenHash, "#"},
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
		{TokenAnd, "&"},
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
    #FriendsWith Person {
      #LivesIn [2..] Country{.name="India"}
      & .salary >= A.salary
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
		{TokenHash, "#"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenHash, "#"},
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
		{TokenAnd, "&"},
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
    Sum(#FriendsWith Person {
    #LivesIn [2..] Country{.name="India"}
      & .salary >= A.salary
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
		{TokenIdentifier, "Sum"},
		{TokenLRB, "("},
		{TokenHash, "#"},
		{TokenIdentifier, "FriendsWith"},
		{TokenIdentifier, "Person"},
		{TokenLCB, "{"},
		{TokenHash, "#"},
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
		{TokenAnd, "&"},
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

func TestUnknownToken(t *testing.T) {
	mockData := ";~"
	l := NewLexer(strings.NewReader(mockData))
	token := l.GetNextToken()
	assert.Equal(t, token.Type, TokenInvalid)
}

func TestAddToken(t *testing.T) {
	mockData := "+"
	l := NewLexer(strings.NewReader(mockData))
	l.GetNextToken()
	token := l.addSLToken(TokenPlus, "+")
	assert.Equal(t, TokenPlus, token.Type)
	assert.Equal(t, "+", token.Value)
	assert.Equal(t, 0, token.Row)
	assert.Equal(t, 0, token.Col)
	assert.Equal(t, 1, token.Span)
}

func TestMultilineToken(t *testing.T) {
	mockData := `"this is a sample
  multiline string"`
	l := NewLexer(strings.NewReader(mockData))
	token := l.GetNextToken()
	assert.Equal(t, TokenStringConstant, token.Type)
	assert.Equal(t, `this is a sample
  multiline string`, token.Value)
	assert.Equal(t, 0, token.Row)
	assert.Equal(t, 1, token.Col)
	assert.Equal(t, len(mockData)-2, token.Span)
}

func TestGetSourceContext(t *testing.T) {
	mockData := `Person as A {
    Sum(#FriendsWith Person {
      #LivesIn[2..] Country{.name="India"}
      & .salary >= A.salary`
	l := NewLexer(strings.NewReader(mockData))
	l.Row = 3
	source := l.GetSourceContext()

	// Tabs characters are present
	assert.Equal(t, `2	|	    Sum(#FriendsWith Person {
3	|	      #LivesIn[2..] Country{.name="India"}
4	|	      & .salary >= A.salary
`, source)
}

func TestNewLexer(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		r    io.Reader
		want *Lexer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLexer(tt.r)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewLexer() = %v, want %v", got, tt.want)
			}
		})
	}
}
