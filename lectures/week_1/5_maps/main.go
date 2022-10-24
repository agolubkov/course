package main

import "fmt"

var mm map[string]string

func main() {
	fmt.Println("uninitialized map", mm, len(mm), mm == nil)
	// panic: assignment to entry in map
	// mm["test"] = "ok"

	// повна ініціалізація
	// var mm2 map[string]string = map[string]string{}
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// коротка ініціалізація
	mm3 := make(map[string]string)
	mm3["firstName"] = "Vasily"
	fmt.Println(mm3)

	// Отримання значення
	firstName := mm3["firstName"]
	fmt.Printf("firstName [%s], len: %d\n", firstName, len(firstName))

	// є звернутися до відсутнього ключа = значення за замовчуванням
	lastName := mm3["lastName"]
	fmt.Printf("lastName [%s], len: %d\n", lastName, len(lastName))

	// Перевірка те що, що значення є
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is:", lastName, ", exist:", ok)

	// Тільки отримання ознаки існування
	_, exist := mm3["firstName"]
	fmt.Println("fistName exist:", exist)

	// видалення значення
	delete(mm3, "firstName")
	_, exist = mm3["firstName"]
	fmt.Println("fistName exist:", exist)
}
