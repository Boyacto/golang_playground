// Package lexer provides lexical analysis (tokenization) for the Patito language.
// It converts raw source code text into a stream of tokens that can be used by the parser.
package lexer

import "patito/token"

// Lexer performs lexical analysis by reading input character by character
// and producing tokens. It uses a two-pointer approach for lookahead capability.
type Lexer struct {
	input        string // the source code being tokenized
	currentIndex int    // position of current character being examined
	nextIndex    int    // position of next character to read (enables 1-char lookahead)
	ch           byte   // current character under examination (0 if at EOF)
}

// New creates and initializes a new Lexer for the given input string.
// It positions the lexer at the first character by calling readChar().
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialize by reading the first character
	return l
}

// readChar advances the lexer by one character in the input.
// It moves both currentIndex and nextIndex forward, and sets ch to the next character.
// If we've reached the end of input, ch is set to 0 (NUL) to signal EOF.
func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.ch = 0 // ASCII NUL character represents EOF
	} else {
		l.ch = l.input[l.nextIndex]
	}
	l.currentIndex = l.nextIndex
	l.nextIndex += 1
}

// peekChar returns the next character without advancing the lexer position.
// This enables lookahead for multi-character tokens like "==", "!=", "<=", ">=".
// Returns 0 (NUL) if at end of input.
func (l *Lexer) peekChar() byte {
	if l.nextIndex >= len(l.input) {
		return 0
	}
	return l.input[l.nextIndex]
}

// NextToken reads and returns the next token from the input.
// This is the main entry point for tokenization. It:
// 1. Skips whitespace
// 2. Identifies the current character and determines what token it starts
// 3. Handles multi-character tokens via lookahead (==, !=, <=, >=)
// 4. Returns the token and advances the lexer position
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip any whitespace characters (spaces, tabs, newlines)
	l.consumeWhitespace()

	// Determine token type based on current character
	switch l.ch {
	// Assignment or equality operator: '=' or '=='
	case '=':
		if l.peekChar() == '=' {
			// Two-character token: '=='
			ch := l.ch
			l.readChar() // consume second '='
			tok = token.Token{Type: token.EQ, Literal: string([]byte{ch, l.ch})}
		} else {
			// Single character: '='
			tok = newToken(token.ASSIGN, l.ch)
		}

	// Not-equal operator: '!=' (standalone '!' is not supported in this language)
	case '!':
		if l.peekChar() == '=' {
			// Two-character token: '!='
			ch := l.ch
			l.readChar() // consume '='
			tok = token.Token{Type: token.NEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			// Standalone '!' is illegal - this language doesn't have a NOT operator
			tok = newToken(token.ILLEGAL, l.ch)
		}

	// Less-than or less-than-or-equal: '<' or '<='
	case '<':
		if l.peekChar() == '=' {
			// Two-character token: '<='
			ch := l.ch
			l.readChar() // consume '='
			tok = token.Token{Type: token.LEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			// Single character: '<'
			tok = newToken(token.LT, l.ch)
		}

	// Greater-than or greater-than-or-equal: '>' or '>='
	case '>':
		if l.peekChar() == '=' {
			// Two-character token: '>='
			ch := l.ch
			l.readChar() // consume '='
			tok = token.Token{Type: token.GEQ, Literal: string([]byte{ch, l.ch})}
		} else {
			// Single character: '>'
			tok = newToken(token.GT, l.ch)
		}

	// Arithmetic operators
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULT, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)

	// Delimiters and punctuation
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)

	// Parentheses: ( )
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)

	// Braces: { }
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	// Brackets: [ ]
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)

	// String literals: "..."
	case '"':
		lit := l.readString()
		tok.Type = token.STRING_TYPE
		tok.Literal = lit
		return tok // return early; readString() already consumed the closing quote

	// End of input
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	// Identifiers (keywords or variable names) and numeric literals
	default:
		if isAlpha(l.ch) {
			// Starts with a letter: read full identifier (keyword or variable name)
			tok.Literal = l.readIdentifier()
			// Check if it's a reserved keyword; if not, it's an IDENT
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // return early; readIdentifier() already consumed the identifier
		} else if isDigit(l.ch) {
			// Starts with a digit: read numeric literal (integer or float)
			literal, isFloat := l.readNumber()
			tok.Literal = literal
			if isFloat {
				tok.Type = token.FLOAT_TYPE
			} else {
				tok.Type = token.INT_TYPE
			}
			return tok // return early; readNumber() already consumed the number
		} else {
			// Unknown character - mark as illegal
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	// Advance to next character for single-character tokens
	// (multi-character tokens return early above)
	l.readChar()
	return tok
}

// consumeWhitespace skips over all whitespace characters (space, tab, newline, carriage return).
// It advances the lexer until a non-whitespace character is found.
func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier consumes and returns a complete identifier from the input.
// Identifiers must start with a letter (A-Z or a-z), and can contain letters,
// digits, and underscores after the first character.
// Examples: "foo", "myVar", "x123", "hello_world"
func (l *Lexer) readIdentifier() string {
	start := l.currentIndex
	// Continue reading while we see letters, digits, or underscores
	for isAlpha(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[start:l.currentIndex]
}

// readNumber consumes and returns a numeric literal (integer or float).
// It returns both the literal string and a boolean indicating if it's a float.
// Integers: "123", "0", "999"
// Floats: "12.34", "0.5" (must have digits after decimal point)
func (l *Lexer) readNumber() (literal string, isFloat bool) {
	start := l.currentIndex
	// Read the integer part
	for isDigit(l.ch) {
		l.readChar()
	}
	// Check for fractional part: '.' followed by at least one digit
	// This ensures "123." is not treated as a float
	if l.ch == '.' && isDigit(l.peekChar()) {
		isFloat = true
		l.readChar() // consume '.'
		// Read the fractional part
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[start:l.currentIndex], isFloat
}

// readString consumes and returns a string literal.
// Assumes the current character is '"' (opening quote).
// Returns the string content without the surrounding quotes.
// Note: This is a simple implementation that doesn't handle escape sequences.
func (l *Lexer) readString() string {
	l.readChar() // consume opening quote '"'
	start := l.currentIndex
	// Read until we find closing quote or EOF
	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}
	lit := l.input[start:l.currentIndex]
	if l.ch == '"' {
		l.readChar() // consume closing quote '"'
	}
	// TODO: Handle unclosed strings (when l.ch == 0) as an error
	return lit
}

// isLetter checks if a character is a letter (A-Z, a-z) or underscore.
// This function includes underscore as a valid "letter" character.
func isLetter(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') ||
		(ch >= 'a' && ch <= 'z') ||
		ch == '_'
}

// isAlpha checks if a character is strictly an alphabetic letter (A-Z, a-z).
// Unlike isLetter(), this does NOT include underscore.
// Used for validating the first character of identifiers.
func isAlpha(ch byte) bool {
	return ('A' <= ch && ch <= 'Z') || ('a' <= ch && ch <= 'z')
}

// isDigit checks if a character is a numeric digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken is a helper function to create a token from a token type and a single character.
// It converts the byte character to a string for the Literal field.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
