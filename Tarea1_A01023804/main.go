package main // Declaramos el paquete principal: punto de entrada del programa

import ( // Bloque de importaciones
	"fmt" // Paquete para imprimir por consola y formatear cadenas
)

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

// ===== Dictionary (Map wrapper) =====
// En Go ya existe map como built-in. Aquí hacemos un “wrapper” con métodos básicos.
type Dict[K comparable, V any] struct { // Dict genérico con K comparable (requisito, hay types que no son comparables en golang) y V cualquiera
	m map[K]V // El almacenamiento interno es un map nativo de Go
}

func NewDict[K comparable, V any]() *Dict[K, V] { // Constructor del diccionario
	return &Dict[K, V]{m: make(map[K]V)} // Inicializamos el map con make y devolvemos el pointer al Dict
}

func (d *Dict[K, V]) Set(k K, v V) { // Set: asigna v a la clave k
	d.m[k] = v // Escritura directa en el map
}

func (d *Dict[K, V]) Get(k K) (V, bool) { // Get: obtiene valor por clave (bool indica si existe ese valor)
	v, ok := d.m[k] // Lectura con “comma-ok” para saber si la clave existe
	return v, ok    // Devolvemos el valor y el flag ok
}

func (d *Dict[K, V]) Has(k K) bool { // Has: true si la clave existe
	_, ok := d.m[k] // Esta interesante, de esta forma se puede ignorar el valor, lo que importa es valir booleano de ok
	return ok       // Devolvemos existencia
}

func (d *Dict[K, V]) Delete(k K) { // Delete: elimina la entrada por clave
	delete(d.m, k) // Usamos la función built-in delete para mapas
}

func (d *Dict[K, V]) Len() int { return len(d.m) } // Len: cantidad de pares almacenados

func (d *Dict[K, V]) Clear() { // Clear: reinicia el map
	d.m = make(map[K]V) // Asigna un nuevo map vacío (el anterior queda para GC si no hay más referencias)
}

func (d *Dict[K, V]) Keys() []K { // Keys: devuelve un slice con todas las claves (orden no definido)
	keys := make([]K, 0, len(d.m)) // Creamos el slice con capacidad igual al tamaño del map
	for k := range d.m {           // Iteramos sobre las claves del map
		keys = append(keys, k) // Agregamos cada clave al slice
	}
	return keys // Devolvemos el slice de claves
}

func (d *Dict[K, V]) Values() []V { // Values: devuelve un slice con todos los valores (orden no definido)
	values := make([]V, 0, len(d.m)) // Preasignamos capacidad
	for _, v := range d.m {          // Iteramos sobre valores del map
		values = append(values, v) // Agregamos cada valor al slice
	}
	return values // Devolvemos el slice de valores
}

func (d *Dict[K, V]) Range(f func(K, V) bool) { // Range: recorre pares y permite corte temprano si f devuelve false
	for k, v := range d.m { // Iteramos sobre todas las entradas del map
		if !f(k, v) { // Si el callback devuelve false
			return // Terminamos la iteración anticipadamente
		}
	}
}

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
