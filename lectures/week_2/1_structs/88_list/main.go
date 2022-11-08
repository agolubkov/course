package main

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

// print nodes recursively
func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func main() {
	first := node{value: 4}
	second := node{value: 5, prev: &first}
	third := node{value: 6, prev: &second}
	first.next = &second
	second.next = &third

	current := &first
	for current.next != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println(current.value)
	fmt.Println("And go back ...")
	for current.prev != nil {
		fmt.Println(current.value)
		current = current.prev
	}
	fmt.Println(current.value)
	fmt.Println("Print recursively")
	printNodeValue(&first)
}
