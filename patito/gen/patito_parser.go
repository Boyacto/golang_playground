// Code generated from Patito.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gen // Patito
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 225,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 4, 7, 4, 66, 10, 4, 12,
	4, 14, 4, 69, 11, 4, 3, 5, 3, 5, 6, 5, 73, 10, 5, 13, 5, 14, 5, 74, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 85, 10, 7, 12, 7, 14,
	7, 88, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 94, 10, 9, 12, 9, 14, 9, 97,
	11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 5, 10, 109, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7, 13, 125, 10, 13, 12,
	13, 14, 13, 128, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 137, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 154, 10, 16,
	3, 16, 3, 16, 5, 16, 158, 10, 16, 3, 16, 5, 16, 161, 10, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 172, 10, 16,
	12, 16, 14, 16, 175, 11, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 5, 20, 194, 10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 199, 10, 21, 12, 21,
	14, 21, 202, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 24, 3, 24, 5, 24, 215, 10, 24, 3, 25, 3, 25, 3, 25, 7, 25,
	220, 10, 25, 12, 25, 14, 25, 223, 11, 25, 3, 25, 2, 3, 30, 26, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 2, 6, 3, 2, 13, 14, 3, 2, 17, 18, 3, 2, 19, 20, 3, 2, 31, 33,
	2, 224, 2, 50, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 6, 67, 3, 2, 2, 2, 8, 70,
	3, 2, 2, 2, 10, 76, 3, 2, 2, 2, 12, 81, 3, 2, 2, 2, 14, 89, 3, 2, 2, 2,
	16, 91, 3, 2, 2, 2, 18, 108, 3, 2, 2, 2, 20, 110, 3, 2, 2, 2, 22, 115,
	3, 2, 2, 2, 24, 121, 3, 2, 2, 2, 26, 129, 3, 2, 2, 2, 28, 138, 3, 2, 2,
	2, 30, 160, 3, 2, 2, 2, 32, 176, 3, 2, 2, 2, 34, 178, 3, 2, 2, 2, 36, 180,
	3, 2, 2, 2, 38, 193, 3, 2, 2, 2, 40, 195, 3, 2, 2, 2, 42, 203, 3, 2, 2,
	2, 44, 207, 3, 2, 2, 2, 46, 214, 3, 2, 2, 2, 48, 216, 3, 2, 2, 2, 50, 51,
	7, 3, 2, 2, 51, 52, 7, 30, 2, 2, 52, 53, 7, 22, 2, 2, 53, 54, 5, 4, 3,
	2, 54, 55, 5, 6, 4, 2, 55, 56, 7, 4, 2, 2, 56, 57, 5, 16, 9, 2, 57, 58,
	7, 5, 2, 2, 58, 59, 7, 2, 2, 3, 59, 3, 3, 2, 2, 2, 60, 63, 5, 8, 5, 2,
	61, 63, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 5, 3, 2,
	2, 2, 64, 66, 5, 36, 19, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67,
	65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 7, 3, 2, 2, 2, 69, 67, 3, 2, 2,
	2, 70, 72, 7, 6, 2, 2, 71, 73, 5, 10, 6, 2, 72, 71, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 9, 3, 2, 2, 2,
	76, 77, 5, 12, 7, 2, 77, 78, 7, 23, 2, 2, 78, 79, 5, 14, 8, 2, 79, 80,
	7, 22, 2, 2, 80, 11, 3, 2, 2, 2, 81, 86, 7, 30, 2, 2, 82, 83, 7, 21, 2,
	2, 83, 85, 7, 30, 2, 2, 84, 82, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84,
	3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 13, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2,
	89, 90, 9, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 95, 7, 26, 2, 2, 92, 94, 5,
	18, 10, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2,
	95, 96, 3, 2, 2, 2, 96, 98, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 99, 7,
	27, 2, 2, 99, 17, 3, 2, 2, 2, 100, 109, 5, 20, 11, 2, 101, 109, 5, 26,
	14, 2, 102, 109, 5, 28, 15, 2, 103, 104, 5, 44, 23, 2, 104, 105, 7, 22,
	2, 2, 105, 109, 3, 2, 2, 2, 106, 109, 5, 22, 12, 2, 107, 109, 5, 16, 9,
	2, 108, 100, 3, 2, 2, 2, 108, 101, 3, 2, 2, 2, 108, 102, 3, 2, 2, 2, 108,
	103, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 19, 3,
	2, 2, 2, 110, 111, 7, 30, 2, 2, 111, 112, 7, 16, 2, 2, 112, 113, 5, 30,
	16, 2, 113, 114, 7, 22, 2, 2, 114, 21, 3, 2, 2, 2, 115, 116, 7, 8, 2, 2,
	116, 117, 7, 24, 2, 2, 117, 118, 5, 24, 13, 2, 118, 119, 7, 25, 2, 2, 119,
	120, 7, 22, 2, 2, 120, 23, 3, 2, 2, 2, 121, 126, 5, 30, 16, 2, 122, 123,
	7, 21, 2, 2, 123, 125, 5, 30, 16, 2, 124, 122, 3, 2, 2, 2, 125, 128, 3,
	2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 25, 3, 2, 2,
	2, 128, 126, 3, 2, 2, 2, 129, 130, 7, 9, 2, 2, 130, 131, 7, 24, 2, 2, 131,
	132, 5, 30, 16, 2, 132, 133, 7, 25, 2, 2, 133, 136, 5, 16, 9, 2, 134, 135,
	7, 10, 2, 2, 135, 137, 5, 16, 9, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3,
	2, 2, 2, 137, 27, 3, 2, 2, 2, 138, 139, 7, 11, 2, 2, 139, 140, 7, 24, 2,
	2, 140, 141, 5, 30, 16, 2, 141, 142, 7, 25, 2, 2, 142, 143, 7, 12, 2, 2,
	143, 144, 5, 16, 9, 2, 144, 145, 7, 22, 2, 2, 145, 29, 3, 2, 2, 2, 146,
	147, 8, 16, 1, 2, 147, 148, 7, 24, 2, 2, 148, 149, 5, 30, 16, 2, 149, 150,
	7, 25, 2, 2, 150, 161, 3, 2, 2, 2, 151, 161, 5, 44, 23, 2, 152, 154, 5,
	32, 17, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 3, 2,
	2, 2, 155, 161, 7, 30, 2, 2, 156, 158, 5, 32, 17, 2, 157, 156, 3, 2, 2,
	2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 5, 34, 18, 2,
	160, 146, 3, 2, 2, 2, 160, 151, 3, 2, 2, 2, 160, 153, 3, 2, 2, 2, 160,
	157, 3, 2, 2, 2, 161, 173, 3, 2, 2, 2, 162, 163, 12, 9, 2, 2, 163, 164,
	7, 15, 2, 2, 164, 172, 5, 30, 16, 10, 165, 166, 12, 8, 2, 2, 166, 167,
	9, 3, 2, 2, 167, 172, 5, 30, 16, 9, 168, 169, 12, 7, 2, 2, 169, 170, 9,
	4, 2, 2, 170, 172, 5, 30, 16, 8, 171, 162, 3, 2, 2, 2, 171, 165, 3, 2,
	2, 2, 171, 168, 3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2,
	173, 174, 3, 2, 2, 2, 174, 31, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 177,
	9, 3, 2, 2, 177, 33, 3, 2, 2, 2, 178, 179, 9, 5, 2, 2, 179, 35, 3, 2, 2,
	2, 180, 181, 7, 7, 2, 2, 181, 182, 7, 30, 2, 2, 182, 183, 7, 24, 2, 2,
	183, 184, 5, 38, 20, 2, 184, 185, 7, 25, 2, 2, 185, 186, 7, 26, 2, 2, 186,
	187, 5, 4, 3, 2, 187, 188, 5, 16, 9, 2, 188, 189, 7, 27, 2, 2, 189, 190,
	7, 22, 2, 2, 190, 37, 3, 2, 2, 2, 191, 194, 5, 40, 21, 2, 192, 194, 3,
	2, 2, 2, 193, 191, 3, 2, 2, 2, 193, 192, 3, 2, 2, 2, 194, 39, 3, 2, 2,
	2, 195, 200, 5, 42, 22, 2, 196, 197, 7, 21, 2, 2, 197, 199, 5, 42, 22,
	2, 198, 196, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200,
	201, 3, 2, 2, 2, 201, 41, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 7,
	30, 2, 2, 204, 205, 7, 23, 2, 2, 205, 206, 5, 14, 8, 2, 206, 43, 3, 2,
	2, 2, 207, 208, 7, 30, 2, 2, 208, 209, 7, 24, 2, 2, 209, 210, 5, 46, 24,
	2, 210, 211, 7, 25, 2, 2, 211, 45, 3, 2, 2, 2, 212, 215, 5, 48, 25, 2,
	213, 215, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215,
	47, 3, 2, 2, 2, 216, 221, 5, 30, 16, 2, 217, 218, 7, 21, 2, 2, 218, 220,
	5, 30, 16, 2, 219, 217, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3,
	2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 49, 3, 2, 2, 2, 223, 221, 3, 2, 2,
	2, 19, 62, 67, 74, 86, 95, 108, 126, 136, 153, 157, 160, 171, 173, 193,
	200, 214, 221,
}
var literalNames = []string{
	"", "'program'", "'main'", "'end'", "'vars'", "'void'", "'print'", "'if'",
	"'else'", "'while'", "'do'", "'int'", "'float'", "", "'='", "'+'", "'-'",
	"'*'", "'/'", "','", "';'", "':'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "PROGRAM", "MAIN", "END", "VARS", "VOID", "PRINT", "IF", "ELSE", "WHILE",
	"DO", "INT_KW", "FLOAT_KW", "RELOP", "ASSIGN", "PLUS", "MINUS", "MULT",
	"DIV", "COMMA", "SEMI", "COLON", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
	"LBRACKET", "RBRACKET", "ID", "INT_LIT", "FLOAT_LIT", "STRING_LIT", "WS",
}

