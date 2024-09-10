package visitors

import (
	"bytes"
	// "encoding/json"
	// "log"

	// "github.com/Jintumoni/vortex/nodes"
)

type DDLRunner struct {
  buffer *bytes.Buffer
}

// func (t *DDLRunner) VisitSchemaDefNode(node *nodes.SchemaDefNode) {
//   schemaName := node.SchemaName
//   log.Println(schemaName)
//   for _, property := range node.Properties {
//     property.Accept(t)
//   }

//   data, err := json.Marshal(node)
//   if err != nil {
//     log.Fatal(err)
//   }
//   log.Println(string(data))
// }

// func (t *DDLRunner) VisitEdgeDefNode(node *nodes.EdgeDefNode) {


// }

// func (t *DDLRunner) VisitRelationInitNode(node *nodes.RelationInitNode) {

// }

// func (t *DDLRunner) VisitPropertyDefNode(node *nodes.PropertyDefNode) {
//   // propertyName := node.PropertyName
//   // log.Println("Property: ", propertyName)
//   // token := node.PropertyType
//   // log.Println("Type: ", token)
// }

// func (t *DDLRunner) VisitPropertyInitNode(node *nodes.PropertyInitNode) {

// }

// func (t *DDLRunner) VisitVertexInitNode(node *nodes.VertexInitNode) {

// }
