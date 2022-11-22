// Запуск go рутин
// іменовані функції чи анонімні
package main

import (
	"fmt"
)

func main() {
	fmt.Println("старт")
	// Можна запустити функцію
	go process(0)
	// Можна запустити анонімну функцію
	go func() {
		fmt.Println("Анонімний запуск")
	}()

	// Можемо запустити багато горутин
	for i := 0; i < 1000; i++ {
		go process(i)
	}

	// Потрібно дочекатися завершення виконання
	_, _ = fmt.Scanln()
}

func process(i int) {
	fmt.Println("обробка: ", i)
}
