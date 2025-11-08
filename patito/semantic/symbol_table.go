package semantic

import "fmt"

// ScopeType representa el nivel de alcance de una variable
type ScopeType int

const (
	GLOBAL ScopeType = iota
	LOCAL
)

func (s ScopeType) String() string {
	switch s {
	case GLOBAL:
		return "global"
	case LOCAL:
		return "local"
	default:
		return "unknown"
	}
}

// VariableSymbol representa una variable en la tabla de símbolos
type VariableSymbol struct {
	Name    string
	Type    Type
	Scope   ScopeType
	Address int // Dirección de memoria virtual (para generación de código)
}

// VariableTable representa una tabla de símbolos para variables
type VariableTable struct {
	symbols map[string]*VariableSymbol
	scope   ScopeType
	nextAddress int
}

// NewVariableTable crea una nueva tabla de variables para el alcance especificado
func NewVariableTable(scope ScopeType) *VariableTable {
	return &VariableTable{
		symbols: make(map[string]*VariableSymbol),
		scope:   scope,
		nextAddress: 0,
	}
}

// AddVariable agrega una nueva variable a la tabla
func (vt *VariableTable) AddVariable(name string, varType Type) error {
	// Verificar si la variable ya existe
	if _, exists := vt.symbols[name]; exists {
		return fmt.Errorf("variable '%s' already declared in %s scope", name, vt.scope)
	}

	// Validar el tipo (solo int y float para variables)
	if err := ValidateType(varType); err != nil {
		return err
	}

	// Crear el símbolo de la variable
	symbol := &VariableSymbol{
		Name:    name,
		Type:    varType,
		Scope:   vt.scope,
		Address: vt.nextAddress,
	}

	vt.symbols[name] = symbol
	vt.nextAddress++

	return nil
}

// Lookup busca una variable en la tabla
func (vt *VariableTable) Lookup(name string) (*VariableSymbol, bool) {
	symbol, exists := vt.symbols[name]
	return symbol, exists
}

// GetAll retorna todas las variables de la tabla
func (vt *VariableTable) GetAll() map[string]*VariableSymbol {
	return vt.symbols
}

// Size retorna el número de variables en la tabla
func (vt *VariableTable) Size() int {
	return len(vt.symbols)
}

// GetNextAddress retorna la siguiente dirección de memoria disponible
func (vt *VariableTable) GetNextAddress() int {
	return vt.nextAddress
}

// VariableManager maneja tanto la tabla de variables globales como locales
type VariableManager struct {
	globalTable *VariableTable
	localTable  *VariableTable
}

// NewVariableManager crea un nuevo manejador de variables
func NewVariableManager() *VariableManager {
	return &VariableManager{
		globalTable: NewVariableTable(GLOBAL),
		localTable:  nil, // La tabla local se crea al entrar a una función
	}
}

// EnterFunction crea un nuevo alcance local (al entrar a una función)
func (vm *VariableManager) EnterFunction() {
	vm.localTable = NewVariableTable(LOCAL)
}

// ExitFunction limpia el alcance local (al salir de una función)
func (vm *VariableManager) ExitFunction() {
	vm.localTable = nil
}

// AddGlobalVariable agrega una variable al alcance global
func (vm *VariableManager) AddGlobalVariable(name string, varType Type) error {
	return vm.globalTable.AddVariable(name, varType)
}

// AddLocalVariable agrega una variable al alcance local
func (vm *VariableManager) AddLocalVariable(name string, varType Type) error {
	if vm.localTable == nil {
		return fmt.Errorf("cannot add local variable outside of a function")
	}
	return vm.localTable.AddVariable(name, varType)
}

// Lookup busca una variable (primero en local, luego en global)
func (vm *VariableManager) Lookup(name string) (*VariableSymbol, error) {
	// Primero revisar el alcance local si existe
	if vm.localTable != nil {
		if symbol, exists := vm.localTable.Lookup(name); exists {
			return symbol, nil
		}
	}

	// Luego revisar el alcance global
	if symbol, exists := vm.globalTable.Lookup(name); exists {
		return symbol, nil
	}

	return nil, fmt.Errorf("variable '%s' not declared", name)
}

// GetGlobalTable retorna la tabla de variables globales
func (vm *VariableManager) GetGlobalTable() *VariableTable {
	return vm.globalTable
}

// GetLocalTable retorna la tabla de variables locales (puede ser nil)
func (vm *VariableManager) GetLocalTable() *VariableTable {
	return vm.localTable
}
