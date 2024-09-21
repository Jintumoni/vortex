package lexer

type TokenType int

const (
	TokenPlus TokenType = iota + 1
	TokenMinus
	TokenDivide
	TokenMultiply
	TokenLessThan
	TokenLessThanEqual
	TokenGreaterThan
	TokenGreaterThanEqual
	TokenEqual
	TokenNotEqual
	TokenOr
	TokenAnd
	TokenAlias
	TokenComma
	TokenDot
	TokenRange
	TokenLCB
	TokenRCB
	TokenLRB
	TokenRRB
	TokenLSB
	TokenRSB
	TokenInteger
	TokenIntegerConstant
	TokenString
	TokenStringConstant
	TokenLiteral
	TokenIdentifier
	TokenFunction
	TokenEOF
	TokenInvalid
	TokenSchema
	TokenVertex
	TokenEdge
	TokenRelation
	TokenQuery
)

func (t TokenType) String() string {
	switch t {
	case TokenPlus:
		return "+"
	case TokenMinus:
		return "-"
	case TokenDivide:
		return "/"
	case TokenMultiply:
		return "*"
	case TokenLessThan:
		return "<"
	case TokenLessThanEqual:
		return "<="
	case TokenGreaterThan:
		return ">"
	case TokenGreaterThanEqual:
		return ">="
	case TokenEqual:
		return "="
	case TokenNotEqual:
		return "!="
	case TokenOr:
		return "or"
	case TokenAnd:
		return "and"
	case TokenAlias:
		return "as"
	case TokenComma:
		return ","
	case TokenDot:
		return "."
	case TokenRange:
		return ".."
	case TokenLCB:
		return "{"
	case TokenRCB:
		return "}"
	case TokenLRB:
		return "("
	case TokenRRB:
		return ")"
	case TokenLSB:
		return "["
	case TokenRSB:
		return "]"
	case TokenInteger:
		return "int"
	case TokenIntegerConstant:
		return "<integer>"
	case TokenString:
		return "string"
	case TokenStringConstant:
		return "<string>"
	case TokenIdentifier:
		return "<identifier>"
	case TokenFunction:
		return "<function>"
	case TokenEOF:
		return "eof"
	case TokenSchema:
		return "Schema"
	case TokenVertex:
		return "Vertex"
	case TokenEdge:
		return "Edge"
	case TokenRelation:
		return "Relation"
	case TokenQuery:
		return "Query"
	default:
		return ""
	}
}

var ReservedKeywords = map[string]TokenType{
	"as":         TokenAlias,
	"and":        TokenAnd,
	"or":         TokenOr,
	"int":        TokenInteger,
	"string":     TokenString,
	"Sum":        TokenFunction,
	"Max":        TokenFunction,
	"Min":        TokenFunction,
	"StartsWith": TokenFunction,
	"Schema":     TokenSchema,
	"Vertex":     TokenVertex,
	"Relation":   TokenRelation,
	"Edge":       TokenEdge,
	"Query":      TokenQuery,
}

type Token struct {
	Type  TokenType
	Value string
	Row   int
	Col   int
	Span  int // stride of the token
}

func GetAllStatementTypes() []TokenType {
	return []TokenType{
		TokenSchema,
		TokenVertex,
		TokenEdge,
		TokenRelation,
		TokenQuery,
	}
}
