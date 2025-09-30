package main // El paquete principal: punto de entrada del programa
// ===== Queue (FIFO) =====

// Queue implementa una cola FIFO genérica usando slice con índice head.
type Queue[T any] struct { // Tipo genérico con parámetro T
	data []T // data almacena todos los elementos en un slice
	head int // head apunta al índice del primer elemento válido (frente)
}

// Enqueue agrega un elemento al final de la cola.
func (q *Queue[T]) Enqueue(v T) { // Receptor puntero para mutar estado
	q.data = append(q.data, v) // insertamos v al final del slice
}

// Dequeue quita y devuelve el elemento del frente; bool indica éxito.
func (q *Queue[T]) Dequeue() (T, bool) { // Retorna (valor, ok) para manejo de vacío
	var zero T       // valor cero de T
	if q.IsEmpty() { // si no hay elementos útiles
		return zero, false // no se puede extraer, devolvemos zero, false
	}
	v := q.data[q.head] // tomamos el elemento en la posición head (frente)
	q.head++            // avanzamos head para “consumir” ese elemento

	// Para evitar retener memoria innecesaria, compactamos si ya consumimos más de la mitad
	if q.head > len(q.data)/2 { // condición de compactación para amortizar costo
		q.data = append([]T(nil), q.data[q.head:]...) // copiamos la porción útil a un slice nuevo
		q.head = 0                                    // reseteamos head al inicio del slice
	}

	return v, true // devolvemos el valor y ok=true
}

// Peek devuelve el elemento del frente sin extraerlo.
func (q *Queue[T]) Peek() (T, bool) { // Mismo patrón (valor, ok)
	var zero T       // valor cero de T
	if q.IsEmpty() { // si no hay elementos
		return zero, false // no se puede mirar, devolvemos zero, false
	}
	return q.data[q.head], true // devolvemos el elemento al frente y ok=true
}

// Len devuelve cuántos elementos hay disponibles en la cola.
func (q *Queue[T]) Len() int { // cálculo basado en head
	return len(q.data) - q.head // total almacenado menos los ya consumidos
}

// IsEmpty indica si la cola no tiene elementos.
func (q *Queue[T]) IsEmpty() bool { // true si Len()==0
	return q.Len() == 0 // reutilizamos Len()
}

// Clear borra el contenido para que el GC pueda liberar memoria.
func (q *Queue[T]) Clear() { // reset completo
	q.data = nil // soltamos referencia del slice
	q.head = 0   // reiniciamos head
}
