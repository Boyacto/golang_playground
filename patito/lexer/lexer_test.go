package lexer

import (
	"testing"

	"patito/token"
)

func TestNextToken(t *testing.T) {
	input := `program test ; var x : int ; main { x = 10 ; } end`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PROGRAM, "program"},
		{token.IDENT, "test"},
		{token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "x"},
		{token.COLON, ":"},
		{token.INT, "int"},
		{token.SEMICOLON, ";"},
		{token.MAIN, "main"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT_TYPE, "10"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "{"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
