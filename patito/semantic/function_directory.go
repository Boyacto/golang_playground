package semantic

import "fmt"

// Parameter representa un parámetro de función
type Parameter struct {
	Name string
	Type Type
}

// FunctionSignature representa una función en el directorio
type FunctionSignature struct {
	Name           string
	ReturnType     Type // En Patito, siempre es VOID
	Parameters     []Parameter
	LocalVarTable  *VariableTable
	StartAddress   int // Dirección donde inicia el código de la función (para generación de código)
}

// NewFunctionSignature crea una nueva firma de función
func NewFunctionSignature(name string, returnType Type) *FunctionSignature {
	return &FunctionSignature{
		Name:          name,
		ReturnType:    returnType,
		Parameters:    make([]Parameter, 0),
		LocalVarTable: NewVariableTable(LOCAL),
		StartAddress:  -1, // Se asignará durante la generación de código
	}
}

// AddParameter agrega un parámetro a la función
func (fs *FunctionSignature) AddParameter(name string, paramType Type) error {
	// Verificar si el parámetro ya existe
	for _, param := range fs.Parameters {
		if param.Name == name {
			return fmt.Errorf("parameter '%s' already declared in function '%s'", name, fs.Name)
		}
	}

	// Validar el tipo
	if err := ValidateType(paramType); err != nil {
		return fmt.Errorf("invalid parameter type for '%s': %v", name, err)
	}

	// Agregar el parámetro
	fs.Parameters = append(fs.Parameters, Parameter{
		Name: name,
		Type: paramType,
	})

	// También agregar a la tabla de variables locales (los parámetros son variables locales)
	return fs.LocalVarTable.AddVariable(name, paramType)
}

// AddLocalVariable agrega una variable local a la función
func (fs *FunctionSignature) AddLocalVariable(name string, varType Type) error {
	// Verificar si el nombre entra en conflicto con los parámetros
	for _, param := range fs.Parameters {
		if param.Name == name {
			return fmt.Errorf("variable '%s' conflicts with parameter name in function '%s'", name, fs.Name)
		}
	}

	return fs.LocalVarTable.AddVariable(name, varType)
}

// GetParameterCount retorna el número de parámetros
func (fs *FunctionSignature) GetParameterCount() int {
	return len(fs.Parameters)
}

// ValidateCall verifica si una llamada a función coincide con la firma
// PUNTO NEURÁLGICO: Validación de tipos de argumentos
func (fs *FunctionSignature) ValidateCall(argTypes []Type) error {
	if len(argTypes) != len(fs.Parameters) {
		return fmt.Errorf("function '%s' expects %d arguments but got %d",
			fs.Name, len(fs.Parameters), len(argTypes))
	}

	// Verificación de tipos de argumentos
	for i, param := range fs.Parameters {
		if !IsCompatible(argTypes[i], param.Type) {
			return fmt.Errorf("argument %d of function '%s': type mismatch (expected %s, got %s)",
				i+1, fs.Name, param.Type, argTypes[i])
		}
	}

	return nil
}

// FunctionDirectory almacena todas las firmas de funciones
type FunctionDirectory struct {
	functions map[string]*FunctionSignature
}

// NewFunctionDirectory crea un nuevo directorio de funciones
func NewFunctionDirectory() *FunctionDirectory {
	return &FunctionDirectory{
		functions: make(map[string]*FunctionSignature),
	}
}

// AddFunction agrega una nueva función al directorio
func (fd *FunctionDirectory) AddFunction(name string, returnType Type) (*FunctionSignature, error) {
	// Verificar si la función ya existe
	if _, exists := fd.functions[name]; exists {
		return nil, fmt.Errorf("function '%s' already declared", name)
	}

	// En Patito, todas las funciones deben ser void
	if returnType != VOID {
		return nil, fmt.Errorf("function '%s' must have void return type", name)
	}

	// Crear la firma de función
	funcSig := NewFunctionSignature(name, returnType)
	fd.functions[name] = funcSig

	return funcSig, nil
}

// Lookup busca una función en el directorio
func (fd *FunctionDirectory) Lookup(name string) (*FunctionSignature, bool) {
	funcSig, exists := fd.functions[name]
	return funcSig, exists
}

// GetAll retorna todas las funciones del directorio
func (fd *FunctionDirectory) GetAll() map[string]*FunctionSignature {
	return fd.functions
}

// Size retorna el número de funciones en el directorio
func (fd *FunctionDirectory) Size() int {
	return len(fd.functions)
}

// ValidateFunctionCall valida una llamada a función
func (fd *FunctionDirectory) ValidateFunctionCall(name string, argTypes []Type) error {
	funcSig, exists := fd.Lookup(name)
	if !exists {
		return fmt.Errorf("function '%s' not declared", name)
	}

	return funcSig.ValidateCall(argTypes)
}
