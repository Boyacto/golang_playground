# Estructuras de Datos en Go: Stack, Queue y Dictionary (ordenado)

## Descripción breve
- Implementé desde cero tres estructuras genéricas en Go 1.22:
  - `Stack[T]` (LIFO) en `ds/stack.go`
  - `Queue[T]` (FIFO) en `ds/queue.go`
  - `OrderedMap[K,V]` (diccionario con orden de inserción) en `ds/dictionary.go`
- No usé librerías externas; solo la biblioteca estándar del lenguaje.
- Esta solución fue asistida por una herramienta de IA para acelerar la escritura, pero el diseño y validación se describen aquí.

## Cómo ejecutar
```bash
# Estando en el directorio del proyecto
go run .
```

La salida muestra una pequeña demo en `main.go` que ejerce las operaciones principales de cada estructura.

## API principal
- Stack:
  - `Push(T)`, `Pop() (T,bool)`, `Peek() (T,bool)`, `Len() int`, `IsEmpty() bool`, `Clear()`, `Iterate(func(T) bool)`
- Queue:
  - `Enqueue(T)`, `Dequeue() (T,bool)`, `Front() (T,bool)`, `Back() (T,bool)`, `Len() int`, `IsEmpty() bool`, `Clear()`, `Iterate(func(T) bool)`
- OrderedMap (Dictionary con orden de inserción):
  - `Set(K,V)`, `Get(K) (V,bool)`, `Has(K) bool`, `Delete(K) bool`, `Len() int`, `Clear()`, `Keys() []K`, `Values() []V`, `Iterate(func(K,V) bool)`

## Diseño y decisiones
- `Stack` usa un `[]T` para operaciones amortizadas O(1).
- `Queue` implementa un buffer circular con crecimiento dinámico para mantener `Enqueue`/`Dequeue` en O(1) amortizado.
- `OrderedMap` combina:
  - `map[K]*node` para acceso O(1) por clave
  - lista doblemente enlazada para preservar el orden de inserción y soportar borrados en O(1) sin perder orden.
  - `Set` actualiza el valor si la clave ya existe, sin cambiar su posición en el orden.

## Test-cases utilizados (validación manual en la demo)
1. Stack
   - Inserciones múltiples (`Push`), `Peek` del tope, `Pop` y verificación de `Len` decreciente.
   - Iteración de tope a base con `Iterate` y parada temprana (implícita en API).
   - `Clear` seguido de `IsEmpty` y `Len==0`.
2. Queue
   - Encolar varios elementos (`Enqueue`), verificar `Front` y `Back`.
   - `Dequeue` mantiene orden FIFO y actualiza `Len`.
   - Iteración de frente a fondo con `Iterate` y `Clear` al final.
   - Casos de crecimiento de capacidad: encolado hasta agrandar el buffer y continuar con operaciones.
3. OrderedMap (Dictionary)
   - `Set` de claves nuevas y `Get` correcto.
   - Actualización de valor con `Set` sobre clave existente no cambia el orden (ver `Keys()` antes/después).
   - `Has` para presencia, `Delete` elimina y orden se actualiza; `Keys()` y `Values()` reflejan el cambio.
   - `Iterate` en orden de inserción; `Clear` deja el tamaño en 0.

> Nota: Para convertir estos test-cases en pruebas automatizadas, se pueden crear archivos en `*_test.go` usando `testing` y replicar los escenarios anteriores con aserciones (`t.Fatalf`, `t.Errorf`).