var ruleNames = []string{
	"program", "varsOpt", "funcsOpt", "varsDecl", "varGroup", "idList", "typeSpec",
	"body", "statement", "assign", "printStmt", "printArgs", "condition", "cycle",
	"expr", "sign", "literal", "funcDecl", "paramListOpt", "paramList", "param",
	"fcall", "argListOpt", "argList",
}

type PatitoParser struct {
	*antlr.BaseParser
}

// NewPatitoParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PatitoParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPatitoParser(input antlr.TokenStream) *PatitoParser {
	this := new(PatitoParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Patito.g4"

	return this
}

// PatitoParser tokens.
const (
	PatitoParserEOF        = antlr.TokenEOF
	PatitoParserPROGRAM    = 1
	PatitoParserMAIN       = 2
	PatitoParserEND        = 3
	PatitoParserVARS       = 4
	PatitoParserVOID       = 5
	PatitoParserPRINT      = 6
	PatitoParserIF         = 7
	PatitoParserELSE       = 8
	PatitoParserWHILE      = 9
	PatitoParserDO         = 10
	PatitoParserINT_KW     = 11
	PatitoParserFLOAT_KW   = 12
	PatitoParserRELOP      = 13
	PatitoParserASSIGN     = 14
	PatitoParserPLUS       = 15
	PatitoParserMINUS      = 16
	PatitoParserMULT       = 17
	PatitoParserDIV        = 18
	PatitoParserCOMMA      = 19
	PatitoParserSEMI       = 20
	PatitoParserCOLON      = 21
	PatitoParserLPAREN     = 22
	PatitoParserRPAREN     = 23
	PatitoParserLBRACE     = 24
	PatitoParserRBRACE     = 25
	PatitoParserLBRACKET   = 26
	PatitoParserRBRACKET   = 27
	PatitoParserID         = 28
	PatitoParserINT_LIT    = 29
	PatitoParserFLOAT_LIT  = 30
	PatitoParserSTRING_LIT = 31
	PatitoParserWS         = 32
)

