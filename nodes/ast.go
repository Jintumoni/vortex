package nodes

type EdgeType int

const (
  OneWayEdge EdgeType = iota
  TwoWayEdge
)

type FuncType int

const (
  SumFunc FuncType = iota
  MaxFunc
  MinFunc
  StartWithFunc
)

type ASTNode interface {
	// TODO: implement visitor class
	Accept(v Visitor)
}

type BuiltinFuncNode interface {
  Apply(args ...ASTNode) ASTNode
}

