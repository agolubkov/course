package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func PrintSome[T any](in T) {
	fmt.Println(in)
}

type MyConstraint interface {
	~string | ~int
}

type (
	myStrings []MyInt
	MyInt     int
)

func Each[S ~[]E, E MyConstraint](in S, call func(e E)) {
	for _, v := range in {
		call(v)
	}
}

func main() {
	println(Min(1.22, 1.23))
	m := GMin[float64](1.22, 1.23)
	println(m)
	m2 := GMin[int](3, 1)
	println(m2)
	// shortened syntax
	println(GMin(1.33, 1.32))
	println(GMin(2, 1))
	println(GMin("2", "1"))
	type doSome struct {
		DoSome   bool
		WhatToDo string
	}
	PrintSome(doSome{true, "something"})
	// GMin(doSome{true, "something"}, doSome{true, "something"}) !!! Error
	var x, y int
	x, y = 9, 8
	println(GMin(x, y))
	Each(myStrings{1, 2, 4, 5, 6, 7}, func(in MyInt) {
		println(in * 42)
	})
}
