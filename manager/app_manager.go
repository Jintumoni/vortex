package manager

import (
	"errors"

	"github.com/Jintumoni/vortex/nodes"
)

var (
	SchemaAlreadyExist   = errors.New("Schema already exist")
	SchemaDoesNotExist   = errors.New("Schema missing")
	VertexAlreadyExist   = errors.New("Vertex already exist")
	VertexDoesNotExist   = errors.New("Vertex missing")
	RelationAlreadyExist = errors.New("Relation already exist")
	RelationDoesNotExist = errors.New("Relation missing")
	EdgeAlreadyExist     = errors.New("Edge already exist")
	EdgeDoesNotExist     = errors.New("Edge missing")
)

type NodeRelationPair struct {
	Vertex   *nodes.VertexInitNode
	Relation *nodes.EdgeDefNode
}
type AdjacencyList map[*nodes.VertexInitNode][]*NodeRelationPair

type AppManager struct {
	schemaStore   map[string]*nodes.SchemaDefNode
	vertexStore   map[string]*nodes.VertexInitNode
	edgeStore     map[string]*nodes.EdgeDefNode
	relationStore map[string]*nodes.RelationInitNode
	graphStore    AdjacencyList
}

func NewAppManager() *AppManager {
	return &AppManager{
		schemaStore:   make(map[string]*nodes.SchemaDefNode),
		vertexStore:   make(map[string]*nodes.VertexInitNode),
		edgeStore:     make(map[string]*nodes.EdgeDefNode),
		relationStore: make(map[string]*nodes.RelationInitNode),
	}
}

func (a *AppManager) WriteSchema(s *nodes.SchemaDefNode) error {
	_, ok := a.schemaStore[s.SchemaName.Value]
	if ok {
		return SchemaAlreadyExist
	}

	a.schemaStore[s.SchemaName.Value] = s
	return nil
}

func (a *AppManager) ReadSchema(s string) (*nodes.SchemaDefNode, error) {
	schemaNode, ok := a.schemaStore[s]
	if !ok {
		return nil, SchemaDoesNotExist
	}

	return schemaNode, nil
}

func (a *AppManager) WriteVertex(v *nodes.VertexInitNode) error {
	_, ok := a.vertexStore[v.VertexName.Value]
	if ok {
		return VertexAlreadyExist
	}

	a.vertexStore[v.VertexName.Value] = v
	return nil
}

func (a *AppManager) ReadVertex(s string) (*nodes.VertexInitNode, error) {
	vertexNode, ok := a.vertexStore[s]
	if !ok {
		return nil, VertexDoesNotExist
	}

	return vertexNode, nil
}

func (a *AppManager) ReadRelation(s string) (*nodes.RelationInitNode, error) {
	relationNode, ok := a.relationStore[s]
	if !ok {
		return nil, RelationDoesNotExist
	}

	return relationNode, nil
}

func (a *AppManager) WriteRelation(r *nodes.RelationInitNode) error {
	_, ok := a.relationStore[r.Relation.Value]
	if ok {
		return RelationAlreadyExist
	}
	return a.JoinVertex(r)
}

func (a *AppManager) ReadEdge(s string) (*nodes.EdgeDefNode, error) {
	edgeNode, ok := a.edgeStore[s]
	if !ok {
		return nil, EdgeDoesNotExist
	}

	return edgeNode, nil
}

func (a *AppManager) WriteEdge(r *nodes.EdgeDefNode) error {
	_, ok := a.edgeStore[r.EdgeName.Value]
	if ok {
		return EdgeAlreadyExist
	}

	return nil
}

func (a *AppManager) JoinVertex(r *nodes.RelationInitNode) error {
	leftVertex, err := a.ReadVertex(r.LeftVertex.Value)
	if err != nil {
		return err
	}
	rightVertex, err := a.ReadVertex(r.RightVertex.Value)
	if err != nil {
		return err
	}
	relation, err := a.ReadEdge(r.Relation.Value)
	if err != nil {
		return err
	}

	a.graphStore[leftVertex] = append(a.graphStore[leftVertex], &NodeRelationPair{Vertex: rightVertex, Relation: relation})
	if relation.EdgeType == nodes.TwoWayEdge {
		a.graphStore[rightVertex] = append(a.graphStore[rightVertex], &NodeRelationPair{Vertex: leftVertex, Relation: relation})
	}
	return nil
}
