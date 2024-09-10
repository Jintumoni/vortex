package nodes

func (node *ProgramStatementNode) Accept(visitor Visitor) {
	visitor.VisitProgramNode(node)
}

func (node *SchemaDefNode) Accept(visitor Visitor) {
	visitor.VisitSchemaDefNode(node)
}

func (node *PropertyDefNode) Accept(visitor Visitor) {
	visitor.VisitPropertyDefNode(node)
}

func (node *EdgeDefNode) Accept(visitor Visitor) {
	visitor.VisitEdgeDefNode(node)
}

func (node *PropertyInitNode) Accept(visitor Visitor) {
	visitor.VisitPropertyInitNode(node)
}

func (node *RelationInitNode) Accept(visitor Visitor) {
	visitor.VisitRelationInitNode(node)
}

func (node *VertexInitNode) Accept(visitor Visitor) {
	visitor.VisitVertexInitNode(node)
}

func (node *StringNode) Accept(visitor Visitor) {
	visitor.VisitStringNode(node)
}

func (node *IntNode) Accept(visitor Visitor) {
	visitor.VisitIntNode(node)
}

func (node *EdgeNode) Accept(visitor Visitor) {
  visitor.VisitEdgeNode(node)
}

func (node *PropertyNode) Accept(visitor Visitor) {
  visitor.VisitPropertyNode(node)
}

func (node *BinaryNode) Accept(visitor Visitor) {
  visitor.VisitBinaryNode(node)
}

func (node *VertexNode) Accept(visitor Visitor) {
  visitor.VisitVertexNode(node)
}

func (node *VertexTermNode) Accept(visitor Visitor) {
  visitor.VisitVertexTermNode(node)
}

func (node *RelationNode) Accept(visitor Visitor) {
  visitor.VisitRelationNode(node)
}

func (node *QueryStatementNode) Accept(visitor Visitor) {
  visitor.VisitQueryStatement(node)
}

func (node *SumFuncNode) Accept(visitor Visitor) {
  visitor.VisitSumFunc(node)
}
