package parser

import (
	"math"
	"testing"

	"github.com/Jintumoni/vortex/errors"
	"github.com/Jintumoni/vortex/lexer"
	"github.com/Jintumoni/vortex/mocks"
	"github.com/Jintumoni/vortex/nodes"
	"github.com/stretchr/testify/assert"
)

func TestPropertyDef(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenString, Value: "string"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenInteger, Value: "int"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	properties, err := p.propertyDef()

	assert.NoError(t, err)
	assert.NotNil(t, properties)
	assert.Len(t, properties, 2)

	value, ok := properties[0].(*nodes.PropertyDefNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.TokenString, value.PropertyType.Type)
	assert.Equal(t, "name", value.PropertyName.Value)

	value, ok = properties[1].(*nodes.PropertyDefNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.TokenInteger, value.PropertyType.Type)
	assert.Equal(t, "age", value.PropertyName.Value)
}

// property_init: (DOT ID EQUAL factor)*
func TestPropertyInit(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "26"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "salary"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "100"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	properties, err := p.propertyInit()

	assert.NoError(t, err)
	assert.NotNil(t, properties)
	assert.Len(t, properties, 3)

	value, ok := properties[0].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "name", value.PropertyName.Value)
	assert.Equal(t, "John", value.PropertyValue.Value)

	value, ok = properties[1].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "age", value.PropertyName.Value)
	assert.Equal(t, "26", value.PropertyValue.Value)

	value, ok = properties[2].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "salary", value.PropertyName.Value)
	assert.Equal(t, "100", value.PropertyValue.Value)
}

// schema_def: SCHEMA ID LCB property_def RCB
func TestSchemaDef(t *testing.T) {
	mockLexer := new(mocks.MockLexer)
	// mockParser := new(mocks.MockParser)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenSchema, Value: "Schema"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()

	// mockParser.On("propertyDef").Return([]nodes.PropertyDefNode{
	// 	{PropertyName: "name", PropertyType: lexer.TokenString},
	// 	{PropertyName: "age", PropertyType: lexer.TokenInteger},
	// }).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenString, Value: "string"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenInteger, Value: "int"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	schemaNode, err := p.schemaDef()
	assert.NoError(t, err)

	value, ok := schemaNode.(*nodes.SchemaDefNode)
	assert.True(t, ok)
	assert.Equal(t, "Person", value.SchemaName.Value)
	assert.Len(t, value.Properties, 2)

	property1 := value.Properties[0].(*nodes.PropertyDefNode)
	assert.Equal(t, "name", property1.PropertyName.Value)
	assert.Equal(t, lexer.TokenString, property1.PropertyType.Type)

	property2 := value.Properties[1].(*nodes.PropertyDefNode)
	assert.Equal(t, "age", property2.PropertyName.Value)
	assert.Equal(t, lexer.TokenInteger, property2.PropertyType.Type)
}

// edge_def: EDGE EDGE_TYPE AS ID
func TestEdgeDef(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEdge, Value: "Edge"}).Once()
	// mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "OneWay"}).Once()
	// mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	edgeNode, err := p.edgeDef()
	assert.NoError(t, err)

	edge, ok := edgeNode.(*nodes.EdgeDefNode)

	assert.True(t, ok)
	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, nodes.OneWayEdge, edge.EdgeType)
}

// vertex_init: VERTEX ID ID LCB property_init RCB
func TestVertexInit(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenVertex, Value: "Vertex"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person1"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()

	// prperty_init
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "26"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "salary"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "100"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)

	vertexNode, err := p.vertexInit()
	assert.NoError(t, err)

	vertex, ok := vertexNode.(*nodes.VertexInitNode)

	assert.True(t, ok)
	assert.Equal(t, "Person", vertex.SchemaName.Value)
	assert.Equal(t, "Person1", vertex.VertexName.Value)

	properties := vertex.Properties

	assert.NotNil(t, properties)
	assert.Len(t, properties, 3)

	value, ok := properties[0].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "name", value.PropertyName.Value)
	assert.Equal(t, "John", value.PropertyValue.Value)

	value, ok = properties[1].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "age", value.PropertyName.Value)
	assert.Equal(t, "26", value.PropertyValue.Value)

	value, ok = properties[2].(*nodes.PropertyInitNode)
	assert.True(t, ok)
	assert.Equal(t, "salary", value.PropertyName.Value)
	assert.Equal(t, "100", value.PropertyValue.Value)
}

