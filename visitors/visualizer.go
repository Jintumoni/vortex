package visitors

import (
	"fmt"
	"strconv"
	"strings"

	// "strings"

	"github.com/Jintumoni/vortex/nodes"
)

type Visualizer struct {
	level       int
	indentation int
	lastChild   []int
}

func NewVisualizer() *Visualizer {
	return &Visualizer{level: 0, indentation: 4, lastChild: []int{1}}
}

func (v *Visualizer) print(s string) {
	for _, isLast := range v.lastChild[:len(v.lastChild)-1] {
		if isLast != 0 {
			fmt.Printf("┃%s", strings.Repeat(" ", v.indentation-1))
		} else {
			fmt.Printf(strings.Repeat(" ", v.indentation))
		}
	}
	fmt.Printf("┗━%s\n", s)
}

func (v *Visualizer) shiftRight(childCnt int) {
	v.lastChild[v.level]--
	v.level++
	v.lastChild = append(v.lastChild, childCnt)
}

func (v *Visualizer) shiftLeft() {
	v.lastChild = v.lastChild[:len(v.lastChild)-1]
	v.level--
}

func (v *Visualizer) VisitProgramNode(node *nodes.ProgramStatementNode) {
	v.print("ProgramStatement")
	v.shiftRight(len(node.Children))

	for _, c := range node.Children {
		c.Accept(v)
	}

	v.shiftLeft()
}

func (v *Visualizer) VisitIntNode(node *nodes.IntNode) {
	v.print(strconv.Itoa(node.Value))
}

func (v *Visualizer) VisitStringNode(node *nodes.StringNode) {
	v.print(node.Value)
}

func (v *Visualizer) VisitSchemaDefNode(node *nodes.SchemaDefNode) {
	v.print("SchemaDef")
	v.shiftRight(1)

	v.print(node.SchemaName.Value)
	v.shiftRight(len(node.Properties))

	for _, p := range node.Properties {
		p.Accept(v)
	}

	v.shiftLeft()
	v.shiftLeft()
}

func (v *Visualizer) VisitEdgeDefNode(node *nodes.EdgeDefNode) {
	v.print("EdgeDef")
	v.shiftRight(1)

	edgeType := "OneWay"
	if node.EdgeType == nodes.TwoWayEdge {
		edgeType = "TwoWay"
	}

	v.print(fmt.Sprintf("%s: %s", node.EdgeName.Value, edgeType))
	v.shiftLeft()
}

func (v *Visualizer) VisitRelationInitNode(node *nodes.RelationInitNode) {
	v.print("RelationInit")
	v.shiftRight(1)

	v.print(node.Relation.Value)
	v.shiftRight(2)

	v.print(node.LeftVertex.Value)
	v.print(node.RightVertex.Value)

	v.shiftLeft()
	v.shiftLeft()
}
func (v *Visualizer) VisitPropertyDefNode(node *nodes.PropertyDefNode) {
	v.print(fmt.Sprintf("%s %s", node.PropertyName.Value, node.PropertyType.Value))
}

func (v *Visualizer) VisitPropertyInitNode(node *nodes.PropertyInitNode) {
	v.print(fmt.Sprintf("%s %s", node.PropertyName.Value, node.PropertyValue.Value))
}

func (v *Visualizer) VisitVertexInitNode(node *nodes.VertexInitNode) {
	v.print("VertexInit")
	v.shiftRight(1)

	v.print(fmt.Sprintf("%s: %s", node.VertexName.Value, node.SchemaName.Value))
	v.shiftRight(len(node.Properties))

	for _, p := range node.Properties {
		v.lastChild[v.level]--
		p.Accept(v)
	}

	v.shiftLeft()
	v.shiftLeft()
}

func (v *Visualizer) VisitEdgeNode(node *nodes.EdgeNode) {
	v.print(fmt.Sprintf("%s %d..%d", node.EdgeName.Value, node.LowerBound.Value, node.UpperBound.Value))
}

func (v *Visualizer) VisitPropertyNode(node *nodes.PropertyNode) {
	if node.Alias != nil {
		v.print(fmt.Sprintf("%s.%s", node.Alias.Value, node.PropertyName.Value))
	} else {
		v.print(fmt.Sprintf(".%s", node.PropertyName.Value))
	}
}

func (v *Visualizer) VisitBinaryNode(node *nodes.BinaryNode) {
	v.print(node.Operator.Value)
	v.shiftRight(2)

	node.LeftChild.Accept(v)
	node.RightChild.Accept(v)

	v.shiftLeft()
}

func (v *Visualizer) VisitVertexNode(node *nodes.VertexNode) {
	if node.VertexName == nil {
		v.print("()")
	} else if node.Alias == nil {
		v.print(node.VertexName.Value)
	} else {
		v.print(fmt.Sprintf("%s: %s", node.VertexName.Value, node.Alias.Value))
	}
}

func (v *Visualizer) VisitVertexTermNode(node *nodes.VertexTermNode) {
	v.print("VertexTerm")
	v.shiftRight(1)

	node.Vertex.Accept(v)
	if node.Conditions != nil {
		node.Conditions.Accept(v)
	}
	v.shiftLeft()
}

func (v *Visualizer) VisitRelationNode(node *nodes.RelationNode) {
	v.print("Relation")
	v.shiftRight(1)

	node.Edge.Accept(v)
	node.Vertex.Accept(v)

	v.shiftLeft()
}

func (v *Visualizer) VisitQueryStatement(node *nodes.QueryStatementNode) {
	v.print("QueryStatement")
	v.shiftRight(1)

	node.Expression.Accept(v)

	v.shiftLeft()
}

func (v *Visualizer) VisitSumFunc(node *nodes.SumFuncNode) {
	v.print("Sum: BuiltinFunc")
	v.shiftRight(len(node.Args))
	for _, arg := range node.Args {
		arg.Accept(v)
	}
	v.shiftLeft()
}
