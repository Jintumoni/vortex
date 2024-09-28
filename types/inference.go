package types

import "github.com/Jintumoni/vortex/nodes"

func GetTypeInference(node nodes.ASTNode) {
	if isInteger(node) {
	}
}

func isInteger(node nodes.ASTNode) bool {
	_, ok := node.(*nodes.IntNode)
	if !ok {
		return false
	}
	return true
}

func isString(node nodes.ASTNode) bool {
	_, ok := node.(*nodes.StringNode)
	if !ok {
		return false
	}
	return true
}

func isList(node nodes.ASTNode) bool {
	_, ok := node.(*nodes.VertexTermNode)
	if ok {
		return true
	}
	_, ok = node.(*nodes.RelationNode)
	if ok {
		return true
	}
	return false
}
