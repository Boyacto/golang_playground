package ds

// OrderedMap is a generic dictionary that preserves insertion order of keys.
// It supports Set, Get, Has, Delete, Len, Clear, Keys, Values, and Iterate.
type OrderedMap[K comparable, V any] struct {
	index map[K]*orderedNode[K, V]
	head  *orderedNode[K, V]
	tail  *orderedNode[K, V]
	len   int
}

type orderedNode[K comparable, V any] struct {
	key   K
	value V
	prev  *orderedNode[K, V]
	next  *orderedNode[K, V]
}

// NewOrderedMap creates an empty ordered map.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		index: make(map[K]*orderedNode[K, V]),
	}
}

// Set inserts or updates the value for key. Insertion order is preserved; updates do not move the key.
func (m *OrderedMap[K, V]) Set(key K, value V) {
	if n, ok := m.index[key]; ok {
		n.value = value
		return
	}
	n := &orderedNode[K, V]{key: key, value: value}
	m.index[key] = n
	if m.tail == nil {
		m.head = n
		m.tail = n
		m.len = 1
		return
	}
	n.prev = m.tail
	m.tail.next = n
	m.tail = n
	m.len++
}

// Get returns the value for key and true if present.
func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	var zero V
	n, ok := m.index[key]
	if !ok {
		return zero, false
	}
	return n.value, true
}

// Has reports whether key exists.
func (m *OrderedMap[K, V]) Has(key K) bool {
	_, ok := m.index[key]
	return ok
}

// Delete removes the key if present and returns true if removed.
func (m *OrderedMap[K, V]) Delete(key K) bool {
	n, ok := m.index[key]
	if !ok {
		return false
	}
	delete(m.index, key)
	m.unlink(n)
	m.len--
	return true
}

func (m *OrderedMap[K, V]) unlink(n *orderedNode[K, V]) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		m.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		m.tail = n.prev
	}
	// clear node's links to help GC in long-lived maps
	n.prev = nil
	n.next = nil
}

// Len returns the number of key-value pairs.
func (m *OrderedMap[K, V]) Len() int { return m.len }

// Clear removes all key-value pairs.
func (m *OrderedMap[K, V]) Clear() {
	if m.len == 0 {
		return
	}
	for k := range m.index {
		delete(m.index, k)
	}
	m.head = nil
	m.tail = nil
	m.len = 0
}

// Keys returns all keys in insertion order.
func (m *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.len)
	for n := m.head; n != nil; n = n.next {
		keys = append(keys, n.key)
	}
	return keys
}

// Values returns all values in insertion order.
func (m *OrderedMap[K, V]) Values() []V {
	vals := make([]V, 0, m.len)
	for n := m.head; n != nil; n = n.next {
		vals = append(vals, n.value)
	}
	return vals
}

// Iterate calls fn for each key and value in insertion order.
// If fn returns false, iteration stops early.
func (m *OrderedMap[K, V]) Iterate(fn func(K, V) bool) {
	for n := m.head; n != nil; n = n.next {
		if !fn(n.key, n.value) {
			return
		}
	}
}
