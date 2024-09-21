package errors

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Jintumoni/vortex/lexer"
	"github.com/Jintumoni/vortex/nodes"
	"github.com/fatih/color"
)

type UnexpectedToken struct {
	SourceContext   string
	ActualToken     *lexer.Token
	ExpectedToken   lexer.TokenType
	SuggestedTokens []lexer.TokenType
}

func (e *UnexpectedToken) Error() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(color.RedString(fmt.Sprintf("Error: Unexpected \"%s\" found\n", e.ActualToken.Value)))

	buffer.WriteString(e.SourceContext)

	buffer.WriteString(strings.Repeat("\t", 2))
	buffer.WriteString(strings.Repeat(" ", e.ActualToken.Col))
	// buffer.WriteString("\033[31m")
	buffer.WriteString(color.BlueString(strings.Repeat("^", e.ActualToken.Span)))
	buffer.WriteString(color.BlueString("--"))

	if e.SuggestedTokens != nil {
		buffer.WriteString(color.BlueString(fmt.Sprintf("Expected one of: ")))
		for i, t := range e.SuggestedTokens {
			if i > 0 {
				buffer.WriteString(color.BlueString(", "))
			}
			buffer.WriteString(color.BlueString(fmt.Sprintf("\"%s\"", t)))
		}
	} else if e.ExpectedToken != 0 {
		buffer.WriteString(color.BlueString(fmt.Sprintf("Did you mean \"%s\"?", e.ExpectedToken)))
	} else {
		buffer.WriteString(color.BlueString("here"))
	}
	// buffer.WriteString("\033[0m")
	buffer.WriteString("\n")

	return buffer.String()
}

type UnknownEdgeType struct {
	SourceContext string
	ActualToken   *lexer.Token
}

func (e *UnknownEdgeType) Error() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(color.RedString(fmt.Sprintf("Error: Unknown \"%s\" found\n", e.ActualToken.Value)))

	buffer.WriteString(e.SourceContext)

	buffer.WriteString(strings.Repeat("\t", 2))
	buffer.WriteString(strings.Repeat(" ", e.ActualToken.Col))

	buffer.WriteString(color.BlueString(strings.Repeat("^", e.ActualToken.Span)))
	buffer.WriteString(color.BlueString("--"))

	buffer.WriteString(color.BlueString(fmt.Sprintf("Expected one of: ")))
	for i, t := range nodes.GetAllEdgeTypes() {
		if i > 0 {
			buffer.WriteString(color.BlueString(", "))
		}
		buffer.WriteString(color.BlueString(fmt.Sprintf("\"%s\"", t)))
	}
	buffer.WriteString("\n")

	return buffer.String()
}

type UnknownStatement struct {
	SourceContext string
	ActualToken   *lexer.Token
}

func (e *UnknownStatement) Error() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(color.RedString(fmt.Sprintf("Error: Unknown \"%s\" found\n", e.ActualToken.Value)))

	buffer.WriteString(e.SourceContext)

	buffer.WriteString(strings.Repeat("\t", 2))
	buffer.WriteString(strings.Repeat(" ", e.ActualToken.Col))

	buffer.WriteString(color.BlueString(strings.Repeat("^", e.ActualToken.Span)))
	buffer.WriteString(color.BlueString("--"))

	buffer.WriteString(color.BlueString(fmt.Sprintf("Expected one of: ")))
	for i, t := range lexer.GetAllStatementTypes() {
		if i > 0 {
			buffer.WriteString(color.BlueString(", "))
		}
		buffer.WriteString(color.BlueString(fmt.Sprintf("\"%s\"", t)))
	}
	buffer.WriteString("\n")

	return buffer.String()
}

type UnknownBuiltinFunc struct {
	SourceContext string
	ActualToken   *lexer.Token
}

func (e *UnknownBuiltinFunc) Error() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString(color.RedString(fmt.Sprintf("Error: Unknown \"%s\" found\n", e.ActualToken.Value)))

	buffer.WriteString(e.SourceContext)

	buffer.WriteString(strings.Repeat("\t", 2))
	buffer.WriteString(strings.Repeat(" ", e.ActualToken.Col))

	buffer.WriteString(color.BlueString(strings.Repeat("^", e.ActualToken.Span)))
	buffer.WriteString(color.BlueString("--"))

	buffer.WriteString(color.BlueString(fmt.Sprintf("Expected one of: ")))
	for i, t := range nodes.GetAllFuncTypes() {
		if i > 0 {
			buffer.WriteString(color.BlueString(", "))
		}
		buffer.WriteString(color.BlueString(fmt.Sprintf("\"%s\"", t)))
	}
	buffer.WriteString("\n")

	return buffer.String()
}
