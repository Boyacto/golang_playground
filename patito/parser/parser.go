package parser

import (
	"patito/ast"
	"patito/lexer"
	"patito/token"
	"strconv"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Decls = append(prog.Decls, stmt)
		}
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	if p.currToken.Type == token.IDENT && p.peekToken.Type == token.ASSIGN {
		return p.parseAssignStatement()
	}
	return nil
}

func (p *Parser) parseAssignStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Name: &ast.Identifier{Value: p.currToken.Literal}}
	p.nextToken()
	p.nextToken()
	stmt.Value = p.parseExpression()
	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	left := p.parsePrimary()
	for isOp(p.peekToken.Type) {
		op := p.peekToken.Literal
		p.nextToken()
		p.nextToken()
		right := p.parsePrimary()
		left = &ast.InfixExpression{Left: left, Operator: op, Right: right}
	}
	return left
}

func (p *Parser) parsePrimary() ast.Expression {
	switch p.currToken.Type {
	case token.IDENT:
		return &ast.Identifier{Value: p.currToken.Literal}
	case token.INT_TYPE:
		v, _ := strconv.ParseInt(p.currToken.Literal, 10, 64)
		return &ast.IntegerLiteral{Value: v}
	case token.FLOAT_TYPE:
		f, _ := strconv.ParseFloat(p.currToken.Literal, 64)
		return &ast.FloatLiteral{Value: f}
	case token.LPAREN:
		p.nextToken()
		exp := p.parseExpression()
		for p.currToken.Type != token.RPAREN && p.currToken.Type != token.EOF {
			p.nextToken()
		}
		return exp
	default:
		return &ast.Identifier{Value: p.currToken.Literal}
	}
}

func isOp(t string) bool {
	switch t {
	case token.PLUS, token.MINUS, token.MULT, token.DIV:
		return true
	}
	return false
}