// relation_init: RELATION ID LCB ID ID RCB
func TestRelationInit(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRelation, Value: "Relation"}).Once()
	// mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person1"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Country1"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	// mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	relationNode, err := p.relationInit()
	assert.NoError(t, err)

	node, ok := relationNode.(*nodes.RelationInitNode)
	assert.True(t, ok)
	assert.Equal(t, "Person1", node.LeftVertex.Value)
	assert.Equal(t, "LivesIn", node.Relation.Value)
	assert.Equal(t, "Country1", node.RightVertex.Value)
}

func TestFactorInteger(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "1"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)

	assert.Equal(t, 1, factorNode.(*nodes.IntNode).Value)
}

func TestFactorString(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)

	assert.Equal(t, "John", factorNode.(*nodes.StringNode).Value)
}

func TestFactorPropertyIDWithoutAlias(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)
	property := factorNode.(*nodes.PropertyNode)

	assert.Equal(t, "name", property.PropertyName.Value)
	assert.Nil(t, property.Alias)
}

func TestFactorPropertyIDWithAlias(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)
	property := factorNode.(*nodes.PropertyNode)

	assert.Equal(t, "name", property.PropertyName.Value)
	assert.Equal(t, "Person", property.Alias.Value)
}

func TestFactorVertexWithAlias(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenAlias, Value: "as"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "A"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)
	vertex := factorNode.(*nodes.VertexTermNode)

	assert.Equal(t, "Person", vertex.Vertex.VertexName.Value)
}

func TestFactorVertexWithoutAlias(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	factorNode, err := p.factor()
	assert.NoError(t, err)
	vertex := factorNode.(*nodes.VertexTermNode)

	assert.Equal(t, "Person", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationDefaultBounds(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 1, edge.LowerBound.Value)
	assert.Equal(t, 1, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationConstantBound(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 2, edge.LowerBound.Value)
	assert.Equal(t, 2, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationNoBounds(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 0, edge.LowerBound.Value)
	assert.Equal(t, math.MaxInt, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationOnlyLowerBound(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 2, edge.LowerBound.Value)
	assert.Equal(t, math.MaxInt, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationOnlyUpperBound(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 0, edge.LowerBound.Value)
	assert.Equal(t, 2, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationBothLowerAndUpperBound(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "4"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "LivesIn"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Equal(t, "LivesIn", edge.EdgeName.Value)
	assert.Equal(t, 2, edge.LowerBound.Value)
	assert.Equal(t, 4, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationWithAnyEdge(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "4"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLRB, Value: "("}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRRB, Value: ")"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "India"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Nil(t, edge.EdgeName)
	assert.Equal(t, 2, edge.LowerBound.Value)
	assert.Equal(t, 4, edge.UpperBound.Value)
	assert.Equal(t, "India", vertex.Vertex.VertexName.Value)
}

func TestFactorRelationWithAnyEdgeAndAnyVertex(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLSB, Value: "["}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "2"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRange, Value: ".."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "4"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRSB, Value: "]"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLRB, Value: "("}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRRB, Value: ")"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLRB, Value: "("}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRRB, Value: ")"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()
	p := NewParser(mockLexer)
	relationNode, err := p.factor()
	assert.NoError(t, err)
	relation := relationNode.(*nodes.RelationNode)
	edge := relation.Edge.(*nodes.EdgeNode)
	vertex := relation.Vertex.(*nodes.VertexTermNode)

	assert.Nil(t, edge.EdgeName)
	assert.Equal(t, 2, edge.LowerBound.Value)
	assert.Equal(t, 4, edge.UpperBound.Value)
	assert.Nil(t, vertex.Vertex.VertexName)
}

func TestFactorVertexTermWithCondition(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenQuery, Value: "Query"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenAnd, Value: "and"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "20"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	queryNode, err := p.queryStatement()
	assert.NoError(t, err)

	query := queryNode.(*nodes.QueryStatementNode)

	condition := query.Expression.(*nodes.VertexTermNode).Conditions.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenAnd, condition.Operator.Type)

	left := condition.LeftChild.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenEqual, left.Operator.Type)
	assert.Equal(t, "name", left.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, "John", left.RightChild.(*nodes.StringNode).Value)

	right := condition.LeftChild.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenEqual, right.Operator.Type)
	assert.Equal(t, "name", right.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, "John", right.RightChild.(*nodes.StringNode).Value)
}

func TestFactorVertexTermWithNestedCondition(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenAnd, Value: "and"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLRB, Value: "("}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "age"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "20"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenOr, Value: "or"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "salary"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenGreaterThanEqual, Value: ">="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIntegerConstant, Value: "1000"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRRB, Value: ")"}).Once()

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	conditionNode, err := p.factor()
	assert.NoError(t, err)

	condition := conditionNode.(*nodes.VertexTermNode).Conditions.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenAnd, condition.Operator.Type)

	// Left Child
	left := condition.LeftChild.(*nodes.BinaryNode)

	assert.Equal(t, lexer.TokenEqual, left.Operator.Type)
	assert.Equal(t, "name", left.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, "John", left.RightChild.(*nodes.StringNode).Value)

	// Right Child
	right := condition.RightChild.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenOr, right.Operator.Type)

	l := right.LeftChild.(*nodes.BinaryNode)
	r := right.RightChild.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenEqual, l.Operator.Type)
	assert.Equal(t, "age", l.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, 20, l.RightChild.(*nodes.IntNode).Value)

	assert.Equal(t, lexer.TokenGreaterThanEqual, r.Operator.Type)
	assert.Equal(t, "salary", r.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, 1000, r.RightChild.(*nodes.IntNode).Value)
}

