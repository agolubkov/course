package main

import "fmt"

type linkedNode[T any] struct {
	value T
	next  *linkedNode[T]
	prev  *linkedNode[T]
}

type linkedList[T any] struct {
	first *linkedNode[T]
	last  *linkedNode[T]
}

func newLinkedList[T any](in T) *linkedList[T] {
	return &linkedList[T]{first: &linkedNode[T]{value: in}}
}

func (ll *linkedList[T]) Add(v T) {
	if ll.first == nil {
		ll.Insert(v)
		return
	}
	next := &linkedNode[T]{value: v}
	last := ll.last
	last.next = next
	next.prev = last
	ll.last = next
}

func (ll *linkedList[T]) Insert(v T) {
	newNode := &linkedNode[T]{value: v}
	if ll.first == nil {
		ll.first = newNode
		ll.last = ll.first
		return
	}
	prev := ll.first
	prev.prev = newNode
	newNode.next = prev
	ll.first = newNode
}

func (ll *linkedList[T]) Each(call func(T)) {
	if ll.first == nil {
		return
	}
	current := ll.first
	for current != nil {
		call(current.value)
		current = current.next
	}
}

type iLinkedList[T any] interface {
	Insert(v T)
	Add(v T)
	Each(call func(T))
}

func main() {
	var m iLinkedList[int]
	m = &linkedList[int]{}
	for i := 0; i < 10; i++ {
		m.Add(i)
		m.Insert(i)
	}
	m.Each(func(in int) {
		println(in)
	})
	initData := [...]string{"some", "body", "help", "me", "!"}
	x := linkedList[string]{}
	for _, v := range initData {
		x.Add(v)
	}
	x.Each(func(in string) {
		println(in)
	})
	out := newLinkedList("1")
	fmt.Printf("%T, %+v", *out, *out)
}
