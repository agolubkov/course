package main

import (
	"fmt"
	"log"
)

func main() {
	showMeTheMoney()

	stuff := []int{10, 2, 3}
	res, _ := sumNatural2(stuff...)
	a := 13
	b := 42
	fmt.Println(a, b)
	sum(&a, b)
	fmt.Println(a, b)
	fmt.Println("Res", res)
	fmt.Println(sumLight("sum:\t %d%s", "$", 1, 2, 3, 4, 5))
	fmt.Println(sumLight("stuff:\t %d%s", "$", stuff...))
	println(sumMore(stuff...))
	if _, err := sumOnlyNatural(stuff...); err != nil {
		log.Printf("sumOnlyNatural error: %v", err)
	}
}

func showMeTheMoney() {
	fmt.Printf("$$$$")
}

// Декілька вхідних значень
func sum(i *int, j int) {
	*i += j
}

// Спрощений запис для кількох значень однакового типу
func sumLight(format, suffix string, summed ...int) string {
	var out int
	for _, v := range summed {
		out += v
	}
	return fmt.Sprintf(format, out, suffix)
}

// Для отримання довільного списку однотипних значень
func sumMore(stuff ...int) (res int) {
	fmt.Printf("\n%T\n", stuff)
	for i := range stuff {
		res += stuff[i]
	}
	return
}

// Повернення кількох значень
func sumOnlyNatural(stuff ...int) (int, error) {
	res := 0
	for i := range stuff {
		if stuff[i] < 0 {
			return 0, fmt.Errorf("only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, nil
}

// Повернення іменованих значень
func sumNatural2(stuff ...int) (res int, err error) {
	resPo := &res
	fmt.Printf("\n%T\n", resPo)

	res = *resPo

	for i := range stuff {
		if stuff[i] < 0 {
			err = fmt.Errorf("only natural numbers expected - given %d", stuff[i])
			return
		}
		res += stuff[i]
	}
	return res, nil
}
