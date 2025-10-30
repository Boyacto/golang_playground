// Code generated from Patito.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gen // Patito
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePatitoVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePatitoVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitVarsOpt(ctx *VarsOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitFuncsOpt(ctx *FuncsOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitVarsDecl(ctx *VarsDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitVarGroup(ctx *VarGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitIdList(ctx *IdListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitPrintArgs(ctx *PrintArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitCycle(ctx *CycleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitSign(ctx *SignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParamListOpt(ctx *ParamListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitFcall(ctx *FcallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitArgListOpt(ctx *ArgListOptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePatitoVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}
