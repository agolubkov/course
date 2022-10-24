package main

import (
	"fmt"
)

// Приклад представляє тип з різними полями.
type example struct {
	Flag    bool
	counter int16
	pi      float32
}

func main() {
	// Оголошення змінної типу example встановлено на її
	// нульове значення.
	var e1 example
	var e11 example = example{
		Flag:    false,
		counter: 11,
		pi:      3.14,
	}

	// Відбити значення.
	fmt.Printf("%+v\n", e1)
	e1 = e11
	fmt.Printf("%+v\n", e1)
	// Оголошення змінної типу example та ініціалізація за допомогою структурного літерала.
	e2 := example{
		Flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	e3 := example{}

	// Display the field values.
	fmt.Println("Flag", e2.Flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
	fmt.Println("Pi", e3.pi)
	e4 := struct {
		Flag    bool
		counter int16
		pi      float32
	}{}
	e4.Flag = true
	fmt.Printf("%+v\n", e4.Flag)
}
