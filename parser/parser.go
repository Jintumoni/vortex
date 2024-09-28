package parser

import (
	"math"
	"strconv"

	"github.com/Jintumoni/vortex/errors"

	"github.com/Jintumoni/vortex/lexer"
	"github.com/Jintumoni/vortex/nodes"
)

type ParserInterface interface {
	// eat(tokenType lexer.TokenType) error
	// edgeDef() nodes.ASTNode
	// propertyDef() []nodes.ASTNode
	// propertyInit() []nodes.ASTNode
	// relationInit() nodes.ASTNode
	// schemaDef() nodes.ASTNode
	// vertexInit() nodes.ASTNode
	Parse() (nodes.ASTNode, error)
}

type Parser struct {
	Lexer        lexer.LexerInterface
	CurrentToken *lexer.Token
}

func NewParser(lexer lexer.LexerInterface) *Parser {
	return &Parser{
		Lexer:        lexer,
		CurrentToken: lexer.GetNextToken(),
	}
}

func (p *Parser) eat(tokenType lexer.TokenType) error {
	if tokenType != p.CurrentToken.Type {
		return &errors.UnexpectedToken{
			SourceContext: p.Lexer.GetSourceContext(),
			ActualToken:   p.CurrentToken,
			ExpectedToken: tokenType,
		}
	}
	p.CurrentToken = p.Lexer.GetNextToken()
	return nil
}

// edge_def: EDGE ID EDGE_TYPE
func (p *Parser) edgeDef() (nodes.ASTNode, error) {
	// EDGE
	if err := p.eat(lexer.TokenEdge); err != nil {
		return new(nodes.EdgeDefNode), err
	}

	rightToken := p.CurrentToken
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.EdgeDefNode), err
	} // ID

	var edgeType nodes.EdgeType
	leftToken := p.CurrentToken
	if leftToken.Value == "OneWay" {
		edgeType = nodes.OneWayEdge
	} else if leftToken.Value == "TwoWay" {
		edgeType = nodes.TwoWayEdge
	} else {
		return new(nodes.EdgeDefNode), &errors.UnknownEdgeType{
			SourceContext: p.Lexer.GetSourceContext(),
			ActualToken:   p.CurrentToken,
		}
	}
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.EdgeDefNode), err
	} // EDGE_TYPE

	// if err := p.eat(lexer.TokenAlias); err != nil {
	// 	return new(nodes.EdgeDefNode), err
	// } // AS

	return &nodes.EdgeDefNode{
		EdgeName: &nodes.StringNode{Value: rightToken.Value},
		EdgeType: edgeType,
	}, nil
}

// schema_def: SCHEMA ID LCB property_def RCB
func (p *Parser) schemaDef() (nodes.ASTNode, error) {
	if err := p.eat(lexer.TokenSchema); err != nil {
		return new(nodes.SchemaDefNode), err
	} // SCHEMA

	schemaName := p.CurrentToken.Value
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.SchemaDefNode), err
	} // ID
	if err := p.eat(lexer.TokenLCB); err != nil {
		return new(nodes.SchemaDefNode), err
	} // LCB

	properties, err := p.propertyDef() // (property_def)*
	if err != nil {
		return new(nodes.SchemaDefNode), err
	}
	if err := p.eat(lexer.TokenRCB); err != nil {
		return new(nodes.SchemaDefNode), err
	} // RCB

	return &nodes.SchemaDefNode{
		SchemaName: &nodes.StringNode{Value: schemaName},
		Properties: properties,
	}, nil
}