// PatitoParser rules.
const (
	PatitoParserRULE_program      = 0
	PatitoParserRULE_varsOpt      = 1
	PatitoParserRULE_funcsOpt     = 2
	PatitoParserRULE_varsDecl     = 3
	PatitoParserRULE_varGroup     = 4
	PatitoParserRULE_idList       = 5
	PatitoParserRULE_typeSpec     = 6
	PatitoParserRULE_body         = 7
	PatitoParserRULE_statement    = 8
	PatitoParserRULE_assign       = 9
	PatitoParserRULE_printStmt    = 10
	PatitoParserRULE_printArgs    = 11
	PatitoParserRULE_condition    = 12
	PatitoParserRULE_cycle        = 13
	PatitoParserRULE_expr         = 14
	PatitoParserRULE_sign         = 15
	PatitoParserRULE_literal      = 16
	PatitoParserRULE_funcDecl     = 17
	PatitoParserRULE_paramListOpt = 18
	PatitoParserRULE_paramList    = 19
	PatitoParserRULE_param        = 20
	PatitoParserRULE_fcall        = 21
	PatitoParserRULE_argListOpt   = 22
	PatitoParserRULE_argList      = 23
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(PatitoParserPROGRAM, 0)
}

func (s *ProgramContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *ProgramContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *ProgramContext) VarsOpt() IVarsOptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsOptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsOptContext)
}

