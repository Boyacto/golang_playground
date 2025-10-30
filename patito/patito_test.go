package main

import (
    "strings"
    "testing"

    "github.com/antlr/antlr4/runtime/Go/antlr"
    gen "patito/gen"
)

type errCatcher struct {
    antlr.DefaultErrorListener
    msgs []string
}

func (e *errCatcher) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
    line, column int, msg string, e2 antlr.RecognitionException) {
    e.msgs = append(e.msgs, msg)
}

func lexTokens(input string) ([]*antlr.CommonToken, *gen.PatitoLexer) {
    is := antlr.NewInputStream(input)
    lex := gen.NewPatitoLexer(is)
    ts := antlr.NewCommonTokenStream(lex, 0)
    ts.Fill()
    var out []*antlr.CommonToken
    for _, t := range ts.GetAllTokens() {
        ct, _ := t.(*antlr.CommonToken)
        if ct.GetTokenType() == antlr.TokenEOF {
            continue
        }
        out = append(out, ct)
    }
    return out, lex
}

func tokName(t *antlr.CommonToken) string {
    tt := t.GetTokenType()
    text := t.GetText()
    switch tt {
    case gen.PatitoLexerPROGRAM:
        return "PROGRAM"
    case gen.PatitoLexerMAIN:
        return "MAIN"
    case gen.PatitoLexerEND:
        return "END"
    case gen.PatitoLexerVARS:
        return "VARS"
    case gen.PatitoLexerVOID:
        return "VOID"
    case gen.PatitoLexerPRINT:
        return "PRINT"
    case gen.PatitoLexerIF:
        return "IF"
    case gen.PatitoLexerELSE:
        return "ELSE"
    case gen.PatitoLexerWHILE:
        return "WHILE"
    case gen.PatitoLexerDO:
        return "DO"
    case gen.PatitoLexerINT_KW:
        return "INT"
    case gen.PatitoLexerFLOAT_KW:
        return "FLOAT"
    case gen.PatitoLexerRELOP:
        return "RELOP(" + text + ")"
    case gen.PatitoLexerASSIGN:
        return "ASSIGN"
    case gen.PatitoLexerPLUS:
        return "PLUS"
    case gen.PatitoLexerMINUS:
        return "MINUS"
    case gen.PatitoLexerMULT:
        return "MULT"
    case gen.PatitoLexerDIV:
        return "DIV"
    case gen.PatitoLexerCOMMA:
        return "COMMA"
    case gen.PatitoLexerSEMI:
        return "SEMICOLON"
    case gen.PatitoLexerCOLON:
        return "COLON"
    case gen.PatitoLexerLPAREN:
        return "LPAREN"
    case gen.PatitoLexerRPAREN:
        return "RPAREN"
    case gen.PatitoLexerLBRACE:
        return "LBRACE"
    case gen.PatitoLexerRBRACE:
        return "RBRACE"
    case gen.PatitoLexerLBRACKET:
        return "LBRACKET"
    case gen.PatitoLexerRBRACKET:
        return "RBRACKET"
    case gen.PatitoLexerID:
        return "ID(" + text + ")"
    case gen.PatitoLexerINT_LIT:
        return "INT_TYPE(" + text + ")"
    case gen.PatitoLexerFLOAT_LIT:
        return "FLOAT_TYPE(" + text + ")"
    case gen.PatitoLexerSTRING_LIT:
        return "STRING_TYPE(" + text + ")"
    default:
        return text
    }
}

func tokenList(input string) []string {
    toks, _ := lexTokens(input)
    out := make([]string, 0, len(toks))
    for _, t := range toks {
        out = append(out, tokName(t))
    }
    return out
}

func parseOK(t *testing.T, input string) []string {
    is := antlr.NewInputStream(input)
    lex := gen.NewPatitoLexer(is)
    ts := antlr.NewCommonTokenStream(lex, 0)
    p := gen.NewPatitoParser(ts)
    errs := &errCatcher{}
    p.RemoveErrorListeners()
    p.AddErrorListener(errs)

    _ = p.Program()

    if len(errs.msgs) > 0 {
        t.Fatalf("unexpected syntax errors: %v", errs.msgs)
    }
    return errs.msgs
}