// property_def: (ID TYPE)*
func (p *Parser) propertyDef() ([]nodes.ASTNode, error) {
	var properties []nodes.ASTNode

	for p.CurrentToken.Type == lexer.TokenIdentifier {
		propertyName := p.CurrentToken.Value
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}

		property := p.CurrentToken
		if property.Type == lexer.TokenInteger {
			if err := p.eat(lexer.TokenInteger); err != nil {
				return nil, err
			}
		} else if property.Type == lexer.TokenString {
			if err := p.eat(lexer.TokenString); err != nil {
				return nil, err
			}
		} else {
			return nil, &errors.UnexpectedToken{
				SourceContext:   p.Lexer.GetSourceContext(),
				ActualToken:     p.CurrentToken,
				SuggestedTokens: []lexer.TokenType{lexer.TokenInteger, lexer.TokenString},
			}

		}

		properties = append(properties, &nodes.PropertyDefNode{
			PropertyName: &nodes.StringNode{Value: propertyName},
			PropertyType: *property,
		})

	}
	return properties, nil
}

// property_init: (DOT ID EQUAL factor)*
func (p *Parser) propertyInit() ([]nodes.ASTNode, error) {
	var arguments []nodes.ASTNode

	for p.CurrentToken.Type == lexer.TokenDot {
		if err := p.eat(lexer.TokenDot); err != nil {
			return nil, err
		}
		propertyName := p.CurrentToken.Value
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}
		if err := p.eat(lexer.TokenEqual); err != nil {
			return nil, err
		}

		literalToken := p.CurrentToken
		switch literalToken.Type {
		case lexer.TokenStringConstant:
			if err := p.eat(lexer.TokenStringConstant); err != nil {
				return nil, err
			}
		case lexer.TokenIntegerConstant:
			if err := p.eat(lexer.TokenIntegerConstant); err != nil {
				return nil, err
			}
		default:
			return nil, &errors.UnexpectedToken{
				SourceContext:   p.Lexer.GetSourceContext(),
				ActualToken:     p.CurrentToken,
				SuggestedTokens: []lexer.TokenType{lexer.TokenIntegerConstant, lexer.TokenStringConstant},
			}
		}

		arguments = append(arguments, &nodes.PropertyInitNode{
			PropertyName:  &nodes.StringNode{Value: propertyName},
			PropertyValue: &nodes.StringNode{Value: literalToken.Value},
		})
	}
	return arguments, nil
}

// vertex_init: VERTEX ID ID LCB property_init RCB
func (p *Parser) vertexInit() (nodes.ASTNode, error) {
	if err := p.eat(lexer.TokenVertex); err != nil {
		return nil, err
	}

	vertexName := p.CurrentToken.Value
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return nil, err
	}

	schemaName := p.CurrentToken.Value
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return nil, err
	}

	// if err := p.eat(lexer.TokenAlias); err != nil {
	// 	return nil, err
	// }

	if err := p.eat(lexer.TokenLCB); err != nil {
		return nil, err
	}

	properties, err := p.propertyInit()
	if err != nil {
		return nil, err
	}

	if err := p.eat(lexer.TokenRCB); err != nil {
		return nil, err
	}

	return &nodes.VertexInitNode{
		SchemaName: &nodes.StringNode{Value: schemaName},
		VertexName: &nodes.StringNode{Value: vertexName},
		Properties: properties,
	}, nil
}

// relation_init: RELATION ID LCB ID ID RCB
func (p *Parser) relationInit() (nodes.ASTNode, error) {
	if err := p.eat(lexer.TokenRelation); err != nil {
		return new(nodes.RelationInitNode), err
	}

	relation := p.CurrentToken
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.RelationInitNode), err
	}

	if err := p.eat(lexer.TokenLCB); err != nil {
		return nil, err
	}

	leftVertex := p.CurrentToken
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.RelationInitNode), err
	}

	rightVertex := p.CurrentToken
	if err := p.eat(lexer.TokenIdentifier); err != nil {
		return new(nodes.RelationInitNode), err
	}

	if err := p.eat(lexer.TokenRCB); err != nil {
		return nil, err
	}

	return &nodes.RelationInitNode{
		LeftVertex:  &nodes.StringNode{Value: leftVertex.Value},
		Relation:    &nodes.StringNode{Value: relation.Value},
		RightVertex: &nodes.StringNode{Value: rightVertex.Value},
	}, nil
}

