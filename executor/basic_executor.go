package executor

import (
	"errors"

	"github.com/Jintumoni/vortex/manager"
	"github.com/Jintumoni/vortex/nodes"
	"github.com/Jintumoni/vortex/parser"
	"github.com/Jintumoni/vortex/visitors"
	// "github.com/vmihailenco/msgpack/v5"
)

var (
	UnknownRootNode = errors.New("Unknown root node detected by parser")
)

type Executor struct {
	AppManager *manager.AppManager
	Parser     parser.ParserInterface
}

func NewExecutor(appManager *manager.AppManager, parser parser.ParserInterface) *Executor {
	return &Executor{AppManager: appManager, Parser: parser}
}

func (q *Executor) Execute() error {
	root, err := q.Parser.Parse()
	if err != nil {
		return err
	}

	visitor := visitors.NewVisualizer()

	for _, node := range root.(*nodes.ProgramStatementNode).Children {
		switch node.(type) {
		case *nodes.SchemaDefNode:
			q.AppManager.WriteSchema(node.(*nodes.SchemaDefNode))
		case *nodes.EdgeDefNode:
			q.AppManager.WriteEdge(node.(*nodes.EdgeDefNode))
		case *nodes.RelationInitNode:
			q.AppManager.WriteRelation(node.(*nodes.RelationInitNode))
		case *nodes.VertexInitNode:
			q.AppManager.WriteVertex(node.(*nodes.VertexInitNode))
		case *nodes.QueryStatementNode:
			root.Accept(visitor)
		default:
			return UnknownRootNode
		}
	}
	// data, err := json.Marshal(root)
	// fileio.WriteToFile(config.SchemaDefPath, bytes.NewReader(data))
	return nil
}