func parseErr(input string) []string {
    is := antlr.NewInputStream(input)
    lex := gen.NewPatitoLexer(is)
    ts := antlr.NewCommonTokenStream(lex, 0)
    p := gen.NewPatitoParser(ts)
    errs := &errCatcher{}
    p.RemoveErrorListeners()
    p.AddErrorListener(errs)

    _ = p.Program()
    return errs.msgs
}

/* ---------------------------
   7.1 Pruebas del Léxico
   --------------------------- */

func Test_7_1_1_TokensBasicos(t *testing.T) {
    input := `program test ; vars x : int ; main { x = 10 ; } end`
    got := tokenList(input)

    want := []string{
        "PROGRAM", "ID(test)", "SEMICOLON",
        "VARS", "ID(x)", "COLON", "INT", "SEMICOLON",
        "MAIN", "LBRACE",
        "ID(x)", "ASSIGN", "INT_TYPE(10)", "SEMICOLON",
        "RBRACE", "END",
    }
    if strings.Join(got, " ") != strings.Join(want, " ") {
        t.Fatalf("\n got: %v\nwant: %v", got, want)
    }
}

func Test_7_1_2_FloatsYStrings(t *testing.T) {
    input := `pi = 3.14 ; print("Hello World") ;`
    got := tokenList(input)
    want := []string{
        "ID(pi)", "ASSIGN", "FLOAT_TYPE(3.14)", "SEMICOLON",
        "PRINT", "LPAREN", "STRING_TYPE(\"Hello World\")", "RPAREN", "SEMICOLON",
    }
    if strings.Join(got, " ") != strings.Join(want, " ") {
        t.Fatalf("\n got: %v\nwant: %v", got, want)
    }
}

func Test_7_1_3_CaracterInvalido(t *testing.T) {
    input := `x = @ 10 ;`
    is := antlr.NewInputStream(input)
    lex := gen.NewPatitoLexer(is)
    errs := &errCatcher{}
    lex.RemoveErrorListeners()
    lex.AddErrorListener(errs)
    ts := antlr.NewCommonTokenStream(lex, 0)
    ts.Fill()
    if len(errs.msgs) == 0 {
        t.Fatalf("expected lexical error for '@'")
    }
}

func Test_7_1_4_StringSinCerrar(t *testing.T) {
    input := `print("texto sin cerrar`
    is := antlr.NewInputStream(input)
    lex := gen.NewPatitoLexer(is)
    errs := &errCatcher{}
    lex.RemoveErrorListeners()
    lex.AddErrorListener(errs)
    ts := antlr.NewCommonTokenStream(lex, 0)
    ts.Fill()
    if len(errs.msgs) == 0 {
        t.Fatalf("expected lexical error for unterminated string")
    }
}

/* ------------------------------
   7.2 Pruebas del Sintáctico
   ------------------------------ */

func Test_7_2_1_ProgramaMinimoValido(t *testing.T) {
    input := `
program minimal ;
main { }
end
`
    parseOK(t, input)
}

func Test_7_2_2_ProgramaConVarsYExpr(t *testing.T) {
    input := `
program calc ;
vars a, b, result : int ;
main {
  a = 5 ;
  b = 3 ;
  result = a + b * 2 ;
  print("Resultado:", result) ;
}
end
`
    parseOK(t, input)
}

func Test_7_2_3_ControlEstructuras(t *testing.T) {
    input := `
program control ;
vars x : int ;
main {
  x = 10 ;
  if ( x > 5 ) { print(x) ; } else { x = 0 ; }
  while ( x > 0 ) do { x = x - 1 ; } ;
}
end
`
    parseOK(t, input)
}

func Test_7_2_4_FaltaPuntoYComa(t *testing.T) {
    input := `
program test ;
vars x : int
main { }
end
`
    errs := parseErr(input)
    if len(errs) == 0 {
        t.Fatalf("expected syntax error for missing ';' after type declaration")
    }
}

func Test_7_2_5_ParentesisDesbalanceados(t *testing.T) {
    input := `
program test ;
main {
  x = ( 5 + 3 ;
}
end
`
    errs := parseErr(input)
    if len(errs) == 0 {
        t.Fatalf("expected syntax error for missing ')'")
    }
}

func Test_7_2_6_PalabraReservadaComoIdent(t *testing.T) {
    input := `
program test ;
vars while : int ;
main { }
end
`
    errs := parseErr(input)
    if len(errs) == 0 {
        t.Fatalf("expected syntax error for reserved keyword used as identifier")
    }
}
