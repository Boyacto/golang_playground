// Code generated from Patito.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gen // Patito
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PatitoParser.
type PatitoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PatitoParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by PatitoParser#varsOpt.
	VisitVarsOpt(ctx *VarsOptContext) interface{}

	// Visit a parse tree produced by PatitoParser#funcsOpt.
	VisitFuncsOpt(ctx *FuncsOptContext) interface{}

	// Visit a parse tree produced by PatitoParser#varsDecl.
	VisitVarsDecl(ctx *VarsDeclContext) interface{}

	// Visit a parse tree produced by PatitoParser#varGroup.
	VisitVarGroup(ctx *VarGroupContext) interface{}

	// Visit a parse tree produced by PatitoParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by PatitoParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by PatitoParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by PatitoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PatitoParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by PatitoParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by PatitoParser#printArgs.
	VisitPrintArgs(ctx *PrintArgsContext) interface{}

	// Visit a parse tree produced by PatitoParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by PatitoParser#cycle.
	VisitCycle(ctx *CycleContext) interface{}

	// Visit a parse tree produced by PatitoParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by PatitoParser#sign.
	VisitSign(ctx *SignContext) interface{}

	// Visit a parse tree produced by PatitoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PatitoParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by PatitoParser#paramListOpt.
	VisitParamListOpt(ctx *ParamListOptContext) interface{}

	// Visit a parse tree produced by PatitoParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by PatitoParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by PatitoParser#fcall.
	VisitFcall(ctx *FcallContext) interface{}

	// Visit a parse tree produced by PatitoParser#argListOpt.
	VisitArgListOpt(ctx *ArgListOptContext) interface{}

	// Visit a parse tree produced by PatitoParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}
}
