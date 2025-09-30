package main

import (
	"fmt"
)

// ===== Stack (LIFO) =====
type Stack[T any] struct {
	data []T
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
