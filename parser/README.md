## DDL

```
property_def: ID TYPE
EDGE_TYPE: OneWay | BothWay
schema_def: SCHEMA ID LCB (property_def)* RCB
edge_def: EDGE LCB EDGE_TYPE AS ID RCB
		
property_init: DOT ID EQUAL factor
vertex_init: VERTEX ID AS ID LCB (property_init)* RCB
relation_init: RELATION LCB ID ID ID RCB
```

## DML

```
=====Definitions=====
sum: Sum LRB (arithmatic_expression | array_expression) RRB
max, min: Max LRB array_expression (COMMA property_id)+ RRB
starts_with: StartsWith LRB property[type: string] RRB
expression: boolean_expression | arithmatic_expression | array_expression
=====================

builtin_func: sum | max | min | starts_with
operand: LT | GT | LE | GE | EQ | NE
unit: LRB RRB
relation: DOT (ID | unit)
property: ID? DOT ID
vertex: ID (as ID)? | unit
query: ID

relation_term: 
		relation (LSB (integer | (integer)? DOT DOT (integer)?) RSB)?
vertex_term: vertex (LCB expression RCB)?

factor:
    (INT | STRING | BOOL)
    | property_id
    | vertex_term
    | relation_term vertex_term
    | builtin_func LCB clause RCB 
    | LRB expression RRB

term:
    factor ((MUL | DIV) factor)*

operation:
    term ((PLUS | MINUS) term)* 

clause:
    operation (operator operation)*

expression:
    clause ((AND | OR) clause)*

query_statement:
    query LCB expression RCB


```

## Program

```
program_statement:
	(query_statment|schema_def|vertex_def|add_vertex|add_relation) program_stamement 
	| NoOp

```

## Examples

```
Schema Person {
	name string
	age int
	salary int
}


Vertex Harry Person {
	.name = "Harry" 
	.age = 1
}

Edge LivesIn OneWay

Relation LivesIn {
    Harry 
    London
}

Query Sum(Person as A {
		.name = "Hi"
		or (StartsWith(.name, 'H') and .age > 10)
		and LivesIn () {
			Within[..] Country {
				.name="India"
				and Within Continent
			}
		}
		and Sum(FriendsWith Person {
			LivesIn () {
				Within Country{.name="USA"}
			}
		}, .salary) < .salary
	}, .age)
}

# Query can be recursive
Query {
	Person as A {
		Sum(FriendsWith Person {
			FriendsWith A 
		}, .salary) < .salary
	}
}
```
