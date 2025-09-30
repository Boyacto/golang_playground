package main // Declaramos el paquete principal: punto de entrada del programa

import ( // Bloque de importaciones
	"fmt" // Paquete para imprimir por consola y formatear cadenas
)

// ===== Main function with demo =====
func main() {
	fmt.Println("main function")
	fmt.Println("Stack")
	// Stack demo
	var s Stack[int]                // Declaramos un Stack de enteros (inicia con slice nil pero usable)
	s.Push(10)                      // Insertamos 10
	s.Push(20)                      // Insertamos 20
	top, _ := s.Peek()              // Vemos el tope sin quitarlo
	fmt.Println("stack peek:", top) // Imprimimos el tope (esperado 20)
	for !s.IsEmpty() {              // Mientras el stack no esté vacío
		v, _ := s.Pop()              // Sacamos el elemento superior
		fmt.Println("stack pop:", v) // Mostramos el elemento sacado
	}

	fmt.Println("Queue")
	// Queue demo
	var q Queue[string]               // Declaramos una Queue de strings
	q.Enqueue("a")                    // Enqueue "a"
	q.Enqueue("b")                    // Enqueue "b"
	front, _ := q.Peek()              // Miramos el frente sin extraer
	fmt.Println("queue peek:", front) // Imprimimos el frente (esperado "a")
	for !q.IsEmpty() {                // Mientras haya elementos en la cola
		v, _ := q.Dequeue()              // Dequeue el primero
		fmt.Println("queue dequeue:", v) // Imprimimos el valor extraído
	}
	fmt.Println("Dictionary")
	// Dict demo
	d := NewDict[string, int]()        // Creamos un Dict(diccionario) con claves string y valores int
	d.Set("manzana", 3)                // Asignamos "manzana" -> 3
	d.Set("pera", 5)                   // Asignamos "pera" -> 5
	if v, ok := d.Get("manzana"); ok { // Intentamos obtener "manzana"
		fmt.Println("Cantidad de manzana:", v) // Imprimimos el valor (esperado 3)
	}
	fmt.Println("dict tiene sandia?", d.Has("sandia")) // Consultamos si existe "sandia" (esperado false)
	fmt.Println("dict len:", d.Len())                  // Imprimimos la cantidad de entradas (esperado 2)
	fmt.Println("keys:", d.Keys())                     // Mostramos las claves
	fmt.Println("values:", d.Values())                 // Mostramos los valores
	d.Range(func(k string, v int) bool {               // Recorremos con un callback que siempre devuelve true
		fmt.Printf("%s -> %d\n", k, v) // Imprimimos cada par clave-valor
		return true                    // Continuamos hasta terminar
	})
	d.Delete("pera")                           // Eliminamos la clave "pera"
	fmt.Println("after delete, len:", d.Len()) // Imprimimos el nuevo tamaño (esperado 1)
}
