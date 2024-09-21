package lexer

type TokenType int

const (
	TokenPlus TokenType = iota
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
