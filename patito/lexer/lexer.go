package lexer

import "patito/token"

type Lexer struct {
	input        string
	currentIndex int  // current index in input (points to current char)
	nextIndex    int  // current reading index in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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

func (l *Lexer) peekChar() byte {
	if l.nextIndex >= len(l.input) {
		return 0
	}
	return l.input[l.nextIndex]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string([]byte{ch, l.ch})}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULT, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		// "string literal"
		lit := l.readString()
		tok.Type = token.STRING_TYPE
		tok.Literal = lit
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isAlpha(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT_TYPE
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok
}

func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.currentIndex
	// first must be letter
	// then letter/digit/'_'
	for isAlpha(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.currentIndex]
}

func (l *Lexer) readNumber() (literal string, isFloat bool) {
	start := l.currentIndex
	for isDigit(l.ch) {
		l.readChar()
	}
	// fractional part: '.' followed by at least one digit
	if l.ch == '.' && isDigit(l.peekChar()) {
		isFloat = true
		l.readChar() // consume '.'
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[start:l.currentIndex], isFloat
}
func (l *Lexer) readString() string {
	// current l.ch == '"'
	l.readChar() // consume opening quote
	start := l.currentIndex
	for l.ch != '"' && l.ch != 0 {
		// simple version; no escapes
		l.readChar()
	}
	lit := l.input[start:l.currentIndex]
	if l.ch == '"' {
		l.readChar() // consume closing quote
	}
	return lit
}

func isLetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		ch == '_'
}

func isAlpha(ch byte) bool { // letters only, no underscore
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
