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
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456
	FLOAT  = "FLOAT"  // 12.12
	STRING = "STRING" // "Hola"

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
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	PROGRAM    = "program"
	MAIN       = "main"
	END        = "end"
	VAR        = "var"
	VOID       = "void"
	PRINT      = "print"
	IF         = "if"
	ELSE       = "else"
	WHILE      = "while"
	DO         = "do"
	FLOAT_TYPE = "float"
	INT_TYPE   = "int"
)
