package main

import (
	"fmt"
)

// ===== Stack (LIFO) =====
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	//tengo que encontrar de alguna forma como append para agregar el elemento
}
func (s *Stack[T]) Pop() (T, bool) {
	//tengo que checar si esta en 0 y si no está vacio quitar la longitud y regresar el valor que hice pop
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
