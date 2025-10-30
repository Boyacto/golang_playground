//This code is based on the book writing_an_interpreter_in_go_1.7 and adapted to patito.

// Package ast defines the Abstract Syntax Tree (AST) node types for the Patito language.
// The AST represents the hierarchical structure of a program after parsing.
//
// Design Philosophy:
// - All AST nodes implement the Node interface
// - Statements and Expressions are distinct categories of nodes
package ast

// Node is the base interface that all AST nodes must implement.
// Every node in the syntax tree can return its token literal, which is useful
// for debugging, error messages, and representing the original source code.
type Node interface {
	TokenLiteral() string // Returns the literal value of the token this node is associated with
}

// Statement represents a statement in the program.
// In Patito, statements are instructions that perform actions but don't produce values.
// Examples: variable declarations, assignments, print statements, control flow (if/while).
//
// The statementNode() method is a marker method (not meant to be called) that ensures
// only statement types can be assigned to Statement variables, providing compile-time type safety.
type Statement interface {
	Node
	statementNode() // Marker method to distinguish statements from expressions
}

// Expression represents an expression in the program.
// In Patito, expressions are constructs that evaluate to values.
// Examples: arithmetic operations (1 + 2), identifiers (x), literals (42, "hello"), comparisons (x < 10).
//
// The expressionNode() method is a marker method (not meant to be called) that ensures
// only expression types can be assigned to Expression variables, providing compile-time type safety.
type Expression interface {
	Node
	expressionNode() // Marker method to distinguish expressions from statements
}

// Program is the root node of every AST produced by the parser.
// It represents the entire source file and contains all top-level statements.
// Every valid Patito program is represented as a Program containing a slice of Statements.
type Program struct {
	Statements []Statement // All statements in the program, in order of appearance
}

// TokenLiteral returns the token literal of the first statement in the program.
// This is primarily used for debugging and error reporting.
// If the program is empty, it returns an empty string.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
