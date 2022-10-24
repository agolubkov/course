package main

import (
	"fmt"
	"time"
)

type Ny func() func(string) string

func main() {
	var myTime Ny
	myTime = sss()
	mt2 := myTime
	mt2()
	f := func() {
		str := myTime()("Johnny")
		println(str)
	}

	defer func() {
		println("defer executed")
	}()

	f()
}

func sss() func() func(string) string {
	start := time.Now()
	return func() func(string) string {
		fmt.Printf("Time from start %v", time.Since(start))
		return getTime
	}
}

func getTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Time from start %v\n", time.Since(start))
	}
}

func getTime(name string) string {
	return fmt.Sprintf("Hi, %s! %v", name, time.Now().String())
}
