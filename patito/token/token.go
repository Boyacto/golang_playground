package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT       TokenType = "IDENT"  // add, foobar, x, y, ...
	INT_TYPE    TokenType = "INT"    // 1343456
	FLOAT_TYPE  TokenType = "FLOAT"  // 12.12
	STRING_TYPE TokenType = "STRING" // "Hola"

	// Assignment operator
	ASSIGN TokenType = "="

	// Arithmetic Operators
	PLUS  TokenType = "+"
	MINUS TokenType = "-"
	MULT  TokenType = "*"
	DIV   TokenType = "/"

	// Comparison (Relational) Operators
	EQ  TokenType = "==" // Equal
	NEQ TokenType = "!=" // Not Equal
	LT  TokenType = "<"  // Less Than
	GT  TokenType = ">"  // Greater Than
	LEQ TokenType = "<=" // Less or Equal
	GEQ TokenType = ">=" // Greater or Equal

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	LBRACKET  TokenType = "["
	RBRACKET  TokenType = "]"

	// Keywords
	PROGRAM TokenType = "program"
	MAIN    TokenType = "main"
	END     TokenType = "end"
	VAR     TokenType = "var" // we'll map both "var" and "vars" → VAR
	VOID    TokenType = "void"
	PRINT   TokenType = "print"
	IF      TokenType = "if"
	ELSE    TokenType = "else"
	WHILE   TokenType = "while"
	DO      TokenType = "do"
	FLOAT   TokenType = "float" // type keyword
	INT     TokenType = "int"   // type keyword
)

var keywords = map[string]TokenType{
	"program": PROGRAM,
	"main":    MAIN,
	"end":     END,

	// variable section keyword(s)
	"var":  VAR,
	"vars": VAR,

	"void":  VOID,
	"print": PRINT,
	"if":    IF,
	"else":  ELSE,
	"while": WHILE,
	"do":    DO,

	// type keywords
	"float": FLOAT,
	"int":   INT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
