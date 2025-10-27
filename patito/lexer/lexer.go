package lexer

import "patito/token"

type Lexer struct {
	input        string
	currentIndex int  // current index in input (points to current char)
	nextIndex    int  // current reading index in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextIndex]
	}
	l.currentIndex = l.nextIndex
	l.nextIndex += 1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	currIndex := l.currentIndex
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[currIndex:l.currentIndex]
}

func isLetter(ch byte) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