func (p *Parser) programStatement() (nodes.ASTNode, error) {
	var programNodes []nodes.ASTNode
	for p.CurrentToken.Type != lexer.TokenEOF {
		switch p.CurrentToken.Type {
		case lexer.TokenSchema:
			schemaNode, err := p.schemaDef()
			if err != nil {
				return nil, err
			}
			programNodes = append(programNodes, schemaNode)
		case lexer.TokenEdge:
			edgeNode, err := p.edgeDef()
			if err != nil {
				return nil, err
			}
			programNodes = append(programNodes, edgeNode)
		case lexer.TokenVertex:
			vertexNode, err := p.vertexInit()
			if err != nil {
				return nil, err
			}
			programNodes = append(programNodes, vertexNode)
		case lexer.TokenRelation:
			relationNode, err := p.relationInit()
			if err != nil {
				return nil, err
			}
			programNodes = append(programNodes, relationNode)
		case lexer.TokenQuery:
			queryNode, err := p.queryStatement()
			if err != nil {
				return nil, err
			}
			programNodes = append(programNodes, queryNode)
		default:
			return nil, &errors.UnknownStatement{SourceContext: p.Lexer.GetSourceContext(), ActualToken: p.CurrentToken}
		}
	}
	return &nodes.ProgramStatementNode{Children: programNodes}, nil
}

// clause:
//
//	operation (operator operation)*
func (p *Parser) clause() (nodes.ASTNode, error) {
	operationLeft, err := p.operation()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == lexer.TokenLessThan ||
		p.CurrentToken.Type == lexer.TokenGreaterThan ||
		p.CurrentToken.Type == lexer.TokenLessThanEqual ||
		p.CurrentToken.Type == lexer.TokenGreaterThanEqual ||
		p.CurrentToken.Type == lexer.TokenEqual ||
		p.CurrentToken.Type == lexer.TokenNotEqual {

		operator := p.CurrentToken
		if err := p.eat(p.CurrentToken.Type); err != nil {
			return nil, err
		}
		operationRight, err := p.operation()
		if err != nil {
			return nil, err
		}

		operationLeft = &nodes.BinaryNode{LeftChild: operationLeft, Operator: *operator, RightChild: operationRight}
	}

	return operationLeft, nil
}

// expression:
//
//	clause ((AND | OR) clause)*
func (p *Parser) expression() (nodes.ASTNode, error) {
	leftClause, err := p.clause()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == lexer.TokenAnd || p.CurrentToken.Type == lexer.TokenOr {
		operator := p.CurrentToken
		if err := p.eat(p.CurrentToken.Type); err != nil {
			return nil, err
		}
		rightClause, err := p.clause()
		if err != nil {
			return nil, err
		}

		leftClause = &nodes.BinaryNode{LeftChild: leftClause, Operator: *operator, RightChild: rightClause}
	}

	return leftClause, nil
}

// query_statement:
//
//	Query expression
func (p *Parser) queryStatement() (nodes.ASTNode, error) {
	if err := p.eat(lexer.TokenQuery); err != nil {
		return nil, err
	}
	// if err := p.eat(lexer.TokenLCB); err != nil {
	// 	return nil, err
	// }
	expression, err := p.expression()
	if err != nil {
		return nil, err
	}
	// if err := p.eat(lexer.TokenRCB); err != nil {
	// 	return nil, err
	// }

	return &nodes.QueryStatementNode{Expression: expression}, nil
}

// factor:
//
//	(INT | STRING | BOOL)
//	| property_id
//	| vertex_term
//	| relation_term vertex_term
//	| builtin_func LCB clause RCB
//	| LRB expression RRB

func (p *Parser) integer() (nodes.ASTNode, error) {
	number, err := strconv.Atoi(p.CurrentToken.Value)
	if err != nil {
		return nil, err
	}
	if err := p.eat(lexer.TokenIntegerConstant); err != nil {
		return nil, err
	}
	return &nodes.IntNode{Value: number}, nil
}

