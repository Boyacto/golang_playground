package main // El paquete principal: punto de entrada del programa

// ===== Queue (FIFO) =====
// Implementación simple basada en slice con índice head amortizando O(1) en dequeue.
type Queue[T any] struct { // Definimos Queue genérica de T
	data []T // Slice de T para datos
	head int // Índice del primer elemento válido (frente)
}

func (q *Queue[T]) Enqueue(v T) { // Enqueue: agrega al final (cola)
	q.data = append(q.data, v) // Usamos append para insertar
}

func (q *Queue[T]) Dequeue() (T, bool) { // Dequeue: extrae el primer elemento (bool indica éxito)
	var zero T       // Valor cero de T para el caso vacío
	if q.IsEmpty() { // Si la cola está vacía
		return zero, false // No hay nada que devolver
	}
	v := q.data[q.head] // Tomamos el elemento al frente (posición head)
	q.head++            // Avanzamos el índice head para “consumir” ese elemento
	// Si ya consumimos más de la mitad, compactamos para liberar memoria retenida
	if q.head > len(q.data)/2 { // Condición de compactación
		q.data = append([]T(nil), q.data[q.head:]...) // Copiamos los elementos restantes a un nuevo slice
		q.head = 0                                    // Reiniciamos head porque ahora el primer elemento está en índice 0
	}
	return v, true // Devolvemos el valor dequeued y true
}
func (q *Queue[T]) Peek() (T, bool) { // Peek: mira el frente sin quitarlo
	var zero T       // Valor cero de T si está vacía
	if q.IsEmpty() { // Si no hay elementos
		return zero, false // Indicamos vacío
	}
	return q.data[q.head], true // Devolvemos el elemento en el frente y true
}

func (q *Queue[T]) Len() int      { return len(q.data) - q.head } // Len: cantidad efectiva = total - consumidos
func (q *Queue[T]) IsEmpty() bool { return q.Len() == 0 }         // IsEmpty: true si Len es 0
func (q *Queue[T]) Clear()        { q.data, q.head = nil, 0 }     // Clear: liberamos referencia y reiniciamos head
