package main // El paquete principal: punto de entrada del programa
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