func (p *Parser) string() (nodes.ASTNode, error) {
	str := p.CurrentToken.Value
	if err := p.eat(lexer.TokenStringConstant); err != nil {
		return nil, err
	}
	return &nodes.StringNode{Value: str}, nil
}

func (p *Parser) builtinFunc() (nodes.ASTNode, error) {
	if p.CurrentToken.Type != lexer.TokenFunction {
		return nil, &errors.UnexpectedToken{
			SourceContext: p.Lexer.GetSourceContext(),
			ActualToken:   p.CurrentToken,
			ExpectedToken: lexer.TokenFunction,
		}
	}

	function := p.CurrentToken

	if err := p.eat(lexer.TokenFunction); err != nil {
		return nil, err
	}

	switch function.Value {
	case "Sum":
		expression, err := p.expression()
		if err != nil {
			return nil, err
		}
		args := []nodes.ASTNode{expression}

		for p.CurrentToken.Type == lexer.TokenComma {
			args = append(args, &nodes.StringNode{Value: p.CurrentToken.Value})
		}
		return &nodes.SumFuncNode{FunctionName: nodes.SumFunc, Args: args}, nil
	default:
		return nil, &errors.UnknownBuiltinFunc{SourceContext: p.Lexer.GetSourceContext(), ActualToken: p.CurrentToken}
	}
}

func (p *Parser) factor() (nodes.ASTNode, error) {
	// LRB expression RRB
	if p.CurrentToken.Type == lexer.TokenLRB {
		// LRB
		if err := p.eat(lexer.TokenLRB); err != nil {
			return nil, err
		}

		expression, err := p.expression()
		if err != nil {
			return nil, err
		}

		// RRB
		if err := p.eat(lexer.TokenRRB); err != nil {
			return nil, err
		}

		return expression, nil
	}

	// builtin_func
	if p.CurrentToken.Type == lexer.TokenFunction {
		return p.builtinFunc()
	}

	// relation_term vertex_term
	if p.CurrentToken.Type == lexer.TokenLSB {
		relationTerm, err := p.relationTerm()
		if err != nil {
			return nil, err
		}
		vertexTerm, err := p.vertexTerm()
		if err != nil {
			return nil, err
		}

		return &nodes.RelationNode{Edge: relationTerm, Vertex: vertexTerm}, nil
	}

	// INT
	if p.CurrentToken.Type == lexer.TokenIntegerConstant {
		return p.integer()
	}

	// STRING
	if p.CurrentToken.Type == lexer.TokenStringConstant {
		return p.string()
	}

	// property_id (eg: .name)
	if p.CurrentToken.Type == lexer.TokenDot {
		if err := p.eat(lexer.TokenDot); err != nil {
			return nil, err
		}
		property := p.CurrentToken.Value
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}

		return &nodes.PropertyNode{
			PropertyName: &nodes.StringNode{Value: property}, Alias: nil,
		}, nil
	}

	// property_id (eg: A.name)
	// vertex_term  (eg: Person)
	if p.CurrentToken.Type == lexer.TokenIdentifier {
		id := p.CurrentToken.Value
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}

		if p.CurrentToken.Type == lexer.TokenDot {
			// property_id (eg: A.name)
			if err := p.eat(lexer.TokenDot); err != nil {
				return nil, err
			}
			property := p.CurrentToken.Value
			if err := p.eat(lexer.TokenIdentifier); err != nil {
				return nil, err
			}

			return &nodes.PropertyNode{
				PropertyName: &nodes.StringNode{Value: property}, Alias: &nodes.StringNode{Value: id},
			}, nil
		}

		alias := new(nodes.StringNode)
		if p.CurrentToken.Type == lexer.TokenAlias {
			// vertex_term (eg: Person)

			// vertex: ID (as ID)*
			if err := p.eat(lexer.TokenAlias); err != nil {
				return nil, err
			}

			alias.Value = p.CurrentToken.Value
			if err := p.eat(lexer.TokenIdentifier); err != nil {
				return nil, err
			}
		}

		vertex := &nodes.VertexNode{
			VertexName: &nodes.StringNode{Value: id},
			Alias:      alias,
		}

		// (LCB expression RCB)?
		if p.CurrentToken.Type == lexer.TokenLCB {
			if err := p.eat(lexer.TokenLCB); err != nil {
				return nil, err
			}
			condition, err := p.expression()
			if err != nil {
				return nil, err
			}
			if err := p.eat(lexer.TokenRCB); err != nil {
				return nil, err
			}
			return &nodes.VertexTermNode{Vertex: vertex, Conditions: condition}, nil
		}
		return &nodes.VertexTermNode{Vertex: vertex, Conditions: nil}, nil
	}

	// vertex_term  (Unit)
	if p.CurrentToken.Type == lexer.TokenLRB {
		return p.vertexTerm()
	}

	return nil, &errors.UnexpectedToken{
		SourceContext: p.Lexer.GetSourceContext(),
		ActualToken:   p.CurrentToken,
		SuggestedTokens: []lexer.TokenType{
			lexer.TokenLRB,
			lexer.TokenFunction,
			lexer.TokenEdge,
			lexer.TokenIntegerConstant,
			lexer.TokenStringConstant,
			lexer.TokenDot,
			lexer.TokenIdentifier,
		},
	}
}

