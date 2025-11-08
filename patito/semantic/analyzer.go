package semantic

import (
	"fmt"
	"patito/gen"
)

// SemanticAnalyzer recorre el árbol de parsing de ANTLR
// y construye las tablas de símbolos mientras valida el programa
type SemanticAnalyzer struct {
	*gen.BasePatitoListener
	vm            *VariableManager   // Maneja variables globales y locales
	fd            *FunctionDirectory // Directorio de funciones
	errors        []string           // Errores semánticos encontrados
	currentFunc   *FunctionSignature // Función actual (nil si estamos en scope global)
	inGlobalScope bool               // true cuando estamos procesando vars globales
	programName   string             // Nombre del programa
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{
		BasePatitoListener: &gen.BasePatitoListener{},
		vm:                 NewVariableManager(),
		fd:                 NewFunctionDirectory(),
		errors:             make([]string, 0),
		currentFunc:        nil,
		inGlobalScope:      false,
	}
}

func (s *SemanticAnalyzer) GetErrors() []string {
	return s.errors
}

func (s *SemanticAnalyzer) HasErrors() bool {
	return len(s.errors) > 0
}

func (s *SemanticAnalyzer) GetVariableManager() *VariableManager {
	return s.vm
}

func (s *SemanticAnalyzer) GetFunctionDirectory() *FunctionDirectory {
	return s.fd
}

func (s *SemanticAnalyzer) addError(msg string) {
	s.errors = append(s.errors, msg)
}

// ========== Programa ==========

// Al entrar al programa, guardamos el nombre y marcamos que estamos en scope global
func (s *SemanticAnalyzer) EnterProgram(ctx *gen.ProgramContext) {
	s.programName = ctx.ID().GetText()
	s.inGlobalScope = true
}

func (s *SemanticAnalyzer) ExitProgram(ctx *gen.ProgramContext) {
	s.inGlobalScope = false
}

// ========== Declaración de Variables ==========

// PUNTO NEURÁLGICO: Aquí agregamos variables a las tablas
// Ejemplo: x, y : int ;
func (s *SemanticAnalyzer) ExitVarGroup(ctx *gen.VarGroupContext) {
	varNames := s.getIdList(ctx.IdList())
	varType := s.getTypeSpec(ctx.TypeSpec())

	if s.currentFunc != nil {
		// Estamos dentro de una función - agregar variables locales
		for _, name := range varNames {
			err := s.currentFunc.AddLocalVariable(name, varType)
			if err != nil {
				s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
			}
			// También agregar al VariableManager para búsquedas
			s.vm.AddLocalVariable(name, varType)
		}
	} else if s.inGlobalScope {
		// Estamos en scope global - agregar variables globales
		for _, name := range varNames {
			err := s.vm.AddGlobalVariable(name, varType)
			if err != nil {
				s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
			}
		}
	}
}

// ========== Declaración de Funciones ==========

// PUNTO NEURÁLGICO: Al entrar a una función, la registramos y creamos su scope local
// Ejemplo: void foo(a : int) { ... }
func (s *SemanticAnalyzer) EnterFuncDecl(ctx *gen.FuncDeclContext) {
	funcName := ctx.ID().GetText()

	// Intentar agregar la función al directorio
	funcSig, err := s.fd.AddFunction(funcName, VOID)
	if err != nil {
		s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
		// Crear una firma dummy para continuar el análisis
		funcSig = NewFunctionSignature(funcName, VOID)
	}

	s.currentFunc = funcSig
	s.inGlobalScope = false
	s.vm.EnterFunction()
}

// Al salir de la función, limpiamos el scope local
func (s *SemanticAnalyzer) ExitFuncDecl(ctx *gen.FuncDeclContext) {
	s.currentFunc = nil
	s.vm.ExitFunction()
}

// ========== Parámetros de Funciones ==========

