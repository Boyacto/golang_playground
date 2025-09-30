package ds

// Stack is a generic LIFO stack.
// It supports Push, Pop, Peek, Len, IsEmpty, Clear, and Iterate.
type Stack[T any] struct {
	data []T
}

// NewStack creates an empty Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element. The boolean is false if empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	lastIdx := len(s.data) - 1
	v := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return v, true
}

// Peek returns the top element without removing it. The boolean is false if empty.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.data)
}

// IsEmpty returns true if the stack has no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	// Avoid memory leak by zeroing slice
	var zero T
	for i := range s.data {
		s.data[i] = zero
	}
	s.data = s.data[:0]
}

// Iterate calls fn for each element from top to bottom.
// If fn returns false, iteration stops early.
func (s *Stack[T]) Iterate(fn func(T) bool) {
	for i := len(s.data) - 1; i >= 0; i-- {
		if !fn(s.data[i]) {
			return
		}
	}
}
