package mocks

import (
	"github.com/Jintumoni/vortex/lexer"
	"github.com/stretchr/testify/mock"
)

type MockLexer struct {
	mock.Mock
}

func (m *MockLexer) GetNextToken() *lexer.Token {
	args := m.Called()
	return args.Get(0).(*lexer.Token)
}