// term:
//
//	factor ((MUL | DIV) factor)*
func (p *Parser) term() (nodes.ASTNode, error) {
	factorLeft, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == lexer.TokenMultiply || p.CurrentToken.Type == lexer.TokenDivide {
		operator := p.CurrentToken
		if err := p.eat(p.CurrentToken.Type); err != nil {
			return nil, err
		}

		factorRight, err := p.factor()
		if err != nil {
			return nil, err
		}

		factorLeft = &nodes.BinaryNode{LeftChild: factorLeft, Operator: *operator, RightChild: factorRight}
	}

	return factorLeft, nil
}

// operation:
//
//	term ((PLUS | MINUS) term)*
func (p *Parser) operation() (nodes.ASTNode, error) {
	termLeft, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.CurrentToken.Type == lexer.TokenPlus || p.CurrentToken.Type == lexer.TokenMinus {
		operator := p.CurrentToken
		if err := p.eat(p.CurrentToken.Type); err != nil {
			return nil, err
		}

		termRight, err := p.term()
		if err != nil {
			return nil, err
		}

		termLeft = &nodes.BinaryNode{LeftChild: termLeft, Operator: *operator, RightChild: termRight}
	}

	return termLeft, nil
}

// relation_term: (LSB (integer | (integer? DOT DOT integer?))? RSB) relation
// relation: ID | Unit
// Unit: LRB RRB
func (p *Parser) relationTerm() (nodes.ASTNode, error) {
	relation := new(nodes.EdgeNode)
	// default Upper/Lower bounds
	relation.LowerBound = &nodes.IntNode{Value: 0}
	relation.UpperBound = &nodes.IntNode{Value: math.MaxInt}

	// relation_term: (LSB (integer | (integer? DOT DOT integer?))? RSB) relation
	if err := p.eat(lexer.TokenLSB); err != nil {
		return nil, err
	}

	if p.CurrentToken.Type == lexer.TokenRSB {
		// []Relation
		relation.LowerBound = &nodes.IntNode{Value: 1}
		relation.UpperBound = &nodes.IntNode{Value: 1}

		if err := p.eat(lexer.TokenRSB); err != nil {
			return nil, err
		}
	} else {
		// Integer
		if p.CurrentToken.Type == lexer.TokenIntegerConstant {
			number, err := strconv.Atoi(p.CurrentToken.Value)
			if err != nil {
				return nil, err
			}
			relation.LowerBound = &nodes.IntNode{Value: number}
			relation.UpperBound = &nodes.IntNode{Value: number}

			if err := p.eat(lexer.TokenIntegerConstant); err != nil {
				return nil, err
			}

		}
		// DOT DOT
		if p.CurrentToken.Type == lexer.TokenRange {
			relation.UpperBound = &nodes.IntNode{Value: math.MaxInt}

			if err := p.eat(lexer.TokenRange); err != nil {
				return nil, err
			}

			// Integer
			if p.CurrentToken.Type == lexer.TokenIntegerConstant {
				number, err := strconv.Atoi(p.CurrentToken.Value)
				if err != nil {
					return nil, err
				}

				// Lower bound is already set
				relation.UpperBound = &nodes.IntNode{Value: number}

				if err := p.eat(lexer.TokenIntegerConstant); err != nil {
					return nil, err
				}
			}
		}

		if err := p.eat(lexer.TokenRSB); err != nil {
			return nil, err
		}
	}

	if p.CurrentToken.Type == lexer.TokenIdentifier {
		// EdgeName
		relation.EdgeName = &nodes.StringNode{Value: p.CurrentToken.Value}
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}
	} else if p.CurrentToken.Type == lexer.TokenLRB {
		// Unit
		if err := p.eat(lexer.TokenLRB); err != nil {
			return nil, err
		}
		if err := p.eat(lexer.TokenRRB); err != nil {
			return nil, err
		}
		relation.EdgeName = nil
	} else {
    panic(p.CurrentToken)
		return nil, &errors.UnexpectedToken{
			SourceContext:   p.Lexer.GetSourceContext(),
			ActualToken:     p.CurrentToken,
			SuggestedTokens: []lexer.TokenType{lexer.TokenIdentifier, lexer.TokenLRB},
		}
	}

	return relation, nil
}

