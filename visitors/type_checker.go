package visitors

import (
	"github.com/Jintumoni/vortex/nodes"
)

type ReturnType int

const (
	IntType ReturnType = iota
	FloatType
	BoolType
	StringType
	ListType
)

type TypeChecker struct {
}

func NewTypeChecker() *TypeChecker {
  return &TypeChecker{}
}

func (v *TypeChecker) VisitProgramNode(node *nodes.ProgramStatementNode) {
}

func (v *TypeChecker) VisitIntNode(node *nodes.IntNode) {
}

func (v *TypeChecker) VisitStringNode(node *nodes.StringNode) {
}

func (v *TypeChecker) VisitSchemaDefNode(node *nodes.SchemaDefNode) {
}

func (v *TypeChecker) VisitEdgeDefNode(node *nodes.EdgeDefNode) {
}

func (v *TypeChecker) VisitRelationInitNode(node *nodes.RelationInitNode) {
}

func (v *TypeChecker) VisitPropertyDefNode(node *nodes.PropertyDefNode) {
}

func (v *TypeChecker) VisitPropertyInitNode(node *nodes.PropertyInitNode) {
}

func (v *TypeChecker) VisitVertexInitNode(node *nodes.VertexInitNode) {
}

func (v *TypeChecker) VisitEdgeNode(node *nodes.EdgeNode) {
}

func (v *TypeChecker) VisitPropertyNode(node *nodes.PropertyNode) {
}

func (v *TypeChecker) VisitBinaryNode(node *nodes.BinaryNode) {
}

func (v *TypeChecker) VisitVertexNode(node *nodes.VertexNode) {
}

func (v *TypeChecker) VisitVertexTermNode(node *nodes.VertexTermNode) {
}

func (v *TypeChecker) VisitRelationNode(node *nodes.RelationNode) {
}

func (v *TypeChecker) VisitQueryStatement(node *nodes.QueryStatementNode) {
}

func (v *TypeChecker) VisitSumFunc(node *nodes.SumFuncNode) {
}
