package nodes

type ASTNode interface {
	// TODO: implement visitor class
	Accept(v Visitor)
}

type BuiltinFuncNode interface {
	Apply(args ...ASTNode) ASTNode
}