func TestFactorWithBuiltinFunction(t *testing.T) {
	mockLexer := new(mocks.MockLexer)

	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenFunction, Value: "Sum"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLRB, Value: "("}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenDot, Value: "."}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "name"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEqual, Value: "="}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenStringConstant, Value: "John"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRCB, Value: "}"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenRRB, Value: ")"}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

	p := NewParser(mockLexer)
	functionNode, err := p.factor()
	assert.NoError(t, err)

	function := functionNode.(*nodes.SumFuncNode)
	assert.Equal(t, function.FunctionName, nodes.SumFunc)

	args := function.Args
	assert.Len(t, args, 1)

	condition := args[0].(*nodes.VertexTermNode).Conditions.(*nodes.BinaryNode)
	assert.Equal(t, lexer.TokenEqual, condition.Operator.Type)
	assert.Equal(t, "name", condition.LeftChild.(*nodes.PropertyNode).PropertyName.Value)
	assert.Equal(t, "John", condition.RightChild.(*nodes.StringNode).Value)
}

func TestUnexpectedTokenErrorMessage(t *testing.T) {
  mockLexer := new(mocks.MockLexer)
  // Query Person as {
  mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenQuery, Value: "Query", Row: 0, Col: 0, Span: 5}).Once()
  mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenIdentifier, Value: "Person", Row: 0, Col: 7, Span: 6}).Once()
  mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenAlias, Value: "as", Row: 0, Col: 14, Span: 2}).Once()
  mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenLCB, Value: "{", Row: 0, Col: 18, Span: 1}).Once()
	mockLexer.On("GetNextToken").Return(&lexer.Token{Type: lexer.TokenEOF, Value: "EOF"}).Once()

  // There are tab characters in the string
  expectedErrMsg := `Error: Unexpected "{" found
1	|	Query  Person as  {
		                  ^--Did you mean "<identifier>"?
`
  mockLexer.On("GetSourceContext").Return("1	|	Query  Person as  {\n").Once()

  p := NewParser(mockLexer)
  node, err := p.Parse()

  assert.Nil(t, node)

  var expectedErr *errors.UnexpectedToken
  assert.ErrorAs(t, err, &expectedErr)
  assert.Equal(t, expectedErrMsg, err.Error())
}