// vertex_term: vertex (LCB expression RCB)?
func (p *Parser) vertexTerm() (nodes.ASTNode, error) {
	// vertex: ID (as ID)? | unit
	vertex := new(nodes.VertexNode)
	if p.CurrentToken.Type == lexer.TokenIdentifier {
		vertex.VertexName = &nodes.StringNode{Value: p.CurrentToken.Value}
		if err := p.eat(lexer.TokenIdentifier); err != nil {
			return nil, err
		}

		if p.CurrentToken.Type == lexer.TokenAlias {
			if err := p.eat(lexer.TokenAlias); err != nil {
				return nil, err
			}
			vertex.Alias = &nodes.StringNode{Value: p.CurrentToken.Value}
			if err := p.eat(lexer.TokenIdentifier); err != nil {
				return nil, err
			}
		}
	} else if p.CurrentToken.Type == lexer.TokenLRB {
		if err := p.eat(lexer.TokenLRB); err != nil {
			return nil, err
		}
		if err := p.eat(lexer.TokenRRB); err != nil {
			return nil, err
		}
		vertex.VertexName = nil
	} else {
		return nil, &errors.UnexpectedToken{
			SourceContext:   p.Lexer.GetSourceContext(),
			ActualToken:     p.CurrentToken,
			SuggestedTokens: []lexer.TokenType{lexer.TokenIdentifier, lexer.TokenLRB},
		}
	}

	// (LCB expression RCB)?
	if p.CurrentToken.Type == lexer.TokenLCB {
		if err := p.eat(lexer.TokenLCB); err != nil {
			return nil, err
		}
		condition, err := p.expression()
		if err != nil {
			return nil, err
		}
		if err := p.eat(lexer.TokenRCB); err != nil {
			return nil, err
		}
		return &nodes.VertexTermNode{Vertex: vertex, Conditions: condition}, nil
	}
	return &nodes.VertexTermNode{Vertex: vertex, Conditions: nil}, nil
}

func (p *Parser) Parse() (nodes.ASTNode, error) {
	// TODO: if else on the root node
	statements, err := p.programStatement()
	if err != nil {
		return nil, err
	}
	return statements, nil
}
