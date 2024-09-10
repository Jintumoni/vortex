package mocks

import (
	"github.com/Jintumoni/vortex/nodes"
	"github.com/stretchr/testify/mock"
)

type MockParser struct {
	mock.Mock
}

func (m *MockParser) PropertyDef() *nodes.ASTNode {
	args := m.Called()
	return args.Get(0).(*nodes.ASTNode)
}

func (m *MockParser) eat() {
	m.Called()
}
