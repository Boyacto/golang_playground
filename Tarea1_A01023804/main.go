package main // Declaramos el paquete principal: punto de entrada del programa

import ( // Bloque de importaciones
	"fmt" // Paquete para imprimir por consola y formatear cadenas
)

// ===== Stack (LIFO) =====
type Stack[T any] struct { // Definimos un tipo genérico Stack con parámetro de tipo T (cualquier tipo)
	data []T // Usamos un slice de T para almacenar los elementos del stack
}

func (s *Stack[T]) Push(v T) { // Método Push: inserta un elemento en la cima
	// append debe de agregar al final del slice, creciendo si es necesario
}
func (s *Stack[T]) Pop() (T, bool) { // Método Pop: quita y devuelve el último elemento; bool indica éxito
	//tengo que checar si esta en 0 y si no está vacio quitar la longitud y regresar el valor que hice pop y valor bool si se logro regresar
}

// ===== Queue (FIFO) =====
type Queue[T any] struct {
	data []T
	head int
}

// ===== Dictionary (Map) =====
type Dict[K comparable, V any] struct {
	m map[K]V
}

// ===== Main function with demo =====
func main() {
	fmt.Println("main function")
	fmt.Println("Stack")
	fmt.Println("Queue")
	fmt.Println("Dictionary")
}
