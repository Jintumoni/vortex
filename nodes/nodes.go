package nodes

import "github.com/Jintumoni/vortex/lexer"

type StringNode struct {
	Value string
}

type IntNode struct {
	Value int
}

type BoolNode struct {
	Value bool
}

type BinaryNode struct {
	LeftChild  ASTNode
	Operator   lexer.Token
	RightChild ASTNode
}

type ProgramStatementNode struct {
	Children []ASTNode
}

type QueryStatementNode struct {
	Expression ASTNode
}

type SumFuncNode struct {
	FunctionName FuncType
	Args         []ASTNode
}

type MaxFuncNode struct {
	FunctionName *StringNode
	Args         []ASTNode
}

type MinFuncNode struct {
	FunctionName *StringNode
	Args         []ASTNode
}

type StartsWithFuncNode struct {
	FunctionName *StringNode
	Args         []ASTNode
}

type VertexTermNode struct {
	Vertex     *VertexNode
	Conditions ASTNode
}

type VertexNode struct {
	VertexName *StringNode
	Alias      *StringNode
}

type RelationNode struct {
	Edge   ASTNode
	Vertex ASTNode
}

type ConditionNode struct {
	Child ASTNode
}

type SchemaDefNode struct {
	SchemaName *StringNode `json:"schema_name"`
	Properties []ASTNode   `json:"properties"`
}

type VertexInitNode struct {
	SchemaName *StringNode
	VertexName *StringNode
	Properties []ASTNode
}

type PropertyNode struct {
	PropertyName *StringNode
	Alias        *StringNode
}

type PropertyDefNode struct {
	PropertyName *StringNode `json:"property_name"`
	PropertyType lexer.Token `json:"property_type"`
}

type PropertyInitNode struct {
	PropertyName  *StringNode
	PropertyValue *StringNode
}

type EdgeNode struct {
	EdgeName   *StringNode
	LowerBound *IntNode
	UpperBound *IntNode
}

type EdgeDefNode struct {
	EdgeName *StringNode // Store the name of the edge
	EdgeType EdgeType    // Type of the edge (enum?)
}

type RelationInitNode struct {
	LeftVertex  *StringNode
	Relation    *StringNode
	RightVertex *StringNode
}
