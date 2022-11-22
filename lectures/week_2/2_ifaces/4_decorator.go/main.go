package main

import "fmt"

type iStuff interface {
	DoStuff()
}

type realStuff string

func (r realStuff) DoStuff() {
	fmt.Println(r)
}

type fakeStuff int

func (f fakeStuff) DoStuff() {
	fmt.Println("It's a trap")
}

type stuff struct {
	iStuff
	Name string
}

func (s stuff) SomeComplex() {
	if s.iStuff != nil {
		s.DoStuff()
	} else {
		fmt.Println("can't do stuff")
	}
}

func showDecorator(in iStuff) {
	fmt.Printf("%T: %+v\n", in, in)
}

func main() {
	r := realStuff("Hey")
	f := fakeStuff(0)

	rS := stuff{r, "stuff"}
	showDecorator(&rS)
	rS.SomeComplex()

	fS := stuff{f, "fake"}

	fS.SomeComplex()
	fS.DoStuff()
	showDecorator(&fS)
	pS := stuff{Name: "panic"}
	showDecorator(&pS)
	pS.SomeComplex()
	// pS.DoStuff() !!panic
}