// PUNTO NEURÁLGICO: Agregar parámetros a la función actual
// Ejemplo: a : int
func (s *SemanticAnalyzer) ExitParam(ctx *gen.ParamContext) {
	if s.currentFunc == nil {
		s.addError(fmt.Sprintf("Line %d: parameter outside of function", ctx.GetStart().GetLine()))
		return
	}

	paramName := ctx.ID().GetText()
	paramType := s.getTypeSpec(ctx.TypeSpec())

	err := s.currentFunc.AddParameter(paramName, paramType)
	if err != nil {
		s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
	}

	// También agregar al VariableManager
	s.vm.AddLocalVariable(paramName, paramType)
}

// ========== Asignaciones ==========

// PUNTO NEURÁLGICO: Validar que la variable existe antes de asignarle un valor
// Ejemplo: x = 10 ;
func (s *SemanticAnalyzer) ExitAssign(ctx *gen.AssignContext) {
	varName := ctx.ID().GetText()

	_, err := s.vm.Lookup(varName)
	if err != nil {
		s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
	}
}

// ========== Llamadas a Funciones ==========

// PUNTO NEURÁLGICO: Validar que la función existe y los argumentos son correctos
// Ejemplo: foo(10, 3.14) ;
func (s *SemanticAnalyzer) ExitFcall(ctx *gen.FcallContext) {
	funcName := ctx.ID().GetText()
	argTypes := s.getArgTypes(ctx.ArgListOpt())

	err := s.fd.ValidateFunctionCall(funcName, argTypes)
	if err != nil {
		s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
	}
}

// ========== Validación de Expresiones ==========

// Validar que las variables usadas en expresiones estén declaradas
// Ejemplo: x + y
func (s *SemanticAnalyzer) EnterExpr(ctx *gen.ExprContext) {
	if ctx.ID() != nil {
		varName := ctx.ID().GetText()
		_, err := s.vm.Lookup(varName)
		if err != nil {
			s.addError(fmt.Sprintf("Line %d: %v", ctx.GetStart().GetLine(), err))
		}
	}
}

// ========== Funciones Auxiliares ==========

// Extrae la lista de identificadores de una declaración
// Ejemplo: x, y, z  →  ["x", "y", "z"]
func (s *SemanticAnalyzer) getIdList(ctx gen.IIdListContext) []string {
	if ctx == nil {
		return []string{}
	}

	idListCtx, ok := ctx.(*gen.IdListContext)
	if !ok {
		return []string{}
	}

	ids := idListCtx.AllID()
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = id.GetText()
	}
	return result
}

// Extrae el tipo de una declaración
// Ejemplo: int  →  INT
func (s *SemanticAnalyzer) getTypeSpec(ctx gen.ITypeSpecContext) Type {
	if ctx == nil {
		return ERROR
	}

	typeSpecCtx, ok := ctx.(*gen.TypeSpecContext)
	if !ok {
		return ERROR
	}

	if typeSpecCtx.INT_KW() != nil {
		return INT
	} else if typeSpecCtx.FLOAT_KW() != nil {
		return FLOAT
	}

	return ERROR
}

// Obtiene la cantidad de argumentos en una llamada a función
// Por ahora solo retorna tipos INT para verificar la cantidad de argumentos
// (la inferencia completa de tipos se haría en fases posteriores del compilador)
func (s *SemanticAnalyzer) getArgTypes(ctx gen.IArgListOptContext) []Type {
	if ctx == nil {
		return []Type{}
	}

	argListOptCtx, ok := ctx.(*gen.ArgListOptContext)
	if !ok {
		return []Type{}
	}

	argList := argListOptCtx.ArgList()
	if argList == nil {
		return []Type{}
	}

	argListCtx, ok := argList.(*gen.ArgListContext)
	if !ok {
		return []Type{}
	}

	exprs := argListCtx.AllExpr()
	types := make([]Type, len(exprs))

	// Por ahora asumimos INT para contar argumentos
	for i := range exprs {
		types[i] = INT
	}

	return types
}

