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

  Query Person as A {
    .name = "Hi"
    & Sum(#LivesIn [1..2] Any {
        #Within [..] Country {
            .name = "India"
            & #Within Continent
        }
    }, .age, .income) = 10
    & #FriendsWith Person {
	     #LivesIn Any {
            #Within Country {.name="USA"}
        }
    }
  }
  `

	lexer := lexer.NewLexer(strings.NewReader(input))
	parser := parser.NewParser(lexer)
	executor := executor.NewExecutor(appManager, parser)
	if err := executor.Execute(); err != nil {
		log.Fatal(err)
	}
}
