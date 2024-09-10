package nodes

type Visitor interface {
	VisitProgramNode(node *ProgramStatementNode)
  VisitIntNode(node *IntNode)
  VisitStringNode(node *StringNode)
	VisitSchemaDefNode(node *SchemaDefNode)
	VisitEdgeDefNode(node *EdgeDefNode)
	VisitRelationInitNode(node *RelationInitNode)
	VisitPropertyDefNode(node *PropertyDefNode)
	VisitPropertyInitNode(node *PropertyInitNode)
	VisitVertexInitNode(node *VertexInitNode)
	VisitEdgeNode(node *EdgeNode)
	VisitPropertyNode(node *PropertyNode)
	VisitBinaryNode(node *BinaryNode)
	VisitVertexNode(node *VertexNode)
	VisitVertexTermNode(node *VertexTermNode)
	VisitRelationNode(node *RelationNode)
	VisitQueryStatement(node *QueryStatementNode)
  VisitSumFunc(node *SumFuncNode)
}
