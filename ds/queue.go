package ds

// Queue is a generic FIFO queue.
// It supports Enqueue, Dequeue, Front, Back, Len, IsEmpty, Clear, and Iterate.
type Queue[T any] struct {
	data []T
	head int
	tail int
	size int
}

// NewQueue creates an empty Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}

// Enqueue adds an element at the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	// Use circular buffer growth strategy
	if q.size == 0 {
		q.data = append(q.data, value)
		q.head = 0
		q.tail = 0
		q.size = 1
		return
	}
	if q.size == len(q.data) {
		// Need to grow: allocate larger slice and copy in order
		newCap := len(q.data) * 2
		if newCap == 0 {
			newCap = 1
		}
		buf := make([]T, newCap)
		// Copy elements in FIFO order
		for i := 0; i < q.size; i++ {
			buf[i] = q.data[(q.head+i)%len(q.data)]
		}
		q.data = buf
		q.head = 0
		q.tail = q.size - 1
	}
	// Insert at next tail position
	q.tail = (q.tail + 1) % len(q.data)
	q.data[q.tail] = value
	q.size++
}

// Dequeue removes and returns the front element. The boolean is false if empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	v := q.data[q.head]
	q.data[q.head] = zero
	q.head = (q.head + 1) % len(q.data)
	q.size--
	return v, true
}

// Front returns the front element without removing it. The boolean is false if empty.
func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	return q.data[q.head], true
}

// Back returns the last element without removing it. The boolean is false if empty.
func (q *Queue[T]) Back() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	return q.data[q.tail], true
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	return q.size
}

// IsEmpty returns true if the queue has no elements.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Clear removes all elements from the queue.
func (q *Queue[T]) Clear() {
	var zero T
	for i := range q.data {
		q.data[i] = zero
	}
	q.head = 0
	q.tail = 0
	q.size = 0
}

// Iterate calls fn for each element from front to back.
// If fn returns false, iteration stops early.
func (q *Queue[T]) Iterate(fn func(T) bool) {
	for i := 0; i < q.size; i++ {
		if !fn(q.data[(q.head+i)%len(q.data)]) {
			return
		}
	}
}
