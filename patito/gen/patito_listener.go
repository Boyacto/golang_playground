// Code generated from Patito.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gen // Patito
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PatitoListener is a complete listener for a parse tree produced by PatitoParser.
type PatitoListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterVarsOpt is called when entering the varsOpt production.
	EnterVarsOpt(c *VarsOptContext)

	// EnterFuncsOpt is called when entering the funcsOpt production.
	EnterFuncsOpt(c *FuncsOptContext)

	// EnterVarsDecl is called when entering the varsDecl production.
	EnterVarsDecl(c *VarsDeclContext)

	// EnterVarGroup is called when entering the varGroup production.
	EnterVarGroup(c *VarGroupContext)

	// EnterIdList is called when entering the idList production.
	EnterIdList(c *IdListContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterPrintArgs is called when entering the printArgs production.
	EnterPrintArgs(c *PrintArgsContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterCycle is called when entering the cycle production.
	EnterCycle(c *CycleContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterParamListOpt is called when entering the paramListOpt production.
	EnterParamListOpt(c *ParamListOptContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFcall is called when entering the fcall production.
	EnterFcall(c *FcallContext)

	// EnterArgListOpt is called when entering the argListOpt production.
	EnterArgListOpt(c *ArgListOptContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitVarsOpt is called when exiting the varsOpt production.
	ExitVarsOpt(c *VarsOptContext)

	// ExitFuncsOpt is called when exiting the funcsOpt production.
	ExitFuncsOpt(c *FuncsOptContext)

	// ExitVarsDecl is called when exiting the varsDecl production.
	ExitVarsDecl(c *VarsDeclContext)

	// ExitVarGroup is called when exiting the varGroup production.
	ExitVarGroup(c *VarGroupContext)

	// ExitIdList is called when exiting the idList production.
	ExitIdList(c *IdListContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitPrintArgs is called when exiting the printArgs production.
	ExitPrintArgs(c *PrintArgsContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitCycle is called when exiting the cycle production.
	ExitCycle(c *CycleContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitParamListOpt is called when exiting the paramListOpt production.
	ExitParamListOpt(c *ParamListOptContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFcall is called when exiting the fcall production.
	ExitFcall(c *FcallContext)

	// ExitArgListOpt is called when exiting the argListOpt production.
	ExitArgListOpt(c *ArgListOptContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)
}
