package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT       = "IDENT"  // add, foobar, x, y, ...
	INT_TYPE    = "INT"    // 1343456
	FLOAT_TYPE  = "FLOAT"  // 12.12
	STRING_TYPE = "STRING" // "Hola"

	// Assignment operator
	ASSIGN = "="

	// Arithmetic Operators
	PLUS  = "+"
	MINUS = "-"
	MULT  = "*"
	DIV   = "/"

	// Comparison (Relational) Operators
	EQ  = "==" //Equal
	NEQ = "!=" //Non Equal
	LT  = "<"  //Less Than
	GT  = ">"  //Greater Than
	LEQ = "<=" // Less or Equal
	GEQ = ">=" // Greater or Equal

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	PROGRAM = "program"
	MAIN    = "main"
	END     = "end"
	VAR     = "var"
	VOID    = "void"
	PRINT   = "print"
	IF      = "if"
	ELSE    = "else"
	WHILE   = "while"
	DO      = "do"
	FLOAT   = "float"
	INT     = "int"
)

var keywords = map[string]TokenType{
	"program": PROGRAM,
	"main":    MAIN,
	"end":     END,
	"var":     VAR,
	"void":    VOID,
	"print":   PRINT,
	"if":      IF,
	"else":    ELSE,
	"while":   WHILE,
	"do":      DO,
	"float":   FLOAT,
	"int":     INT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