func (s *ProgramContext) FuncsOpt() IFuncsOptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncsOptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncsOptContext)
}

func (s *ProgramContext) MAIN() antlr.TerminalNode {
	return s.GetToken(PatitoParserMAIN, 0)
}

func (s *ProgramContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ProgramContext) END() antlr.TerminalNode {
	return s.GetToken(PatitoParserEND, 0)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PatitoParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PatitoParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(PatitoParserPROGRAM)
	}
	{
		p.SetState(49)
		p.Match(PatitoParserID)
	}
	{
		p.SetState(50)
		p.Match(PatitoParserSEMI)
	}
	{
		p.SetState(51)
		p.VarsOpt()
	}
	{
		p.SetState(52)
		p.FuncsOpt()
	}
	{
		p.SetState(53)
		p.Match(PatitoParserMAIN)
	}
	{
		p.SetState(54)
		p.Body()
	}
	{
		p.SetState(55)
		p.Match(PatitoParserEND)
	}
	{
		p.SetState(56)
		p.Match(PatitoParserEOF)
	}

	return localctx
}

// IVarsOptContext is an interface to support dynamic dispatch.
type IVarsOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsOptContext differentiates from other interfaces.
	IsVarsOptContext()
}

type VarsOptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsOptContext() *VarsOptContext {
	var p = new(VarsOptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_varsOpt
	return p
}

func (*VarsOptContext) IsVarsOptContext() {}

func NewVarsOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsOptContext {
	var p = new(VarsOptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_varsOpt

	return p
}

func (s *VarsOptContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsOptContext) VarsDecl() IVarsDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsDeclContext)
}

func (s *VarsOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterVarsOpt(s)
	}
}

func (s *VarsOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitVarsOpt(s)
	}
}

func (s *VarsOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitVarsOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) VarsOpt() (localctx IVarsOptContext) {
	this := p
	_ = this

	localctx = NewVarsOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PatitoParserRULE_varsOpt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PatitoParserVARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.VarsDecl()
		}

	case PatitoParserMAIN, PatitoParserVOID, PatitoParserLBRACE:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncsOptContext is an interface to support dynamic dispatch.
type IFuncsOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncsOptContext differentiates from other interfaces.
	IsFuncsOptContext()
}

type FuncsOptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsOptContext() *FuncsOptContext {
	var p = new(FuncsOptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_funcsOpt
	return p
}

func (*FuncsOptContext) IsFuncsOptContext() {}

func NewFuncsOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsOptContext {
	var p = new(FuncsOptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_funcsOpt

	return p
}

func (s *FuncsOptContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsOptContext) AllFuncDecl() []IFuncDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem())
	var tst = make([]IFuncDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncDeclContext)
		}
	}

	return tst
}

func (s *FuncsOptContext) FuncDecl(i int) IFuncDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *FuncsOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterFuncsOpt(s)
	}
}

func (s *FuncsOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitFuncsOpt(s)
	}
}

func (s *FuncsOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitFuncsOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) FuncsOpt() (localctx IFuncsOptContext) {
	this := p
	_ = this

	localctx = NewFuncsOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PatitoParserRULE_funcsOpt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserVOID {
		{
			p.SetState(62)
			p.FuncDecl()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarsDeclContext is an interface to support dynamic dispatch.
type IVarsDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarsDeclContext differentiates from other interfaces.
	IsVarsDeclContext()
}

type VarsDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsDeclContext() *VarsDeclContext {
	var p = new(VarsDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_varsDecl
	return p
}

func (*VarsDeclContext) IsVarsDeclContext() {}

func NewVarsDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsDeclContext {
	var p = new(VarsDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_varsDecl

	return p
}

func (s *VarsDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsDeclContext) VARS() antlr.TerminalNode {
	return s.GetToken(PatitoParserVARS, 0)
}

func (s *VarsDeclContext) AllVarGroup() []IVarGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarGroupContext)(nil)).Elem())
	var tst = make([]IVarGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarGroupContext)
		}
	}

	return tst
}

