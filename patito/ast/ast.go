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

type Statement interface {
	Node
	statementNode() // Marker method to distinguish statements from expressions
}

type Expression interface {
	Node
	expressionNode() // Marker method to distinguish expressions from statements
}

// Assignment: ID = Expression ;
type AssignStatement struct {
	Name  *Identifier
	Value Expression
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Name.Value }

// Print: print(ExpressionList)
type PrintStatement struct {
	Expressions []Expression
}

func (ps *PrintStatement) statementNode()       {}
func (ps *PrintStatement) TokenLiteral() string { return "print" }

// If / Else
type IfStatement struct {
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return "if" }

// While
type WhileStatement struct {
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return "while" }

// ---------- Expressions ----------

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Value }

type IntegerLiteral struct {
	Value int64 // Idk why but int in go varies between 32 and 64 bit (maybe because it is for system programming.)
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return "int" }

type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return "float" }

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Operator }

type CallExpression struct {
	Function  *Identifier
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Function.Value }

// ---------- Blocks ----------

type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return "{" }
