package semantic

import "fmt"

// Type representa los tipos de datos soportados en Patito
type Type int

const (
	INT Type = iota
	FLOAT
	VOID
	ERROR // Se usa para errores de verificación de tipos
)

// String retorna la representación en string de un Type
func (t Type) String() string {
	switch t {
	case INT:
		return "int"
	case FLOAT:
		return "float"
	case VOID:
		return "void"
	case ERROR:
		return "error"
	default:
		return "unknown"
	}
}

// TypeFromString convierte un string a un Type
func TypeFromString(s string) Type {
	switch s {
	case "int":
		return INT
	case "float":
		return FLOAT
	case "void":
		return VOID
	default:
		return ERROR
	}
}

// SemanticCube - Cubo Semántico
// Estructura de 3 dimensiones: operator × leftType × rightType → resultType
//
// Ejemplo de uso:
//   result := SemanticCube["+"][INT][FLOAT]  // retorna FLOAT
//
// Operadores soportados:
//   Aritméticos: +, -, *, /
//   Relacionales: <, >, <=, >=, ==, !=
var SemanticCube = map[string]map[Type]map[Type]Type{
	// Operadores aritméticos
	"+": {
		INT: {
			INT:   INT,
			FLOAT: FLOAT,
		},
		FLOAT: {
			INT:   FLOAT,
			FLOAT: FLOAT,
		},
	},
	"-": {
		INT: {
			INT:   INT,
			FLOAT: FLOAT,
		},
		FLOAT: {
			INT:   FLOAT,
			FLOAT: FLOAT,
		},
	},
	"*": {
		INT: {
			INT:   INT,
			FLOAT: FLOAT,
		},
		FLOAT: {
			INT:   FLOAT,
			FLOAT: FLOAT,
		},
	},
	"/": {
		INT: {
			INT:   INT,
			FLOAT: FLOAT,
		},
		FLOAT: {
			INT:   FLOAT,
			FLOAT: FLOAT,
		},
	},
	// Operadores relacionales (siempre retornan INT como booleano)
	"<": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
	">": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
	"<=": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
	">=": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
	"==": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
	"!=": {
		INT: {
			INT:   INT,
			FLOAT: INT,
		},
		FLOAT: {
			INT:   INT,
			FLOAT: INT,
		},
	},
}

// GetOperationType consulta el cubo semántico para obtener el tipo resultado
// Retorna ERROR si la operación no es válida
func GetOperationType(operator string, leftType, rightType Type) Type {
	if opMap, exists := SemanticCube[operator]; exists {
		if leftMap, exists := opMap[leftType]; exists {
			if resultType, exists := leftMap[rightType]; exists {
				return resultType
			}
		}
	}
	return ERROR
}

// IsCompatible verifica si dos tipos son compatibles para operaciones
func IsCompatible(t1, t2 Type) bool {
	// Los mismos tipos siempre son compatibles
	if t1 == t2 {
		return true
	}

	// INT y FLOAT son compatibles (con conversión implícita)
	if (t1 == INT || t1 == FLOAT) && (t2 == INT || t2 == FLOAT) {
		return true
	}

	return false
}

// ResultType retorna el tipo resultante de una operación entre dos tipos
// Este método se mantiene para compatibilidad, pero ahora usa el cubo semántico
func ResultType(t1, t2 Type) Type {
	if !IsCompatible(t1, t2) {
		return ERROR
	}

	// Si alguno es FLOAT, el resultado es FLOAT
	if t1 == FLOAT || t2 == FLOAT {
		return FLOAT
	}

	// Ambos son INT
	return INT
}

// ValidateType verifica si un tipo es válido para declaraciones de variables
func ValidateType(t Type) error {
	if t != INT && t != FLOAT {
		return fmt.Errorf("invalid variable type: %s (only int and float are allowed)", t)
	}
	return nil
}