func (s *VarsDeclContext) VarGroup(i int) IVarGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarGroupContext)
}

func (s *VarsDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterVarsDecl(s)
	}
}

func (s *VarsDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitVarsDecl(s)
	}
}

func (s *VarsDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitVarsDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) VarsDecl() (localctx IVarsDeclContext) {
	this := p
	_ = this

	localctx = NewVarsDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PatitoParserRULE_varsDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(PatitoParserVARS)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PatitoParserID {
		{
			p.SetState(69)
			p.VarGroup()
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVarGroupContext is an interface to support dynamic dispatch.
type IVarGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarGroupContext differentiates from other interfaces.
	IsVarGroupContext()
}

type VarGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarGroupContext() *VarGroupContext {
	var p = new(VarGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_varGroup
	return p
}

func (*VarGroupContext) IsVarGroupContext() {}

func NewVarGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarGroupContext {
	var p = new(VarGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_varGroup

	return p
}

func (s *VarGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *VarGroupContext) IdList() IIdListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *VarGroupContext) COLON() antlr.TerminalNode {
	return s.GetToken(PatitoParserCOLON, 0)
}

func (s *VarGroupContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *VarGroupContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *VarGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterVarGroup(s)
	}
}

func (s *VarGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitVarGroup(s)
	}
}

func (s *VarGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitVarGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) VarGroup() (localctx IVarGroupContext) {
	this := p
	_ = this

	localctx = NewVarGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PatitoParserRULE_varGroup)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.IdList()
	}
	{
		p.SetState(75)
		p.Match(PatitoParserCOLON)
	}
	{
		p.SetState(76)
		p.TypeSpec()
	}
	{
		p.SetState(77)
		p.Match(PatitoParserSEMI)
	}

	return localctx
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_idList
	return p
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserID)
}

func (s *IdListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserID, i)
}

func (s *IdListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserCOMMA)
}

func (s *IdListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserCOMMA, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterIdList(s)
	}
}

func (s *IdListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitIdList(s)
	}
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) IdList() (localctx IIdListContext) {
	this := p
	_ = this

	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PatitoParserRULE_idList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(PatitoParserID)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserCOMMA {
		{
			p.SetState(80)
			p.Match(PatitoParserCOMMA)
		}
		{
			p.SetState(81)
			p.Match(PatitoParserID)
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) INT_KW() antlr.TerminalNode {
	return s.GetToken(PatitoParserINT_KW, 0)
}

func (s *TypeSpecContext) FLOAT_KW() antlr.TerminalNode {
	return s.GetToken(PatitoParserFLOAT_KW, 0)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) TypeSpec() (localctx ITypeSpecContext) {
	this := p
	_ = this

	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PatitoParserRULE_typeSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PatitoParserINT_KW || _la == PatitoParserFLOAT_KW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PatitoParserLBRACE, 0)
}

func (s *BodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PatitoParserRBRACE, 0)
}

func (s *BodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Body() (localctx IBodyContext) {
	this := p
	_ = this

	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PatitoParserRULE_body)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(PatitoParserLBRACE)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PatitoParserPRINT)|(1<<PatitoParserIF)|(1<<PatitoParserWHILE)|(1<<PatitoParserLBRACE)|(1<<PatitoParserID))) != 0 {
		{
			p.SetState(90)
			p.Statement()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Match(PatitoParserRBRACE)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatementContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *StatementContext) Cycle() ICycleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICycleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICycleContext)
}

func (s *StatementContext) Fcall() IFcallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFcallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFcallContext)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *StatementContext) PrintStmt() IPrintStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PatitoParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Condition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Cycle()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Fcall()
		}
		{
			p.SetState(102)
			p.Match(PatitoParserSEMI)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.PrintStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.Body()
		}

	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PatitoParserASSIGN, 0)
}

