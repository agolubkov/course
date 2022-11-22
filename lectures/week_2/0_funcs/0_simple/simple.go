package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	showMeTheMoney()

	stuff := []int{10, 2, 3}
	res, inputType, _ := sumNatural2(stuff...)
	if res, inputType, err := sumNatural2(1, 2, -3, 4); err != nil {
		log.Printf("%d, %s, %s", res, inputType, err.Error())
	}
	fmt.Println(inputType)
	a := 13
	b := 42
	fmt.Println(a, b)
	if sum(&a, b) != nil {
		return
	}
	if err := sum(nil, b); err != nil {
		log.Printf("code error: %v", err)
	}
	fmt.Println(a, b)
	fmt.Println("Res", res)
	fmt.Println(sumLight("sum:\t %d%s", "$"))
	fmt.Println(sumLight("sum:\t %d%s", "$", 1, 2, 3, 4, 5))
	fmt.Println(sumLight("stuff:\t %d%s", "$", stuff...))
	println(sumMore(stuff...))
	if _, err := sumOnlyNatural(stuff...); err != nil {
		log.Printf("sumOnlyNatural error: %v", err)
	}
	if _, err := sumOnlyNatural(1, 2, -3, 4); err != nil {
		log.Printf("sumOnlyNatural error: %v", err)
	}
	some("")
	some("one", true)
	some("one", false)
	someParams(inPar{two: a})
}

func showMeTheMoney() {
	fmt.Printf("$$$$")
}

// Декілька вхідних значень
func sum(i *int, j int) error {
	if i == nil {
		return errors.New("nil pointer")
	}
	*i += j
	return nil
}

func some(message string, out ...bool) {
	if len(out) > 0 && out[0] {
		println(message)
		return
	}
	println("nothing to say")
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
	fmt.Printf("\n%T : %v\n", stuff, stuff)
	for i := range stuff {
		res += stuff[i]
	}
	return res
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
func sumNatural2(stuff ...int) (int, string, error) {
	var res int
	resPo := &res
	valueFormat := fmt.Sprintf("%T", resPo)

	res = *resPo

	for i := range stuff {
		if stuff[i] < 0 {
			return 0, valueFormat, fmt.Errorf("only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, valueFormat, nil
}

type inPar struct {
	one  string
	two  int
	some *int
}

func someParams(in inPar,
) string {
	if in.one == "" {
		in.one = "data %09d %v"
	}
	return fmt.Sprintf(in.one, in.two, in.some)
}
