// Один із механізмів синхронізації - канали
// Канали, це об'єкт через який можна забезпечити взаємодію кількох горутин
// У приймаючій (або повертаючій) канал функції, можна вказати напрям роботи з каналом
// Тільки для читання - "<-chan" або тільки для запису "chan<-"
package main

import "fmt"

// var c chan int

func main() {
	// Створюємо канал
	c := make(chan string)
	// стартуємо пишучу горутину
	go greet(c)
	for i := 0; i < 5; i++ {
		// Читаємо пару рядків із каналу
		fmt.Println(<-c, ",", <-c)
	}

	stuff := make(chan int, 7)
	for i := 0; i < 19; i = i + 3 {
		stuff <- i
	}
	close(stuff)
	fmt.Println("Res", process(stuff))
}

func greet(c chan<- string) {
	// Запускаємо нескінченний цикл
	for {
		// і пишемо в канал кілька рядків
		// Підпрограму буде заблоковано до того, як хтось захоче прочитати з каналу
		c <- fmt.Sprintf("Владика")
		c <- fmt.Sprintf("Штурмовик")
	}
}

func process(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}