func (s *AssignContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Assign() (localctx IAssignContext) {
	this := p
	_ = this

	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PatitoParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PatitoParserID)
	}
	{
		p.SetState(109)
		p.Match(PatitoParserASSIGN)
	}
	{
		p.SetState(110)
		p.expr(0)
	}
	{
		p.SetState(111)
		p.Match(PatitoParserSEMI)
	}

	return localctx
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_printStmt
	return p
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(PatitoParserPRINT, 0)
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *PrintStmtContext) PrintArgs() IPrintArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintArgsContext)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *PrintStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) PrintStmt() (localctx IPrintStmtContext) {
	this := p
	_ = this

	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PatitoParserRULE_printStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(PatitoParserPRINT)
	}
	{
		p.SetState(114)
		p.Match(PatitoParserLPAREN)
	}
	{
		p.SetState(115)
		p.PrintArgs()
	}
	{
		p.SetState(116)
		p.Match(PatitoParserRPAREN)
	}
	{
		p.SetState(117)
		p.Match(PatitoParserSEMI)
	}

	return localctx
}

// IPrintArgsContext is an interface to support dynamic dispatch.
type IPrintArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintArgsContext differentiates from other interfaces.
	IsPrintArgsContext()
}

type PrintArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintArgsContext() *PrintArgsContext {
	var p = new(PrintArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_printArgs
	return p
}

func (*PrintArgsContext) IsPrintArgsContext() {}

func NewPrintArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintArgsContext {
	var p = new(PrintArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_printArgs

	return p
}

func (s *PrintArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintArgsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PrintArgsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserCOMMA)
}

func (s *PrintArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserCOMMA, i)
}

func (s *PrintArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterPrintArgs(s)
	}
}

func (s *PrintArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitPrintArgs(s)
	}
}

func (s *PrintArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitPrintArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) PrintArgs() (localctx IPrintArgsContext) {
	this := p
	_ = this

	localctx = NewPrintArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PatitoParserRULE_printArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.expr(0)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserCOMMA {
		{
			p.SetState(120)
			p.Match(PatitoParserCOMMA)
		}
		{
			p.SetState(121)
			p.expr(0)
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) IF() antlr.TerminalNode {
	return s.GetToken(PatitoParserIF, 0)
}

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *ConditionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *ConditionContext) AllBody() []IBodyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBodyContext)(nil)).Elem())
	var tst = make([]IBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBodyContext)
		}
	}

	return tst
}

func (s *ConditionContext) Body(i int) IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *ConditionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PatitoParserELSE, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PatitoParserRULE_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(PatitoParserIF)
	}
	{
		p.SetState(128)
		p.Match(PatitoParserLPAREN)
	}
	{
		p.SetState(129)
		p.expr(0)
	}
	{
		p.SetState(130)
		p.Match(PatitoParserRPAREN)
	}
	{
		p.SetState(131)
		p.Body()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PatitoParserELSE {
		{
			p.SetState(132)
			p.Match(PatitoParserELSE)
		}
		{
			p.SetState(133)
			p.Body()
		}

	}

	return localctx
}

// ICycleContext is an interface to support dynamic dispatch.
type ICycleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCycleContext differentiates from other interfaces.
	IsCycleContext()
}

type CycleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCycleContext() *CycleContext {
	var p = new(CycleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_cycle
	return p
}

func (*CycleContext) IsCycleContext() {}

func NewCycleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CycleContext {
	var p = new(CycleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_cycle

	return p
}

func (s *CycleContext) GetParser() antlr.Parser { return s.parser }

func (s *CycleContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PatitoParserWHILE, 0)
}

func (s *CycleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *CycleContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CycleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *CycleContext) DO() antlr.TerminalNode {
	return s.GetToken(PatitoParserDO, 0)
}

func (s *CycleContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *CycleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *CycleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CycleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CycleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterCycle(s)
	}
}

func (s *CycleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitCycle(s)
	}
}

func (s *CycleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitCycle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Cycle() (localctx ICycleContext) {
	this := p
	_ = this

	localctx = NewCycleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PatitoParserRULE_cycle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(PatitoParserWHILE)
	}
	{
		p.SetState(137)
		p.Match(PatitoParserLPAREN)
	}
	{
		p.SetState(138)
		p.expr(0)
	}
	{
		p.SetState(139)
		p.Match(PatitoParserRPAREN)
	}
	{
		p.SetState(140)
		p.Match(PatitoParserDO)
	}
	{
		p.SetState(141)
		p.Body()
	}
	{
		p.SetState(142)
		p.Match(PatitoParserSEMI)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *ExprContext) Fcall() IFcallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFcallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFcallContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *ExprContext) Sign() ISignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *ExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprContext) RELOP() antlr.TerminalNode {
	return s.GetToken(PatitoParserRELOP, 0)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PatitoParserPLUS, 0)
}

