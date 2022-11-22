package main

import "fmt"

type Flyer interface {
	Fly()
	Greet()
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name + " is flying")
}

func (b Bird) Greet() {
	fmt.Println("Hey there")
}

func DoFly(f Flyer) {
	f.Greet()
	f.Fly()
}

type Mig45 struct{}

func (m Mig45) Fly() {
	fmt.Println("Mig Flied away")
}

func main() {
	duckPlane := &Bird{"Duck plane"}

	GoFly(duckPlane)
	GoFly("something")
	GoFly(1111)
	GoFly(&struct {
		Alpha string
		Beta  int
	}{
		"10",
		10,
	})
}

func GoFly(x interface{}) {
	fmt.Printf("%T: %+v\n", x, x)
	f, isFlyer := x.(Flyer)
	if isFlyer {
		f.Fly()
	}

	// b := f.(Bird)
	if _, ok := x.(string); !ok {
		fmt.Println("not a string")
	} else {
		fmt.Println("this is string")
	}
	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}
