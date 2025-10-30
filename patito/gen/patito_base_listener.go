// Code generated from Patito.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gen // Patito
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePatitoListener is a complete listener for a parse tree produced by PatitoParser.
type BasePatitoListener struct{}

var _ PatitoListener = &BasePatitoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePatitoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePatitoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePatitoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePatitoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePatitoListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePatitoListener) ExitProgram(ctx *ProgramContext) {}

// EnterVarsOpt is called when production varsOpt is entered.
func (s *BasePatitoListener) EnterVarsOpt(ctx *VarsOptContext) {}

// ExitVarsOpt is called when production varsOpt is exited.
func (s *BasePatitoListener) ExitVarsOpt(ctx *VarsOptContext) {}

// EnterFuncsOpt is called when production funcsOpt is entered.
func (s *BasePatitoListener) EnterFuncsOpt(ctx *FuncsOptContext) {}

// ExitFuncsOpt is called when production funcsOpt is exited.
func (s *BasePatitoListener) ExitFuncsOpt(ctx *FuncsOptContext) {}

// EnterVarsDecl is called when production varsDecl is entered.
func (s *BasePatitoListener) EnterVarsDecl(ctx *VarsDeclContext) {}

// ExitVarsDecl is called when production varsDecl is exited.
func (s *BasePatitoListener) ExitVarsDecl(ctx *VarsDeclContext) {}

// EnterVarGroup is called when production varGroup is entered.
func (s *BasePatitoListener) EnterVarGroup(ctx *VarGroupContext) {}

// ExitVarGroup is called when production varGroup is exited.
func (s *BasePatitoListener) ExitVarGroup(ctx *VarGroupContext) {}

// EnterIdList is called when production idList is entered.
func (s *BasePatitoListener) EnterIdList(ctx *IdListContext) {}

// ExitIdList is called when production idList is exited.
func (s *BasePatitoListener) ExitIdList(ctx *IdListContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BasePatitoListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BasePatitoListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterBody is called when production body is entered.
func (s *BasePatitoListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BasePatitoListener) ExitBody(ctx *BodyContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePatitoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePatitoListener) ExitStatement(ctx *StatementContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasePatitoListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasePatitoListener) ExitAssign(ctx *AssignContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BasePatitoListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BasePatitoListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterPrintArgs is called when production printArgs is entered.
func (s *BasePatitoListener) EnterPrintArgs(ctx *PrintArgsContext) {}

// ExitPrintArgs is called when production printArgs is exited.
func (s *BasePatitoListener) ExitPrintArgs(ctx *PrintArgsContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasePatitoListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasePatitoListener) ExitCondition(ctx *ConditionContext) {}

// EnterCycle is called when production cycle is entered.
func (s *BasePatitoListener) EnterCycle(ctx *CycleContext) {}

// ExitCycle is called when production cycle is exited.
func (s *BasePatitoListener) ExitCycle(ctx *CycleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePatitoListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePatitoListener) ExitExpr(ctx *ExprContext) {}

// EnterSign is called when production sign is entered.
func (s *BasePatitoListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BasePatitoListener) ExitSign(ctx *SignContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePatitoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePatitoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BasePatitoListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BasePatitoListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterParamListOpt is called when production paramListOpt is entered.
func (s *BasePatitoListener) EnterParamListOpt(ctx *ParamListOptContext) {}

// ExitParamListOpt is called when production paramListOpt is exited.
func (s *BasePatitoListener) ExitParamListOpt(ctx *ParamListOptContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BasePatitoListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BasePatitoListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BasePatitoListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasePatitoListener) ExitParam(ctx *ParamContext) {}

// EnterFcall is called when production fcall is entered.
func (s *BasePatitoListener) EnterFcall(ctx *FcallContext) {}

// ExitFcall is called when production fcall is exited.
func (s *BasePatitoListener) ExitFcall(ctx *FcallContext) {}

// EnterArgListOpt is called when production argListOpt is entered.
func (s *BasePatitoListener) EnterArgListOpt(ctx *ArgListOptContext) {}

// ExitArgListOpt is called when production argListOpt is exited.
func (s *BasePatitoListener) ExitArgListOpt(ctx *ArgListOptContext) {}

// EnterArgList is called when production argList is entered.
func (s *BasePatitoListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BasePatitoListener) ExitArgList(ctx *ArgListContext) {}