func (s *ExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PatitoParserMINUS, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(PatitoParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(PatitoParserDIV, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PatitoParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, PatitoParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(145)
			p.Match(PatitoParserLPAREN)
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(PatitoParserRPAREN)
		}

	case 2:
		{
			p.SetState(149)
			p.Fcall()
		}

	case 3:
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PatitoParserPLUS || _la == PatitoParserMINUS {
			{
				p.SetState(150)
				p.Sign()
			}

		}
		{
			p.SetState(153)
			p.Match(PatitoParserID)
		}

	case 4:
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PatitoParserPLUS || _la == PatitoParserMINUS {
			{
				p.SetState(154)
				p.Sign()
			}

		}
		{
			p.SetState(157)
			p.Literal()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PatitoParserRULE_expr)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(161)
					p.Match(PatitoParserRELOP)
				}
				{
					p.SetState(162)
					p.expr(8)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PatitoParserRULE_expr)
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(164)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PatitoParserPLUS || _la == PatitoParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(165)
					p.expr(7)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PatitoParserRULE_expr)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(167)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PatitoParserMULT || _la == PatitoParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(168)
					p.expr(6)
				}

			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_sign
	return p
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PatitoParserPLUS, 0)
}

func (s *SignContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PatitoParserMINUS, 0)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterSign(s)
	}
}

func (s *SignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitSign(s)
	}
}

func (s *SignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Sign() (localctx ISignContext) {
	this := p
	_ = this

	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PatitoParserRULE_sign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PatitoParserPLUS || _la == PatitoParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(PatitoParserINT_LIT, 0)
}

func (s *LiteralContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(PatitoParserFLOAT_LIT, 0)
}

func (s *LiteralContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PatitoParserSTRING_LIT, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PatitoParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PatitoParserINT_LIT)|(1<<PatitoParserFLOAT_LIT)|(1<<PatitoParserSTRING_LIT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_funcDecl
	return p
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) VOID() antlr.TerminalNode {
	return s.GetToken(PatitoParserVOID, 0)
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *FuncDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *FuncDeclContext) ParamListOpt() IParamListOptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListOptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListOptContext)
}

func (s *FuncDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *FuncDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PatitoParserLBRACE, 0)
}

func (s *FuncDeclContext) VarsOpt() IVarsOptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsOptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsOptContext)
}

func (s *FuncDeclContext) Body() IBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *FuncDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PatitoParserRBRACE, 0)
}

func (s *FuncDeclContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PatitoParserSEMI, 0)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) FuncDecl() (localctx IFuncDeclContext) {
	this := p
	_ = this

	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PatitoParserRULE_funcDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(PatitoParserVOID)
	}
	{
		p.SetState(179)
		p.Match(PatitoParserID)
	}
	{
		p.SetState(180)
		p.Match(PatitoParserLPAREN)
	}
	{
		p.SetState(181)
		p.ParamListOpt()
	}
	{
		p.SetState(182)
		p.Match(PatitoParserRPAREN)
	}
	{
		p.SetState(183)
		p.Match(PatitoParserLBRACE)
	}
	{
		p.SetState(184)
		p.VarsOpt()
	}
	{
		p.SetState(185)
		p.Body()
	}
	{
		p.SetState(186)
		p.Match(PatitoParserRBRACE)
	}
	{
		p.SetState(187)
		p.Match(PatitoParserSEMI)
	}

	return localctx
}

// IParamListOptContext is an interface to support dynamic dispatch.
type IParamListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListOptContext differentiates from other interfaces.
	IsParamListOptContext()
}

type ParamListOptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListOptContext() *ParamListOptContext {
	var p = new(ParamListOptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_paramListOpt
	return p
}

func (*ParamListOptContext) IsParamListOptContext() {}

func NewParamListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListOptContext {
	var p = new(ParamListOptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_paramListOpt

	return p
}

func (s *ParamListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListOptContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParamListOpt(s)
	}
}

func (s *ParamListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParamListOpt(s)
	}
}

