package main

import (
	"fmt"
)

// Приклад представляє тип з різними полями.
type Example struct {
	Flag    bool
	counter int16
	pi      float32
}

var some = Example{counter: 4, Flag: true, pi: 3.15}

type contact struct {
	email string
	phone string
	age   int
}

type person struct {
	name string
	age  int
	contact
}

func main() {
	// Оголошення змінної типу Example встановлено на її
	// нульове значення.
	var e1 Example
	var e11 Example = Example{
		Flag:    true,
		counter: 11,
		pi:      3.14,
	}

	// Відбити значення.
	fmt.Printf("%+v\n", e1)
	e1 = e11
	e11.Flag = false
	fmt.Printf("e1: %+v, e11: %+v\n", e1, e11)

	// Оголошення змінної типу Example та ініціалізація за допомогою структурного літерала.
	e2 := Example{
		Flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	e3 := Example{}

	// Display the field values.
	fmt.Println("Flag", e2.Flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
	fmt.Println("Pi", e3.pi)
	e4 := struct {
		Flag    bool
		counter int16
		pi      float32
	}{
		counter: 4,
	}
	e4.Flag = true
	fmt.Printf("%+v\n", e4)
	fmt.Printf("%+v\n", some)
	tom := person{
		name: "Tom",
		age:  24,
		contact: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.email = "tomcat@gmail.com"
	tom.age = 33
	tom.contact.age = 22

	fmt.Printf("%+v\n", tom)
}
