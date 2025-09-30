package main // El paquete principal: punto de entrada del programa

// ===== Stack (LIFO) =====
type Stack[T any] struct { // Definimos un tipo genérico Stack con parámetro de tipo T (cualquier tipo)
	data []T // Usamos un slice de T para almacenar los elementos del stack
}

// comentar para debugear otros
func (s *Stack[T]) Push(v T) { // Método Push: inserta un elemento en la cima
	s.data = append(s.data, v) // append agrega al final del slice, creciendo si es necesario
}
func (s *Stack[T]) Pop() (T, bool) { // Método Pop: quita y devuelve el último elemento. bool indica éxito
	var zero T            // zero contendrá el valor cero para T por si el stack está vacío
	if len(s.data) == 0 { // Si no hay elementos
		return zero, false // Devolvemos el cero de T y false
	}
	v := s.data[len(s.data)-1]      // Obtenemos el último elemento (cima)
	s.data = s.data[:len(s.data)-1] // Reducimos el slice para eliminar el último elemento
	return v, true                  // Devolvemos el valor extraído y true
}

func (s *Stack[T]) Peek() (T, bool) { // Método Peek: mira el último elemento sin quitarlo
	var zero T            // Valor cero de T para el caso vacío
	if len(s.data) == 0 { // Si está vacío
		return zero, false // Indicamos que no hay elemento
	}
	return s.data[len(s.data)-1], true // Devolvemos la cima y true
}

func (s *Stack[T]) Len() int      { return len(s.data) }      // Len: cantidad de elementos actuales
func (s *Stack[T]) IsEmpty() bool { return len(s.data) == 0 } // IsEmpty: true si no hay elementos
func (s *Stack[T]) Clear()        { s.data = nil }            // Clear: libera la referencia al slice para que el GC recupere memoria
