package main

import (
	"log"
	"strings"

	"github.com/Jintumoni/vortex/executor"
	"github.com/Jintumoni/vortex/lexer"
	"github.com/Jintumoni/vortex/manager"
	"github.com/Jintumoni/vortex/parser"
)

func main() {
	appManager := manager.NewAppManager()

	input := `Schema Person {
    name string
    age int
  }

  Vertex Harry Person {
    .name = "Harry"
    .age = 1
  }

  Edge LivesIn OneWay

  Relation LivesIn {
    Harry London
  }

  Query Sum(Person as P {
      []FriendsWith Person {
        2 + .income + 100 > P.income + 1
        and .name = "John" and (.age = 1 or P.salary > 100)
        and []FriendsWith P
      }
  })
  `

	lexer := lexer.NewLexer(strings.NewReader(input))
	parser := parser.NewParser(lexer)
	executor := executor.NewExecutor(appManager, parser)
	if err := executor.Execute(); err != nil {
		log.Fatal(err)
	}
}