func (s *ParamListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParamListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ParamListOpt() (localctx IParamListOptContext) {
	this := p
	_ = this

	localctx = NewParamListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PatitoParserRULE_paramListOpt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PatitoParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.ParamList()
		}

	case PatitoParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserCOMMA)
}

func (s *ParamListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserCOMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ParamList() (localctx IParamListContext) {
	this := p
	_ = this

	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PatitoParserRULE_paramList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Param()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserCOMMA {
		{
			p.SetState(194)
			p.Match(PatitoParserCOMMA)
		}
		{
			p.SetState(195)
			p.Param()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *ParamContext) COLON() antlr.TerminalNode {
	return s.GetToken(PatitoParserCOLON, 0)
}

func (s *ParamContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Param() (localctx IParamContext) {
	this := p
	_ = this

	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PatitoParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(PatitoParserID)
	}
	{
		p.SetState(202)
		p.Match(PatitoParserCOLON)
	}
	{
		p.SetState(203)
		p.TypeSpec()
	}

	return localctx
}

// IFcallContext is an interface to support dynamic dispatch.
type IFcallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFcallContext differentiates from other interfaces.
	IsFcallContext()
}

type FcallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFcallContext() *FcallContext {
	var p = new(FcallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_fcall
	return p
}

func (*FcallContext) IsFcallContext() {}

func NewFcallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FcallContext {
	var p = new(FcallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_fcall

	return p
}

func (s *FcallContext) GetParser() antlr.Parser { return s.parser }

func (s *FcallContext) ID() antlr.TerminalNode {
	return s.GetToken(PatitoParserID, 0)
}

func (s *FcallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserLPAREN, 0)
}

func (s *FcallContext) ArgListOpt() IArgListOptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListOptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListOptContext)
}

func (s *FcallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PatitoParserRPAREN, 0)
}

func (s *FcallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FcallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FcallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterFcall(s)
	}
}

func (s *FcallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitFcall(s)
	}
}

func (s *FcallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitFcall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) Fcall() (localctx IFcallContext) {
	this := p
	_ = this

	localctx = NewFcallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PatitoParserRULE_fcall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(PatitoParserID)
	}
	{
		p.SetState(206)
		p.Match(PatitoParserLPAREN)
	}
	{
		p.SetState(207)
		p.ArgListOpt()
	}
	{
		p.SetState(208)
		p.Match(PatitoParserRPAREN)
	}

	return localctx
}

// IArgListOptContext is an interface to support dynamic dispatch.
type IArgListOptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListOptContext differentiates from other interfaces.
	IsArgListOptContext()
}

type ArgListOptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListOptContext() *ArgListOptContext {
	var p = new(ArgListOptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_argListOpt
	return p
}

func (*ArgListOptContext) IsArgListOptContext() {}

func NewArgListOptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListOptContext {
	var p = new(ArgListOptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_argListOpt

	return p
}

func (s *ArgListOptContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListOptContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *ArgListOptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListOptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListOptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterArgListOpt(s)
	}
}

func (s *ArgListOptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitArgListOpt(s)
	}
}

func (s *ArgListOptContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitArgListOpt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ArgListOpt() (localctx IArgListOptContext) {
	this := p
	_ = this

	localctx = NewArgListOptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PatitoParserRULE_argListOpt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PatitoParserPLUS, PatitoParserMINUS, PatitoParserLPAREN, PatitoParserID, PatitoParserINT_LIT, PatitoParserFLOAT_LIT, PatitoParserSTRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.ArgList()
		}

	case PatitoParserRPAREN:
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatitoParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatitoParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ArgListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PatitoParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PatitoParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatitoListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PatitoVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PatitoParser) ArgList() (localctx IArgListContext) {
	this := p
	_ = this

	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PatitoParserRULE_argList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.expr(0)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PatitoParserCOMMA {
		{
			p.SetState(215)
			p.Match(PatitoParserCOMMA)
		}
		{
			p.SetState(216)
			p.expr(0)
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

func (p *PatitoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PatitoParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
