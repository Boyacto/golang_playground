package main

import (
	"fmt"

	"golang_playground/ds"
)

func demoStack() {
	fmt.Println("== Stack demo (LIFO) ==")
	s := ds.NewStack[int]()
	fmt.Println("IsEmpty:", s.IsEmpty())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Len after pushes:", s.Len())
	if top, ok := s.Peek(); ok {
		fmt.Println("Peek:", top)
	}
	v, _ := s.Pop()
	fmt.Println("Pop:", v)
	fmt.Println("Len after pop:", s.Len())
	fmt.Print("Iterate top->bottom:")
	s.Iterate(func(x int) bool { fmt.Print(" ", x); return true })
	fmt.Println()
	s.Clear()
	fmt.Println("Len after Clear:", s.Len(), "IsEmpty:", s.IsEmpty())
}

func demoQueue() {
	fmt.Println("\n== Queue demo (FIFO) ==")
	q := ds.NewQueue[string]()
	fmt.Println("IsEmpty:", q.IsEmpty())
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")
	fmt.Println("Len after enqueues:", q.Len())
	if f, ok := q.Front(); ok {
		fmt.Println("Front:", f)
	}
	if b, ok := q.Back(); ok {
		fmt.Println("Back:", b)
	}
	x, _ := q.Dequeue()
	fmt.Println("Dequeue:", x)
	fmt.Println("Len after dequeue:", q.Len())
	fmt.Print("Iterate front->back:")
	q.Iterate(func(s string) bool { fmt.Print(" ", s); return true })
	fmt.Println()
	q.Clear()
	fmt.Println("Len after Clear:", q.Len(), "IsEmpty:", q.IsEmpty())
}

func demoOrderedMap() {
	fmt.Println("\n== OrderedMap demo (Dictionary with insertion order) ==")
	m := ds.NewOrderedMap[string, int]()
	fmt.Println("Len:", m.Len())
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)
	fmt.Println("Len after inserts:", m.Len())
	if v, ok := m.Get("two"); ok {
		fmt.Println("Get('two'):", v)
	}
	m.Set("two", 22) // update should not change order
	fmt.Println("Keys:", m.Keys())
	fmt.Println("Values:", m.Values())
	fmt.Print("Iterate (k=v):")
	m.Iterate(func(k string, v int) bool { fmt.Printf(" %s=%d", k, v); return true })
	fmt.Println()
	fmt.Println("Has('one'):", m.Has("one"))
	fmt.Println("Delete('one'):", m.Delete("one"))
	fmt.Println("Has('one'):", m.Has("one"))
	fmt.Println("Keys after delete:", m.Keys())
	m.Clear()
	fmt.Println("Len after Clear:", m.Len())
}

func main() {
	demoStack()
	demoQueue()
	demoOrderedMap()
}
